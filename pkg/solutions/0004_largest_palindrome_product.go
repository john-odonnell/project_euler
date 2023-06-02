package solutions

import (
	"fmt"
	"math"
	"reflect"
)

// https://projecteuler.net/problem=4
//
// "Find the largest palindrome made from the product of two 3-digit numbers."
//
// Generally: Find the largest palindrome made from the project of n m digit
// numbers.

type LargestPalindromeProduct struct{}

type LargestPalindromeProductInput struct {
	n int
	m int
}

func (l LargestPalindromeProduct) description(n, m string) string {
	return fmt.Sprintf(
		"Find the largest palindrome made from the product of %s %s digit numbers",
		n, m,
	)
}

func (l LargestPalindromeProduct) PrintHelp() {
	fmt.Println(l.description("n", "m"))
	fmt.Println("This solution is not a general form!")
	fmt.Println("Usage: ./project_euler 4")
}

func (l LargestPalindromeProduct) DefaultInput() interface{} {
	return LargestPalindromeProductInput{
		n: 2,
		m: 3,
	}
}

func (l LargestPalindromeProduct) Describe(in interface{}) string {
	if in, ok := in.(LargestPalindromeProductInput); ok {
		return fmt.Sprintf(l.description(fmt.Sprintf("%d", in.n), fmt.Sprintf("%d", in.m)))
	} else {
		panic("Input type assertion failed")
	}
}

func (l LargestPalindromeProduct) Process(args []int) interface{} {
	fmt.Println("Note: This solution is not a general form!")

	return l.DefaultInput()
}

func IntToRevArray(n int) []int {
	rev_array := []int{}
	for n > 0 {
		rev_array = append(rev_array, n%10)
		n = n / 10
	}
	return rev_array
}

func ReverseList(a []int) []int {
	rev_array := []int{}
	for i := len(a) - 1; i >= 0; i-- {
		rev_array = append(rev_array, a[i])
	}
	return rev_array
}

func (s LargestPalindromeProduct) Solve(in interface{}) int {
	if in, ok := in.(LargestPalindromeProductInput); ok {
		highest_palidrome_product := 0
		for i := int(math.Pow10(in.m - 1)); i < int(math.Pow10(in.m)); i++ {
			for j := int(math.Pow10(in.m - 1)); j < int(math.Pow10(in.m)); j++ {
				product := i * j
				rev_product := IntToRevArray(product)
				for_product := ReverseList(rev_product)
				if product > highest_palidrome_product && reflect.DeepEqual(rev_product, for_product) {
					highest_palidrome_product = product
				}
			}
		}
		return highest_palidrome_product
	} else {
		panic("Input type assertion failed")
	}
}

// ======================================================================
// Find the largest palindrome made from the product of 2 3 digit numbers
// Answer: 906609
// Process Duration: 6.991µs
// Solve Duration:   177.76942ms
// Total Duration:   177.776411ms
// ======================================================================
