package conversions

import "strconv"

// StoIA converts a string of digits into an array of integers.
func StoIA(s string) []int {
	A := make([]int, len(s))
	for i, c := range s {
		A[i] = int(c - '0')
	}
	return A
}

// IAtoS converts an array of integers into a string of digits.
func IAtoS(N []int) string {
	S := ""
	for _, n := range N {
		S = S + strconv.Itoa(n)
	}
	return S
}

// ItoIA converts an integer into an array of its digits.
func ItoIA(n int) []int {
	N := []int{}
	for n > 0 {
		last := n % 10
		N = append([]int{last}, N...)
		n = n / 10
	}
	return N
}
