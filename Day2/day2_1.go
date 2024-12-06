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
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		split_line := strings.Split(line, " ")
		numbers = []int{}

		for i := range split_line {
			convert_me, _ := strconv.ParseInt(split_line[i], 10, 64)
			numbers = append(numbers, int(convert_me))
		}
		if is_it_safe(numbers) == true {
			fmt.Println(numbers)
			count++
		}
	}
	fmt.Printf("Number of safe reports: %d", count)
}

// Check if the numbers in the array pass the Day 2 Part 1 test.
// Must increase or decrease all the way through, no change is considered a failure
// Change can be no more than 3
// false == unsafe | true == safe
func is_it_safe(numbers []int) bool {
	iterator := 1
	for i := range(len(numbers) - 1) {
		test := numbers[iterator] - numbers[i]
		if test > 0 && test <= 3 {
			iterator++
		} else if test == 0 || test > 3 {
			iterator = 1
			return false
		} else if test < 0 && test > -4 {
			iterator++
		} else if test < -3 {
			iterator = 1
			return false
		}
	}
	return true
}