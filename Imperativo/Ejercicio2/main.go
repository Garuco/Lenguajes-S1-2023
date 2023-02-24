package main

import "fmt"

// Function to print the shape.
// It receives a positive odd number that indicates the amount of elements in the center of the shape.
func printShape(elements int) {

	// Loop that prints the first half and middle of the shape.
	// "i" will also work as a counter of * that needs to be print for each call of the loop.
	for i := 1; i <= elements; i += 2 {
		//Loop to print/calculate the spaces before printing an *. First time will tell us where the center is.
		//Quantity of spaces will decrease by one each time we get to this loop.
		for j := 1; j <= (elements-i)/2; j++ {
			fmt.Print("   ") //2 extra spaces needed for the shape to look exactly as the picture.
		}
		//Loop to print the quantity of * that we need per row.
		//Quantity of * will increase by 2 each time we get to this loop.
		for j := 1; j <= i; j++ {
			fmt.Print(" * ") //1 extra space before and after the * to get the exact shape that we are looking for.
		}
		fmt.Println()
	}

	// Now we need to do the process backwards, from the maximum amount of * to the less amount,
	// and the minimum amount of spaces to the maximum amount.
	for i := elements - 2; i >= 1; i -= 2 {
		for j := 1; j <= (elements-i)/2; j++ {
			fmt.Print("   ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print(" * ")
		}
		fmt.Println()
	}
}

func main() {
	middleElements := 5 // Positive odd number that represents the quantity of elements in the middle of the shape.
	printShape(middleElements)
}
