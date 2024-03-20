package main

import (
	"fmt"
	"os"
)

const pi = 3.14 // Constant for the value of pi

// circleArea calculates the area of a circle given its radius
func circleArea(r float64) float64 {
	return pi * (r * r)
}

// trapezoidArea calculates the area of a trapezoid given its bases and height
func trapezoidArea(base1, base2, height float64) float64 {
	return 0.5 * (base1 + base2) * height
}

// parallelogramArea calculates the area of a parallelogram given its base and height
func parallelogramArea(base, height float64) float64 {
	return base * height
}

func main() {
	var base, height float64
	var err error

	// Prompt user to enter the base of the parallelogram
	fmt.Print("Enter the base of the parallelogram: ")
	_, err = fmt.Scanln(&base)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Prompt user to enter the height of the parallelogram
	fmt.Print("Enter the height of the parallelogram: ")
	_, err = fmt.Scanln(&height)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	areaCircle := circleArea(float64(base))
	fmt.Printf("The area of the circle with radius %.2f is %.2f\n", base, areaCircle)

	base1 := float64(5)
	base2 := float64(7)
	heightTrapezoid := float64(4)
	areaTrapezoid := trapezoidArea(base1, base2, heightTrapezoid)
	fmt.Printf("The area of the trapezoid with bases %.2f and %.2f and height %.2f is %.2f\n", base1, base2, heightTrapezoid, areaTrapezoid)

	areaParallelogram := parallelogramArea(base, height)
	fmt.Printf("The area of the parallelogram with base %.2f and height %.2f is %.2f\n", base, height, areaParallelogram)
}
