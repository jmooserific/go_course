package main

import "fmt"

func main() {
	numbers := []int{}

	// Make a slice of the integers from 1 to 10
	for i := 1; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	// Iterate over the numbers, determining if each is odd or even.
	for _, value := range numbers {
		if value%2 == 0 {
			fmt.Println(value, "is even.")
		} else {
			fmt.Println(value, "is odd.")
		}
	}
}
