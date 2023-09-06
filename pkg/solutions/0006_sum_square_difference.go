package solutions

import (
	"fmt"
	"math"
)

// https://projecteuler.net/problem=6
//
// "Find the difference between the sum of the squares of the first hundred
// natural numbers and the square of the sum."
//
// Generally: Find the difference between the sum of the squares of the first n
// natural numbers and the square of the sum.

type SumSquareDifference struct{}

type SumSquareDifferenceInput struct {
	n int
}

func (s SumSquareDifference) description(n string) string {
	return fmt.Sprintf(
		"Find the difference between the sum of squares of the first %s natural numbers and the square of the sum",
		n,
	)
}

func (s SumSquareDifference) PrintHelp() {
	fmt.Println(s.description("n"))
	fmt.Println("Usage: ./project_euler 6 n")
}

func (s SumSquareDifference) DefaultInput() interface{} {
	return SumSquareDifferenceInput{
		n: 100,
	}
}

func (s SumSquareDifference) Describe(in interface{}) string {
	if in, ok := in.(SumSquareDifferenceInput); ok {
		return fmt.Sprintf(s.description(fmt.Sprintf("%d", in.n)))
	} else {
		panic("Input type assertion failed")
	}
}

func (s SumSquareDifference) Process(args []int) interface{} {
	if len(args) == 0 {
		return s.DefaultInput()
	}

	return SumSquareDifferenceInput{
		n: args[0],
	}
}

func sumOfSquares(n int) int {
	total := 0
	for i := 1; i < n+1; i++ {
		total += int(math.Pow(float64(i), 2))
	}
	return total
}

func squareOfSums(n int) int {
	total := 0
	for i := 1; i < n+1; i++ {
		total += i
	}
	total = int(math.Pow(float64(total), 2))
	return total
}

func (s SumSquareDifference) Solve(in interface{}) int {
	if in, ok := in.(SumSquareDifferenceInput); ok {
		return squareOfSums(in.n) - sumOfSquares(in.n)
	} else {
		panic("Input type assertion failed")
	}
}

// =========================================================================================================
// Find the difference between the sum of squares of the first 100 natural numbers and the square of the sum
// Answer: 25164150
// Process Duration: 340ns
// Solve Duration:   6.182µs
// Total Duration:   6.522µs
// =========================================================================================================
