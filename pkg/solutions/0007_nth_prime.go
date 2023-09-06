package solutions

import (
	"fmt"

	"github.com/john-odonnell/project_euler/internal/primes"
)

// https://projecteuler.net/problem=7
//
// "What is the 10001st prime number?"
//
// Generally: Find the nth prime number.

type NthPrime struct{}

type NthPrimeInput struct {
	n int
}

func (p NthPrime) description(n string) string {
	return fmt.Sprintf(
		"Find the %sth prime number",
		n,
	)
}

func (p NthPrime) PrintHelp() {
	fmt.Println(p.description("n"))
	fmt.Println("Usage: ./project_euler 7 n")
}

func (p NthPrime) DefaultInput() interface{} {
	return NthPrimeInput{
		n: 10001,
	}
}

func (p NthPrime) Describe(in interface{}) string {
	if in, ok := in.(NthPrimeInput); ok {
		return fmt.Sprintf(p.description(fmt.Sprintf("%d", in.n)))
	} else {
		panic("Input type assertion failed")
	}
}

func (p NthPrime) Process(args []int) interface{} {
	if len(args) == 0 {
		return p.DefaultInput()
	}

	return NthPrimeInput{
		n: args[0],
	}
}

func (p NthPrime) Solve(in interface{}) int {
	if in, ok := in.(NthPrimeInput); ok {
		p := primes.FirstN(in.n)
		return p[in.n-1]
	} else {
		panic("Input type assertion failed")
	}
}

// =============================
// Find the 10001th prime number
// Answer: 104743
// Process Duration: 2.131µs
// Solve Duration:   28.422115ms
// Total Duration:   28.424246ms
// =============================
