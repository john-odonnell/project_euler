package solutions

import "time"

// ProcessFunc is a type that represents a function that consumes a list of
// integer inputs and converts it into a problem-specific input struct.
type ProcessFunc func([]int) interface{}

// ProcessFunc is a type that represents a function that consumes a list of
// integer inputs and converts it into a problem-specific input struct and
// captures the computation time.
type TimedProcessFunc func([]int) (interface{}, time.Duration)

// SolveFunc is a type that represents a function that can compute the solution
// for a particular Project Euler problem, assuming that the provided empty
// interface is an instance of a problem-specific input struct.
type SolveFunc func(interface{}) int

// TimedSolveFunc is a type that represents a function that can compute the
// solution for a particular Project Euler problem and capture the computation
// time.
type TimedSolveFunc func(interface{}) (int, time.Duration)

// Strategy is an interface for solving a Project Euler problem.
//   - PrintHelp prints the problem description string and usage information.
//   - DefaultInput returns an instance of a problem-specific input struct with
//     default values (those expected by Project Euler).
//   - Process consumes a list of integer inputs and converts it into a
//     problem-specific input struct.
//   - Describe consumes a problem-specific input and returns a string
//     describing the problem and/or solution.
//   - Solve consumes a problem-specific input and computes a solution.
type Strategy interface {
	PrintHelp()
	DefaultInput() interface{}
	Process([]int) interface{}
	Describe(interface{}) string
	Solve(interface{}) int
}

// SolveWithDuration converts a provided SolveFunc into a TimedSolveFunc.
func SolveWithDuration(f SolveFunc) TimedSolveFunc {
	return func(in interface{}) (int, time.Duration) {
		start := time.Now()
		answer := f(in)
		duration := time.Since(start)
		return answer, duration
	}
}

// ProcessWithDuration convers a provided ProcessFunc into a TimedProcessFunc.
func ProcessWithDuration(f ProcessFunc) TimedProcessFunc {
	return func(in []int) (interface{}, time.Duration) {
		start := time.Now()
		answer := f(in)
		duration := time.Since(start)
		return answer, duration
	}
}

// Map relates problem numbers to Strategy implementations.
var Map = map[string]Strategy{
	"1":  SumOfMultiples{},
	"2":  EvenFibonacci{},
	"3":  LargestPrimeFactor{},
	"4":  LargestPalindromeProduct{},
	"5":  SmallestMultiple{},
	"6":  SumSquareDifference{},
	"7":  NthPrime{},
	"8":  LargestProductInSeries{},
	"9":  SpecialPythagoreanTriplet{},
	"10": SummationOfPrimes{},
	"11": LargestProductGrid{},
}
