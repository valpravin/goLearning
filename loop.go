package main

import (
	"fmt"
)

var toBeSorted [10]int = [10]int{1, 3, 2, 4, 8, 6, 7, 2, 3, 0}

func bubbleSort(input [10]int) {
	n := 10
	swapped := true
	for swapped {
		fmt.Print("hello loop")
		swapped = false
		for i := 1; i < n; i++ {
			// if the current element is greater than the next
			// element, swap them
			if input[i-1] > input[i] {
				// log that we are swapping values for posterity
				fmt.Println("Swapping")
				// swap values using Go's tuple assignment
				input[i], input[i-1] = input[i-1], input[i]
				// set swapped to true - this is important
				// if the loop ends and swapped is still equal
				// to false, our algorithm will assume the list is
				// fully sorted.
				swapped = true
			}
		}
	}
	fmt.Print(input)
}

func main() {
	fmt.Print("hello")
	bubbleSort(toBeSorted)
}
