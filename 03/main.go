package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

const lineLen = 31

func main() {
	lines, err := loadInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	numTrees := getNumTrees(lines, 3, 1)
	fmt.Println("part1:", numTrees)

	product := 1
	for _, c := range []struct {
		right, down int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	} {
		product *= getNumTrees(lines, c.right, c.down)
	}
	fmt.Println("part2:", product)
}

func loadInput() ([]string, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		return nil, errors.Wrap(err, "error reading input.txt")
	}
	return lines, nil
}

func getNumTrees(lines []string, right int, down int) int {
	num := 0
	xInd := 0
	for lineNum, line := range lines {
		if lineNum == 0 {
			continue
		}
		if down != 1 {
			if lineNum%2 != 0 {
				continue
			}
		}
		xInd += right
		thing := string(line[xInd%lineLen])
		if thing == "#" {
			num++
		}
	}
	return num
}
