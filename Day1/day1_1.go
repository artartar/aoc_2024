package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// Open file
	// create scanner
	// for scanner.Scan and scanner.Text to read line by line
	// data type is string, line example: 18102   93258
	// take line and gather the first full number, i.e 18102
	// convert from string to int
	// append to slice1
	// ignore spacing and collect second number
	// repeat convert and append, but append to slice2
	// repeat over the entirety of the text file.

	test_slice1 := []string{"18102", "93258"}
	num_slice1 := make([]int, 2)

	for i := range 2 {
		j, con_err := strconv.Atoi(test_slice1[i])
		err_check(con_err)
		// num_slice1 = append(num_slice1, j)

	}

	fmt.Print(num_slice1)
}

func select_sort_measure(slc1, slc2 []int) int {
	final_distance := 0
	my_range := len(slc1)
	var swap_num int
	var low_num int

	for i := range(my_range) {
		low_num_position := i

		for j := range(my_range) {
			if slc1[j] > slc1[low_num_position] {
				low_num = slc1[low_num_position]
				swap_num = slc1[j]
				slc1[j] = low_num
				slc1[low_num_position] = swap_num
			}
		}
	}

	for i := range(my_range) {
		low_num_position := i

		for j := range(my_range) {
			if slc2[j] > slc2[low_num_position] {
				low_num = slc2[low_num_position]
				swap_num = slc2[j]
				slc2[j] = low_num
				slc2[low_num_position] = swap_num
			}
		}
	}

	for k := range(my_range) {
		result := slc2[k] - slc1[k]
		final_distance += result
	}

	return final_distance
}

func err_check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}