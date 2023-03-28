// Created by: Alyssia fung
// Created on: Mar 2023
//
// This program finds the area of a circle

package main

import (
	"fmt"
	"math"
)

func main() {
	// Input
	var radius float64
	fmt.Print("Enter the radius of the circle (in cm): ")
	fmt.Scan(&radius)

	// Process
	area := math.Pi * math.Pow(radius, 2)

	// Output
	fmt.Printf("The area of the circle is: %.2f cm²\n", area)
}
