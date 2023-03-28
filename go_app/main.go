// Created by: Alyssia fung
// Created on: Mar 2023
//
// This program finds the area of a circle

package main

import "fmt"

func main() {
	// This function finds the area of a circle
	var radius int
	var area int

	// input
	fmt.Println("This program finds the area of a circle")
	fmt.Println()
	fmt.Print("Enter the radius (cm): ")
	fmt.Scanln(&radius)

	// process
	area = 3.14 * radius**2

	// output
	fmt.Println()
	fmt.Println("The area is:", area, " cm².")
	fmt.Println("\nDone.")
}
