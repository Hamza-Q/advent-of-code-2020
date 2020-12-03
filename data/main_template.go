/*
	A templated main.go file to use for a new day of advent of code.
*/

package main

import (
	"fmt"

	"github.com/Hamza-Q/advent-of-code-2020/utils"
)

func main() {
	lines, err := utils.Readlines("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("input.txt information:")
	fmt.Printf("\tNumber of lines:\t%d\n", len(lines))
	fmt.Printf("\tLength of first line:\t%d\n", len(lines[0]))
}
