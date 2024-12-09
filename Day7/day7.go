package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Day 7 of AoC 2024!\n")	
}

func test_permutations(numbers []int) bool {
	oh_god := len(numbers)**

}
// How to determine the number of possiblities?
// 4^spaces

// 196487536: 7 3 6 4 393 4 1 93 3 7 9 7
// 52253144412: 99 5 9 86 816 52 9 2
// 52182: 5 2 473 1 99 388 2 1
// 83680255: 703 1 5 2 390 2 4 352 8
// 43: 4 1 3
// + + | - - | * + | / + | 
// + - | - + | * - | / - | 
// + * | - * | * * | / * | 
// + / | - / | * / | / / | 
 