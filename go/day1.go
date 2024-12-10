package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Helper function to calculate the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

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
		// Split each line into two numbers
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

	// Sort both lists
	sort.Ints(leftList)
	sort.Ints(rightList)

	// Calculate the total distance
	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		totalDistance += abs(leftList[i] - rightList[i])
	}

	// Print the result
	fmt.Printf("The total distance between the two lists is: %d\n", totalDistance)
}
