package solutions

import (
	"fmt"
	"math"
)

// https://projecteuler.net/problem=11
//
// "What is the greatest product of four adjacent numbers in the same direction
// (up, down, left, right, or diagonally) in the 20 x 20 grid?"
//
// Generally: Find the greatest product of n adjacent numbers in the same
// direction in a m x m grid.

type LargestProductGrid struct{}

type LargestProductGridInput struct {
	n           int
	m           int
	grid        [][]int
	coordinates [][][]int
}

func (l LargestProductGrid) description(n, m string) string {
	return fmt.Sprintf(
		"Find the greatest product of %s adjacent numbers in the same direction in a %s x %s grid",
		n, m, m,
	)
}

func (l LargestProductGrid) PrintHelp() {
	fmt.Println(l.description("n", "m"))
	fmt.Println("This solution is not a general form!")
	fmt.Println("Usage: ./project_euler 11")
}

func (l LargestProductGrid) DefaultInput() interface{} {
	return LargestProductGridInput{
		n: 4,
		m: 20,
		grid: [][]int{
			{8, 2, 22, 97, 38, 15, 0, 40, 0, 75, 4, 5, 7, 78, 52, 12, 50, 77, 91, 8},
			{49, 49, 99, 40, 17, 81, 18, 57, 60, 87, 17, 40, 98, 43, 69, 48, 4, 56, 62, 0},
			{81, 49, 31, 73, 55, 79, 14, 29, 93, 71, 40, 67, 53, 88, 30, 3, 49, 13, 36, 65},
			{52, 70, 95, 23, 4, 60, 11, 42, 69, 24, 68, 56, 1, 32, 56, 71, 37, 2, 36, 91},
			{22, 31, 16, 71, 51, 67, 63, 89, 41, 92, 36, 54, 22, 40, 40, 28, 66, 33, 13, 80},
			{24, 47, 32, 60, 99, 3, 45, 2, 44, 75, 33, 53, 78, 36, 84, 20, 35, 17, 12, 50},
			{32, 98, 81, 28, 64, 23, 67, 10, 26, 38, 40, 67, 59, 54, 70, 66, 18, 38, 64, 70},
			{67, 26, 20, 68, 2, 62, 12, 20, 95, 63, 94, 39, 63, 8, 40, 91, 66, 49, 94, 21},
			{24, 55, 58, 5, 66, 73, 99, 26, 97, 17, 78, 78, 96, 83, 14, 88, 34, 89, 63, 72},
			{21, 36, 23, 9, 75, 0, 76, 44, 20, 45, 35, 14, 0, 61, 33, 97, 34, 31, 33, 95},
			{78, 17, 53, 28, 22, 75, 31, 67, 15, 94, 3, 80, 4, 62, 16, 14, 9, 53, 56, 92},
			{16, 39, 5, 42, 96, 35, 31, 47, 55, 58, 88, 24, 0, 17, 54, 24, 36, 29, 85, 57},
			{86, 56, 0, 48, 35, 71, 89, 7, 5, 44, 44, 37, 44, 60, 21, 58, 51, 54, 17, 58},
			{19, 80, 81, 68, 5, 94, 47, 69, 28, 73, 92, 13, 86, 52, 17, 77, 4, 89, 55, 40},
			{4, 52, 8, 83, 97, 35, 99, 16, 7, 97, 57, 32, 16, 26, 26, 79, 33, 27, 98, 66},
			{88, 36, 68, 87, 57, 62, 20, 72, 3, 46, 33, 67, 46, 55, 12, 32, 63, 93, 53, 69},
			{4, 42, 16, 73, 38, 25, 39, 11, 24, 94, 72, 18, 8, 46, 29, 32, 40, 62, 76, 36},
			{20, 69, 36, 41, 72, 30, 23, 88, 34, 62, 99, 69, 82, 67, 59, 85, 74, 4, 36, 16},
			{20, 73, 35, 29, 78, 31, 90, 1, 74, 31, 49, 71, 48, 86, 81, 16, 23, 57, 5, 54},
			{1, 70, 54, 71, 83, 51, 54, 69, 16, 92, 33, 48, 61, 43, 52, 1, 89, 19, 67, 48},
		},
	}
}

func (l LargestProductGrid) Describe(in interface{}) string {
	if in, ok := in.(LargestProductGridInput); ok {
		return fmt.Sprintf(l.description(fmt.Sprintf("%d", in.n), fmt.Sprintf("%d", in.m)))
	} else {
		panic("Input type assertion failed")
	}
}

func (l LargestProductGrid) Process(args []int) interface{} {
	fmt.Println("Note: This solution is not a general form!")

	baseIn := l.DefaultInput().(LargestProductGridInput)

	greatestValue := 0
	for _, row := range baseIn.grid {
		for _, value := range row {
			if value > greatestValue {
				greatestValue = value
			}
		}
	}

	coordinates := make([][][]int, greatestValue+1)
	for i, row := range baseIn.grid {
		for j, value := range row {
			coordinates[value] = append(coordinates[value], []int{i, j})
		}
	}

	baseIn.coordinates = coordinates
	return baseIn
}

