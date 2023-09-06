package arguments

import "strconv"

func GatherIntegers(args []string) []int {
	var ints []int
	for _, arg := range args {
		i, err := strconv.Atoi(arg)
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}
	return ints
}
