package main

import (
	"fmt"
	"os"

	"github.com/bit-cmdr/adventofcode2024-go/pkg/day01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No puzzle specified")
		return
	}

	fmt.Printf("Running %s\n", args[0])
	switch args[0] {
	case "day01":
		fmt.Println("Running puzzle 1")
		day01.Puzzle1()
		fmt.Println("Running puzzle 2")
		day01.Puzzle2()

	default:
		fmt.Println("No known puzzle specified")
	}
}
