package solutions

import "fmt"

// https://projecteuler.net/problem=1
//
// "Find the sum of the multiples of 3 and 5 that are less than 1000."
//
// Generally: Find the sum of the multiples of the integer elements of the
// sorted list N that are less than y.

type SumOfMultiples struct{}

type SumOfMultiplesInput struct {
	N []int
	y int
}

func (s SumOfMultiples) description(N, y string) string {
	return fmt.Sprintf(
		"Find the sum of the multiples of the integer elements of the sorted list %s that are less than %s.",
		N, y,
	)
}

func (s SumOfMultiples) PrintHelp() {
	fmt.Println(s.description("N", "y"))
	fmt.Println("Usage: ./project_euler 1 N[0] N[1] ... N[-1] y")
}

func (s SumOfMultiples) DefaultInput() interface{} {
	return SumOfMultiplesInput{
		N: []int{3, 5},
		y: 1000,
	}
}

func (s SumOfMultiples) Describe(in interface{}) string {
	if in, ok := in.(SumOfMultiplesInput); ok {
		return s.description(fmt.Sprintf("%d", in.N), fmt.Sprintf("%d", in.y))
	} else {
		panic("Input type assertion failed")
	}
}

func (s SumOfMultiples) Process(args []int) interface{} {
	if len(args) == 0 {
		return s.DefaultInput()
	}

	return SumOfMultiplesInput{
		N: args[:len(args)-1],
		y: args[len(args)-1],
	}
}

func (s SumOfMultiples) Solve(in interface{}) int {
	if in, ok := in.(SumOfMultiplesInput); ok {
		x := 0
		for i := in.N[0]; i < in.y; i++ {
			for _, n := range in.N {
				if i%n == 0 {
					x += i
					break
				}
			}
		}
		return x
	} else {
		panic("Input type assertion failed")
	}
}

// ======================================================================================================
// Find the sum of the multiples of the integer elements of the sorted list [3 5] that are less than 1000
// Answer: 233168
// Process Duration: 624ns
// Solve Duration:   18.739µs
// Total Duration:   19.363µs
// ======================================================================================================
