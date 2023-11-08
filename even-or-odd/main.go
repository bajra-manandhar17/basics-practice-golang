package main

import "log"

// In the main function, create a slice of ints from 0 through 10.
// Iterate through the slice with a for loop.
// For each element, check to see if the number is even or odd.
// If the value is even, print out "even". If it is odd, print out "odd"
func main() {
	var numbers []int

	for i := 0; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	for _, number := range numbers {
		if number%2 == 0 {
			log.Printf("%v is even", number)
		} else {
			log.Printf("%v is odd", number)
		}
	}
}
