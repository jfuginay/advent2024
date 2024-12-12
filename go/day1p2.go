package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Define the path to your input file
	inputFilePath := "../inputs/dayOneInput.txt"

	// Open the file
	file, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Initialize slices for the two lists
	var leftList []int
	var rightList []int

	// Read the input file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Invalid input format: each line must contain exactly two numbers")
			return
		}

		// Convert the numbers from strings to integers
		leftNum, err1 := strconv.Atoi(parts[0])
		rightNum, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error converting input to integers:", err1, err2)
			return
		}

		// Append the numbers to their respective lists
		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	// Check for errors during file reading
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Create a frequency map for the right list
	rightFreq := make(map[int]int)
	for _, num := range rightList {
		rightFreq[num]++
	}

	// Calculate the total similarity score
	totalScore := 0
	for _, num := range leftList {
		totalScore += num * rightFreq[num]
	}

	// Print the result
	fmt.Printf("The total similarity score is: %d\n", totalScore)
}
