package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/john-odonnell/project_euler/pkg/arguments"
	"github.com/john-odonnell/project_euler/pkg/output"
	"github.com/john-odonnell/project_euler/pkg/solutions"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please provide Problem number!\n" +
			"Usage: ./project_euler [problem number] [input...|help]")
		os.Exit(1)
	}

	problem := os.Args[1]
	if problem == "help" {
		fmt.Println("Usage: ./project_euler [problem #] [input...|help]")
		os.Exit(0)
	}

	strategy := solutions.Map[problem]
	if strategy == nil {
		fmt.Printf("Problem %s not implemented!\n", problem)
		os.Exit(1)
	}

	var args []int
	if len(os.Args[2:]) == 0 {
		args = []int{}
	} else if os.Args[2] == "help" {
		strategy.PrintHelp()
		os.Exit(0)
	} else {
		args = arguments.GatherIntegers(os.Args[2:])
	}

	process := solutions.ProcessWithDuration(strategy.Process)
	solve := solutions.SolveWithDuration(strategy.Solve)

	fmt.Println("Processing...")
	in, pDuration := process(args)
	fmt.Println("Solving...")
	answer, sDuration := solve(in)

	output.PrintAnswer(
		strategy.Describe(in),
		strconv.Itoa(answer),
		pDuration, sDuration,
	)
}
