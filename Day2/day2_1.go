package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	data, data_err := os.Open("day2_input.txt")

	if data_err != nil {
		log.Fatal(data_err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	var my_line [][]string
	for scanner.Scan() {
		iter := 0
		my_line = append(my_line, []string{scanner.Text()})

		var conv_err error
		for i := range(len(my_line[iter])) {
			my_line[iter][i], conv_err = strconv.Atoi(my_line[iter][i])
			if conv_err != nil {
				log.Fatal(conv_err)
			}
		}
	}
}