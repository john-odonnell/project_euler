package solutions

import "fmt"

// https://projecteuler.net/problem=2
//
// "By considering the terms in the Fibonacci sequence whose values do not
// exceed 4e6, find the sum of the even-valued terms."
//
// Generally: By considering the terms in the Fibonacci sequence whose values do
// not exceed y, find the sum of the terms divisible by x.

type EvenFibonacci struct{}

type EvenFiboncciInput struct {
	y int
	x int
}

func (e EvenFibonacci) description(y, x string) string {
	return fmt.Sprintf(
		"By considering the terms of the Fibonacci sequence whose values do not exceed %s, find the sum of the terms divisble by %s",
		y, x,
	)
}

func (e EvenFibonacci) PrintHelp() {
	fmt.Println(e.description("y", "x"))
	fmt.Println("Usage: ./project_euler 2 x y")
}

func (e EvenFibonacci) DefaultInput() interface{} {
	return EvenFiboncciInput{
		y: 4000000,
		x: 2,
	}
}

func (e EvenFibonacci) Describe(in interface{}) string {
	if in, ok := in.(EvenFiboncciInput); ok {
		return fmt.Sprintf(e.description(fmt.Sprintf("%d", in.y), fmt.Sprintf("%d", in.x)))
	} else {
		panic("Input type assertion failed")
	}
}

func (s EvenFibonacci) Process(args []int) interface{} {
	if len(args) == 0 {
		return s.DefaultInput()
	}

	return EvenFiboncciInput{
		x: args[0],
		y: args[1],
	}
}

func (s EvenFibonacci) Solve(in interface{}) int {
	if in, ok := in.(EvenFiboncciInput); ok {
		n1 := 1
		n2 := 1
		next := 0
		total := 0
		for next < in.y {
			next = n2 + n1
			if next%in.x == 0 {
				total += next
			}
			n1 = n2
			n2 = next
		}
		return total
	} else {
		panic("Input type assertion failed")
	}
}

// =========================================================================================
// Find the sum of the terms of the Fibonacci sequence below 4000000 that are divisible by 2
// Answer: 4613732
// Process Duration: 498ns
// Solve Duration:   480ns
// Total Duration:   978ns
// =========================================================================================
