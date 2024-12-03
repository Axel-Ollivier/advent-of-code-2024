package one

import (
	"advent-of-code-2024/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	url string = "https://adventofcode.com/2024/day/1/input"
)

func PartOne() {

	input, err := utils.GetInput(url)
	if err != nil {
		panic(err)
	}

	var left, right []int
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		leftNum, _ := strconv.Atoi(parts[0])
		rightNum, _ := strconv.Atoi(parts[1])
		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	sort.Ints(left)
	sort.Ints(right)

	total := 0
	for i := range left {
		total += utils.Abs(left[i] - right[i])
	}

	fmt.Println("Total Distance:", total)
}

func PartTwo() {

	input, err := utils.GetInput(url)
	if err != nil {
		panic(err)
	}

	var left, right []int
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			panic(fmt.Sprintf("Invalid input line: %s", line))
		}

		leftNum, err1 := strconv.Atoi(parts[0])
		rightNum, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			panic(fmt.Sprintf("Invalid numbers in line: %s", line))
		}

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	rightFrequency := make(map[int]int)
	for _, num := range right {
		rightFrequency[num]++
	}

	similarityScore := 0
	for _, num := range left {
		if count, exists := rightFrequency[num]; exists {
			similarityScore += num * count
		}
	}

	fmt.Println("Similarity Score:", similarityScore)
}
