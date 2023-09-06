package solutions

import "fmt"

// https://projecteuler.net/problem=9
//
// "A Pythagorean triplet is a set of three natural numbers, a < b < c, for which
// a^2 + b^2 = c^2. There exists exactly one Pythagorean triplet for which
// a + b + c = 1000. Find the product abc."
//
// No general form.

type SpecialPythagoreanTriplet struct{}

type SpecialPythagoreanTripletInput struct{}

func (s SpecialPythagoreanTriplet) description() string {
	return "Find the product abc of the Pythagorean triplet a^2 + b^2 = c^2 where a + b + c = 1000."
}

func (s SpecialPythagoreanTriplet) PrintHelp() {
	fmt.Println(s.description())
	fmt.Println("Usage: ./project_euler 9")
}

func (s SpecialPythagoreanTriplet) DefaultInput() interface{} {
	return SpecialPythagoreanTripletInput{}
}

func (s SpecialPythagoreanTriplet) Describe(in interface{}) string {
	if _, ok := in.(SpecialPythagoreanTripletInput); ok {
		return s.description()
	} else {
		panic("Input type assertion failed")
	}
}

func (s SpecialPythagoreanTriplet) Process(args []int) interface{} {
	fmt.Println("This problem does not support a general solution.")
	return s.DefaultInput()
}

func (s SpecialPythagoreanTriplet) Solve(in interface{}) int {
	if _, ok := in.(SpecialPythagoreanTripletInput); ok {
		c := 3
		for true {
			for b := 1; b < c; b++ {
				for a := 1; a < b; a++ {
					if (a*a)+(b*b) == (c * c) {
						if a+b+c == 1000 {
							return a * b * c
						}
					}
				}
			}
			c++
		}
		return -1
	} else {
		panic("Input type assertion failed")
	}
}

// =======================================================================================
// Find the product abc of the Pythagorean triplet a^2 + b^2 = c^2 where a + b + c = 1000.
// Answer:           31875000
// Process Duration: 6.008µs
// Solve Duration:   14.342032ms
// Total Duration:   14.34804ms
// =======================================================================================
