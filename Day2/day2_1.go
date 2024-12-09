package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, data_err := os.Open("day2_input.txt")

	if data_err != nil {
		log.Fatal(data_err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	var numbers []int
	increase_count := 0
	// decrease_count := 0
	var failed_results [][]int

	for scanner.Scan() {
		line := scanner.Text()
		split_line := strings.Split(line, " ")
		numbers = []int{}

		for i := range split_line {
			convert_me, _ := strconv.ParseInt(split_line[i], 10, 64)
			numbers = append(numbers, int(convert_me))
		}
		i_result, i_fail_return := increase_validation(numbers)
		if i_result == true {
			increase_count++
		} else {
			failed_results = append(failed_results, i_fail_return)
		}

	// 	d_result, d_fail_return := decrease_validation(numbers)
	// 	if d_result == true {
	// 		decrease_count++
	// 	} else {
	// 		failed_results = append(failed_results, d_fail_return)
	// 	}
	}

	for j := range len(failed_results) {
		j_result, _ := increase_validation(failed_results[j])
		if j_result == true {
			increase_count++
		}
	}
	final_count := increase_count
	fmt.Printf("Number of safe reports: %d\n", final_count)

}

// Check if the numbers in the array pass the Day 2 Part 1 test.
// Must increase or decrease all the way through, no change is considered a failure
// Change can be no more than 3
// false == unsafe | true == safe
func increase_validation(numbers []int) (bool, []int) {
	iterator := 1
	success_return := []int{}

	for i := range len(numbers) - 1 {
		test := numbers[iterator] - numbers[i]
		if test > 0 && test <= 3 {
			iterator++
		} else {
			var remove_error []int
			if numbers[iterator] > numbers[i] {
				remove_error = remove_value(numbers, iterator)
			} else {
				remove_error = remove_value(numbers, i)
			}
			return false, remove_error
		}
	}
	return true, success_return
}

func decrease_validation(numbers []int) (bool, []int) {
	iterator := 1
	success_return := []int{}

	for i := range len(numbers) - 1 {
		test := numbers[i] - numbers[iterator]
		if test > 0 && test <= 3 {
			iterator++
		} else {
			var remove_error []int
			fmt.Printf("Original: %v\niterator: %d\n", numbers, iterator)
			if numbers[iterator] > numbers[i] {
				remove_error = remove_value(numbers, iterator)
			} else {
				remove_error = remove_value(numbers, i)
			}
			
			fmt.Printf("Mod: %v\n", remove_error)
			return false, remove_error
		}
	}
	return true, success_return
}

func remove_value(numbers []int, s int) []int {
	return append(numbers[:s], numbers[s+1:]...)
}
