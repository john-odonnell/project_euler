package solutions

import (
	"fmt"

	"github.com/john-odonnell/project_euler/internal/primes"
)

// https://projecteuler.net/problem=10
//
// "Find the sum of all the primes below two million."
//
// Generally: Find the sum of all the primes below n.

type SummationOfPrimes struct{}

type SummationOfPrimesInput struct {
	n int
}

func (s SummationOfPrimes) description(n string) string {
	return fmt.Sprintf(
		"Find the sum of all the primes below %s",
		n,
	)
}

func (s SummationOfPrimes) PrintHelp() {
	fmt.Println(s.description("n"))
	fmt.Println("Usage: ./project_euler 10 n")
}

func (s SummationOfPrimes) DefaultInput() interface{} {
	return SummationOfPrimesInput{
		n: 2000000,
	}
}

func (s SummationOfPrimes) Describe(in interface{}) string {
	if in, ok := in.(SummationOfPrimesInput); ok {
		return fmt.Sprintf(s.description(fmt.Sprintf("%d", in.n)))
	} else {
		panic("Input type assertion failed")
	}
}

func (s SummationOfPrimes) Process(args []int) interface{} {
	if len(args) == 0 {
		return s.DefaultInput()
	}

	return SummationOfPrimesInput{
		n: args[0],
	}
}

func (s SummationOfPrimes) Solve(in interface{}) int {
	if in, ok := in.(SummationOfPrimesInput); ok {
		P := primes.SieveOfEratosthenes(in.n)

		sum := 0
		for _, p := range P {
			sum += p
		}

		return sum
	} else {
		panic("Input type assertion failed")
	}
}

// ============================================
// Find the sum of all the primes below 2000000
// Answer:           142913828922
// Process Duration: 650ns
// Solve Duration:   28.951777ms
// Total Duration:   28.952427ms
// ============================================
