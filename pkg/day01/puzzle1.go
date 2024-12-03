package day01

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return i
}

func puzzleInput(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %s", err)
	}
	return lines
}

func Puzzle1() {
	inputLines := puzzleInput("pkg/day01/input.txt")

	var left []int
	var right []int
	for _, line := range inputLines {
		segments := strings.Split(line, "   ")
		if len(segments) != 2 {
			continue
		}
		left = append(left, toInt(segments[0]))
		right = append(right, toInt(segments[1]))
	}

	sort.Ints(left)
	sort.Ints(right)

	var distances []int
	for i := 0; i < len(left); i++ {
		if i >= len(right) {
			break
		}
		distance := int(math.Abs(float64(left[i] - right[i])))
		distances = append(distances, distance)
	}

	sum := 0
	for _, distance := range distances {
		sum += distance
	}
	fmt.Println(sum)
}