// Updated algorithm:
//  1. Consume the grid into an array of arrays of points - the coordinates for
//     each number in the grid.
//  2. Visit the coordinates of the highest possible value in the grid n and
//     find the largest possible product m that uses n. Maybe use the sliding
//     window technique from 0008.
//  3. Divide m by the cube of the next highest value in the grid n-x - this is
//     the new floor, called y. If a window contains any digit less than or
//     equal to y, ignore it, because it can't be part of a product greater than
//     m that doesn't contain n.
//  4. Repeat from 2 for n-x. If n-x <= y, the current max is the global max.
func (l LargestProductGrid) Solve(in interface{}) int {
	if in, ok := in.(LargestProductGridInput); ok {
		gMax := 0
		floor := 0
		n := len(in.coordinates) - 1
		for n > floor {
			lMax := l.localMax(in.grid, n, in.coordinates[n], floor)
			if lMax > gMax {
				gMax = lMax
			}

			n--
			for len(in.coordinates[n]) == 0 {
				n--
			}

			floor = int(math.Floor(float64(gMax) / math.Pow(float64(n), 3)))
		}
		return gMax
	} else {
		panic("Input type assertion failed")
	}
}

func (l LargestProductGrid) localMax(grid [][]int, n int, coordinates [][]int, floor int) int {
	localMax := 0
	for _, coord := range coordinates {
		newMax := max(grid, coord[0], coord[1], floor)
		if newMax > localMax {
			localMax = newMax
		}
	}
	return localMax
}

func max(grid [][]int, i, j, floor int) int {
	maxes := []int{
		verticalWindow(grid, i, j, floor),
		horizontalWindow(grid, i, j, floor),
		tLbRDiagonal(grid, i, j, floor),
		bLtRDiagonal(grid, i, j, floor),
	}

	localMax := 0
	for _, max := range maxes {
		if max > localMax {
			localMax = max
		}
	}

	return localMax
}

// constant column j
func verticalWindow(grid [][]int, i, j, floor int) int {
	max := 0

	start := i - 4 + 1
	if start < 0 {
		start = 0
	}

windowLoop:
	for start <= i {
		if start+4 > len(grid) {
			break windowLoop
		}

		for k := start; k < start+4; k++ {
			if grid[k][j] <= floor {
				start = k + 1
				continue windowLoop
			}
		}

		product := 1
		for k := start; k < start+4; k++ {
			product *= grid[k][j]
		}

		if product > max {
			max = product
		}

		start++
	}

	return max
}

// constant row i
func horizontalWindow(grid [][]int, i, j, floor int) int {
	max := 0

	start := j - 4 + 1
	if start < 0 {
		start = 0
	}

windowLoop:
	for start <= j {
		if start+4 > len(grid[i]) {
			break windowLoop
		}

		for k := start; k < start+4; k++ {
			if grid[i][k] < floor {
				start = k + 1
				continue windowLoop
			}
		}

		product := 1
		for k := start; k < start+4; k++ {
			product *= grid[i][k]
		}

		if product > max {
			max = product
		}

		start++
	}

	return max
}

// i+1, j+1
func tLbRDiagonal(grid [][]int, i, j, floor int) int {
	max := 0

	startI := i - 4 + 1
	if startI < 0 {
		startI = 0
	}
	startJ := j - 4 + 1
	if startJ < 0 {
		startJ = 0
	}

windowLoop:
	for startI <= i && startJ <= j {
		if startI+4 > len(grid) || startJ+4 > len(grid[0]) {
			break windowLoop
		}

		for increment := 0; increment < 4; increment++ {
			if grid[startI+increment][startJ+increment] <= floor {
				startI = startI + increment + 1
				startJ = startJ + increment + 1
				continue windowLoop
			}
		}

		product := 1
		for increment := 0; increment < 4; increment++ {
			product *= grid[startI+increment][startJ+increment]
		}

		if product > max {
			max = product
		}

		startI++
		startJ++
	}

	return max
}

// i-1, j+1
func bLtRDiagonal(grid [][]int, i, j, floor int) int {
	max := 0

	startI := i + 4 - 1
	if startI >= len(grid) {
		startI = len(grid) - 1
	}
	startJ := j - 4 + 1
	if startJ < 0 {
		startJ = 0
	}

windowLoop:
	for startI >= i && startJ <= j {
		if startI-4 < 0 || startJ+4 > len(grid[0]) {
			break windowLoop
		}

		for increment := 0; increment < 4; increment++ {
			if grid[startI-increment][startJ+increment] <= floor {
				startI = startI - increment - 1
				startJ = startJ + increment + 1
				continue windowLoop
			}
		}

		product := 1
		for increment := 0; increment < 4; increment++ {
			product *= grid[startI-increment][startJ+increment]
		}

		if product > max {
			max = product
		}

		startI--
		startJ++
	}

	return max
}

// =======================================================================================
// Find the greatest product of 4 adjacent numbers in the same direction in a 20 x 20 grid
// Answer:           70600674
// Process Duration: 74.042µs
// Solve Duration:   7.056µs
// Total Duration:   81.098µs
// =======================================================================================
