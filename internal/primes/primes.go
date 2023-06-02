package primes

import (
	"fmt"
	"math"
)

// PrimesBelowLimitFunc returns a list of prime numbers that are less than the
// provided upper limit n.
type PrimesBelowLimitFunc func(n int) []int

// SieveOfEratosthenese returns a list of prime numbers Z that are less than the
// provided upper limit n.
func SieveOfEratosthenes(n int) []int {
	fmt.Printf("Finding primes under %d...\n", n)

	if n < 2 {
		return []int{}
	}

	// A indicates the primeness of the index, with 0 indicating prime.
	A := make([]int, n)
	A[0] = 1
	A[1] = 1

	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if A[i] == 0 {
			j := i * i
			for j < n {
				A[j] = 1
				j = j + i
			}
		}
	}

	Z := []int{}
	for i, a := range A {
		if a == 0 {
			Z = append(Z, i)
		}
	}

	return Z
}

// Check calculates if a given integer is prime.
func Check(n int) bool {
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}

// FirstN returns a list of the first n prime integers.
func FirstN(n int) []int {
	Z := make([]int, n)
	Z[0] = 2

	i := 1
	j := 3
	for i < n {
		if Check(j) {
			Z[i] = j
			i += 1
		}
		j += 1
	}

	return Z
}
