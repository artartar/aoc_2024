package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {
	data, data_err := os.Open("day1_input.txt")

	if data_err != nil {
		log.Fatal(data_err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		my_line := scanner.Text()
		
		break
	}
}
