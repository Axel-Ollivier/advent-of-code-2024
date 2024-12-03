package two

import (
	"advent-of-code-2024/utils"
	"fmt"
	"strconv"
	"strings"
)

const (
	url string = "https://adventofcode.com/2024/day/2/input"
)

func PartOne() {
	input, err := utils.GetInput(url)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(input), "\n")
	safeCount := 0

	for _, line := range lines {
		levels := parseLevelspartOne(line)
		if isSafepartOne(levels) {
			safeCount++
		}
	}

	fmt.Println("Number of safe reports:", safeCount)
}

func parseLevelspartOne(line string) []int {
	parts := strings.Fields(line)
	levels := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			panic(fmt.Sprintf("Invalid number in report: %s", line))
		}
		levels[i] = num
	}
	return levels
}

func isSafepartOne(levels []int) bool {
	if len(levels) < 2 {
		return false
	}

	direction := 0

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}

		if direction == 0 {
			if diff > 0 {
				direction = 1
			} else {
				direction = -1
			}
		} else if (direction == 1 && diff < 0) || (direction == -1 && diff > 0) {
			return false
		}
	}

	return true
}
