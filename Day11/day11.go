package main

import (
	"fmt"
	"strconv"
)

func main() {
	day11_input := []int{3279, 998884, 1832781, 517, 8, 18864, 28, 0}
	var output []int

	blinks := 75
	// Outer for loop to determine the number of blinks
	for i := 0; i < blinks; i++ {
		output = []int{}
		for i := range day11_input {
			as_string := convert_to_string(day11_input[i])

			if day11_input[i] == 0 {
				output = append(output, 1)
			} else if len(as_string) % 2 == 0 {
				// Append the left side
				output = append(output, convert_to_int(as_string[0:len(as_string)/2]))

				// Append the right side
				output = append(output, convert_to_int(as_string[len(as_string)/2:]))
			} else {
				output = append(output, day11_input[i] * 2024)
			}
		}
		day11_input = output
	}

	fmt.Printf("Total stones is %d", len(output))
}

func convert_to_string(number int) string {
	return strconv.Itoa(number)
}

func convert_to_int(word string) int {
	converted, _ := strconv.Atoi(word)
	return converted
}