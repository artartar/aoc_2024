package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	open_file, input_err := os.Open("day3_input.txt")
	err_check(input_err)	
	defer open_file.Close()

	scanner := bufio.NewScanner(open_file)

	read := make([]string, 1)

	for scanner.Scan() {
		read = append(read, scanner.Text())
	}

	fmt.Print(read)
}

func err_check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}