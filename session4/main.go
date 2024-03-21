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

func calculateCorrelation(data1, data2 []float64) float64 {
	if len(data1) != len(data2) {
		fmt.Println("Error: Input data sets must have the same length")
		os.Exit(1)
	}

	mean1 := calculateMean(data1)
	mean2 := calculateMean(data2)

	var covariance, stdDev1, stdDev2 float64

	for i := 0; i < len(data1); i++ {
		covariance += (data1[i] - mean1) * (data2[i] - mean2)
	}

	stdDev1 = calculateStandardDeviation(data1)
	stdDev2 = calculateStandardDeviation(data2)

	correlation := covariance / (stdDev1 * stdDev2)
	return correlation
}

func main() {
	var numbers []float64
	// Example data sets
	data1 := []float64{1, 2, 3, 4, 5}
	data2 := []float64{2, 3, 4, 5, 6}

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

	// Calculate correlation coefficient
	correlation := calculateCorrelation(data1, data2)
	fmt.Printf("The correlation coefficient between the two data sets is: %.2f\n", correlation)

}
