package solutions

import "fmt"

// https://projecteuler.net/problem=5
//
// "What is the smallest positibe number that is evenly divisible by all the
// numbers from 1 to 20?"
//
// Generally: Find the smallest possible number that is evenly divisible by all
// the numbers from 1 to n?

type SmallestMultiple struct{}

type SmallestMultipleInput struct {
	n int
}

func (s SmallestMultiple) description(n string) string {
	return fmt.Sprintf(
		"Find the smallest possible number that is divisible by integers 1 to %s",
		n,
	)
}

func (s SmallestMultiple) PrintHelp() {
	fmt.Println(s.description("n"))
	fmt.Println("Usage: ./project_euler 5 n")
}

func (s SmallestMultiple) DefaultInput() interface{} {
	return SmallestMultipleInput{
		n: 20,
	}
}

func (s SmallestMultiple) Describe(in interface{}) string {
	if in, ok := in.(SmallestMultipleInput); ok {
		return fmt.Sprintf(s.description(fmt.Sprintf("%d", in.n)))
	} else {
		panic("Input type assertion failed")
	}
}

func (s SmallestMultiple) Process(args []int) interface{} {
	if len(args) == 0 {
		return s.DefaultInput()
	}

	return SmallestMultipleInput{
		n: args[0],
	}
}

func (s SmallestMultiple) Solve(in interface{}) int {
	if in, ok := in.(SmallestMultipleInput); ok {
		i := in.n
		for true {
			passed := 2
			for j := in.n - 1; j > 1; j-- {
				if i%j == 0 {
					passed += 1
				}
			}
			if passed == in.n {
				return i
			} else {
				i += in.n
			}
		}
		return -1
	} else {
		panic("Input type assertion failed")
	}
}

// =======================================================================
// Find the smallest possible number that is divisible by integers 1 to 20
// Answer:           232792560
// Process Duration: 2.078µs
// Solve Duration:   1.346246087s
// Total Duration:   1.346248165s
// =======================================================================
