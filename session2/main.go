package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// calculateMean calculates the mean (average) of an array of float64 numbers
func calculateMean(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

// calculateStandardDeviation calculates the standard deviation of an array of float64 numbers
func calculateStandardDeviation(numbers []float64) float64 {
	mean := calculateMean(numbers)
	variance := 0.0
	for _, num := range numbers {
		variance += math.Pow(num-mean, 2)
	}
	variance /= float64(len(numbers))
	return math.Sqrt(variance)
}

func main() {
	var numbers []float64

	// Create a buffered reader for user input
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user to enter numbers
	fmt.Println("Enter comma-separated numbers (enter 'done' to finish):")

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}

		input = input[:len(input)-1] // Remove trailing newline

		if input == "done" {
			break
		}

		// Split the input string by commas
		stringNumbers := strings.Split(input, ",")

		for _, strNum := range stringNumbers {
			num, err := strconv.ParseFloat(strings.TrimSpace(strNum), 64)
			if err != nil {
				fmt.Println("Invalid number:", err)
				continue
			}

			numbers = append(numbers, num)
		}
	}
	// Calculate standard deviation
	standardDeviation := calculateStandardDeviation(numbers)
	fmt.Printf("The standard deviation of the numbers is: %.2f\n", standardDeviation)
}
