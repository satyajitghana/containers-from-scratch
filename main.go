package main

import (
	"os"

	steps "github.com/satyajitghana/container-from-scratch/steps"
)

func main() {
	switch os.Args[1] {
	case "run_one":
		steps.RunOne()
	case "run_two":
		steps.RunTwo()
	case "child":
		steps.RunChild()
	default:
		panic("what? you? doin?")
	}
}
