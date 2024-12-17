package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseReport(line string) []int {
	fields := strings.Fields(line)
	numbers := make([]int, len(fields))
	for i, field := range fields {
		num, _ := strconv.Atoi(field)
		numbers[i] = num
	}
	return numbers
}

func isSafeReport(report []int) bool {
	increasing := true
	decreasing := true

	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		if diff <= 0 || diff > 3 {
			increasing = false
		}
		if diff >= 0 || -diff > 3 {
			decreasing = false
		}
		if !increasing && !decreasing {
			break
		}
	}

	return increasing || decreasing
}

func isSafeWithDampener(report []int) bool {
	if isSafeReport(report) {
		return true
	}

	for i := range report {
		dampened := make([]int, 0, len(report)-1)
		dampened = append(dampened, report[:i]...)
		dampened = append(dampened, report[i+1:]...)

		if isSafeReport(dampened) {
			return true
		}
	}

	return false
}

func solveBothParts(data []string) (int, int) {
	part1Count := 0
	part2Count := 0

	for _, line := range data {
		report := parseReport(line)
		if isSafeReport(report) {
			part1Count++
		}
		if isSafeWithDampener(report) {
			part2Count++
		}
	}

	return part1Count, part2Count
}

func main() {
	// Read from stdin
	scanner := bufio.NewScanner(os.Stdin)
	var data []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		data = append(data, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	part1, part2 := solveBothParts(data)
	fmt.Printf("Part 1: %d safe reports\n", part1)
	fmt.Printf("Part 2: %d safe reports with dampener\n", part2)
}
