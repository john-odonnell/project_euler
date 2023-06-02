package solutions

import (
	"fmt"
	"math"

	"github.com/john-odonnell/project_euler/internal/primes"
)

// https://projecteuler.net/problem=3
//
// "What is the largest prime factor of the number 600851475143?"
//
// Generally: Find the largest prime factor of the positive integer n.

type LargestPrimeFactor struct{}

type LargestPrimeFactorInput struct {
	n int
}

func (l LargestPrimeFactor) description(n string) string {
	return fmt.Sprintf(
		"Find the largest prime factor of the positive integer %s", n,
	)
}

func (l LargestPrimeFactor) PrintHelp() {
	fmt.Println(l.description("n"))
	fmt.Println("Usage: ./project_euler 3 n")
}

func (l LargestPrimeFactor) DefaultInput() interface{} {
	return LargestPrimeFactorInput{
		n: 600851475143,
	}
}

func (l LargestPrimeFactor) Describe(in interface{}) string {
	if in, ok := in.(LargestPrimeFactorInput); ok {
		return fmt.Sprintf(l.description(fmt.Sprintf("%d", in.n)))
	} else {
		panic("Input type assertion failed")
	}
}

func (l LargestPrimeFactor) Process(args []int) interface{} {
	if len(args) == 0 {
		return l.DefaultInput()
	}

	return LargestPrimeFactorInput{
		n: args[0],
	}
}

func (l LargestPrimeFactor) Solve(in interface{}) int {
	if in, ok := in.(LargestPrimeFactorInput); ok {
		primes := primes.SieveOfEratosthenes(
			int(math.Ceil(
				math.Sqrt(float64(in.n)),
			)),
		)
		if len(primes) == 0 {
			return -1
		}

		largestFactor := -1
		for _, prime := range primes {
			if (in.n % prime) == 0 {
				largestFactor = prime
			}
		}

		return largestFactor
	} else {
		panic("Input type assertion failed")
	}
}

// ========================================================
// Find the largest prime factor of the number 600851475143
// Answer: 6857
// Process Duration: 2.244µs
// Solve Duration:   10.556655ms
// Total Duration:   10.558899ms
// ========================================================
