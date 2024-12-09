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
	fmt.Printf("Advent of Code - Day 9!\n\n")
	open_me, open_err := os.Open("day9_input.txt")
	err_check(open_err)
	defer open_me.Close()

	scanner := bufio.NewScanner(open_me)
	var numbers []string
	var block []string
	block_index := make(map[int]string)
	count := 0

	// scan file opened earlier in the file
	// split string and append to a slice
	for scanner.Scan() {
		line := scanner.Text()
		split_line := strings.Split(line, "")

		for i := range split_line {
			numbers = append(numbers, split_line[i])
		}
	}

	// iterate through slice and convert to format described in the Advent of Code 2024 day 9. 
	// Including zero, even numbered items in the slice are file blocks. File ID's start at zero, and continue sequentially as the file appeared prior to any moves. Does not count any wwhite space
	// File blocks represented by the ID repeated the number of times equal to the size of file block.
	// Odd numbered items are free space, represented by a .
	// . symbol repeated a number of times equal to the size of free space
	for i := range numbers {
		if i == 0 || i%2 == 0 {
			block = append(block, strings.Repeat(convert_int_to_string(count), convert_string_to_int(numbers[i])))
			block_index[i] = convert_int_to_string(count)
			count++
		} else if i%2 != 0 {
			block = append(block, strings.Repeat(string('.'), convert_string_to_int(numbers[i])))
		}
	}

	// for loop - iterate through block
		// if index[x][0] == 46 // 46 is the code for .
		// then, len index[x]
		// for loop through block, start at the end
			// once a number is found, create a function that loops through the string an ddetermines how many times the number repeats. Remember to use the map. If less than the result of the function index[x] save this x
			// continue loop until enough numbers are found
			// save to a new slice in order, replace those with white space.
			// replace the white space using the outer loops x
}

func err_check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func convert_string_to_int(string_number string) int {
	converted, _ := strconv.Atoi(string_number)

	return converted
}

func convert_int_to_string(number int) string {
	converted := strconv.Itoa(number)

	return converted
}
