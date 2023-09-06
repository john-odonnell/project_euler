package solutions

import (
	"fmt"

	"github.com/john-odonnell/project_euler/internal/conversions"
)

// https://projecteuler.net/problem=8
//
// "The four adjacent digits in the 1000-digit number that have the greatest
// product are (9 x 9 x 8 x 9) = 5832. Find the thirteen adjacent digits in the
// 1000-digit number that have the greatest product. What is the value of this
// product?"
//
// Generally: Find the m adjacent digits in number n that have the greatest
// product.

type LargestProductInSeries struct{}

type LargestProductInSeriesInput struct {
	m int
	n []int
}

func (l LargestProductInSeries) description(m, n string) string {
	return fmt.Sprintf(
		"Find the %s adjacent digits in the number %s that have the greatest product",
		m, n,
	)
}

func (l LargestProductInSeries) PrintHelp() {
	fmt.Println(l.description("m", "n"))
	fmt.Println("Usage: ./project_euler 8 m n")
}

func (l LargestProductInSeries) DefaultInput() interface{} {
	return LargestProductInSeriesInput{
		m: 13,
		n: conversions.StoIA("7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450"),
	}
}

func (l LargestProductInSeries) Describe(in interface{}) string {
	if in, ok := in.(LargestProductInSeriesInput); ok {
		return fmt.Sprintf(l.description(fmt.Sprintf("%d", in.m), conversions.IAtoS(in.n)))
	} else {
		panic("Input type assertion failed")
	}
}

func (l LargestProductInSeries) Process(args []int) interface{} {
	if len(args) == 0 {
		return l.DefaultInput()
	}

	return LargestProductInSeriesInput{
		m: args[0],
		n: conversions.ItoIA(args[1]),
	}
}

func (l LargestProductInSeries) Solve(in interface{}) int {
	if in, ok := in.(LargestProductInSeriesInput); ok {
		return solve(in.n, 0, in.m)
	} else {
		panic("Input type assertion failed")
	}
}

func solve(A []int, start, length int) int {
	// If the sliding window passes the end of the provided array, return 0.
	next := start + length
	if next > len(A) {
		return 0
	}

	// If the sliding window contains a 0 value, slide the window past the 0
	// and call the solve function.
	zeroIndex := find(A[start:next], 0)
	if zeroIndex != -1 {
		return solve(A, start+zeroIndex+1, length)
	}

	// Find the product of the sliding window.
	windowProduct := product(A[start:next])
	largestProduct := windowProduct

	// While the upcoming value is not 0 and the sliding window has not exited
	// the array, find the product of the next window and slide the window.
	for next < len(A) && A[next] != 0 {
		windowProduct = windowProduct * A[next] / A[start]
		if windowProduct > largestProduct {
			largestProduct = windowProduct
		}
		start++
		next++
	}

	// If the sliding window has not exited the array but the upcoming value is
	// a 0, slide the window past the zero and call the solve function.
	if next < len(A) {
		futureProduct := solve(A, next, length)
		if futureProduct > largestProduct {
			return futureProduct
		}
	}

	return largestProduct
}

func find(N []int, m int) int {
	for i, n := range N {
		if n == m {
			return i
		}
	}
	return -1
}

func product(N []int) int {
	S := 1
	for _, n := range N {
		if n == 0 {
			return 0
		}
		S = S * n
	}
	return S
}

// ==================================================================================
// Find the 13 adjacent digits in the number 731...450 that have the greatest product
// Answer:           23514624000
// Process Duration: 13.72µs
// Solve Duration:   54.395µs
// Total Duration:   68.115µs
// ==================================================================================
