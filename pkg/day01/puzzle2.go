package day01

import (
	"fmt"
	"sort"
	"strings"
)

func Puzzle2() {
	inputLines := puzzleInput("pkg/day01/input.txt")

	var left, right []int
	for _, line := range inputLines {
		segments := strings.Split(line, "   ")
		if len(segments) > 0 && segments[0] != "" {
			left = append(left, toInt(segments[0]))
		}
		if len(segments) > 1 && segments[1] != "" {
			right = append(right, toInt(segments[1]))
		}
	}

	sort.Ints(left)
	sort.Ints(right)

	hash := make(map[int]int)
	for _, r := range right {
		hash[r]++
	}

	var similarities []int
	for _, l := range left {
		similarity := l * hash[l]
		similarities = append(similarities, similarity)
	}

	sum := 0
	for _, similarity := range similarities {
		sum += similarity
	}

	fmt.Println(sum)
}
