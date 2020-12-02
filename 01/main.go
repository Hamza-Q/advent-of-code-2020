package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

func main() {
	data, err := loadInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	val1, val2, err := getTwoSumToTarget(data, 2020)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part1:", val1*val2)

	val1, val2, val3, err := getThreeSumToTarget(data, 2020)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Part 2", val1*val2*val3)

}

func loadInput() ([]int, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data := make([]int, 0, 1000)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num, convErr := strconv.Atoi(scanner.Text())
		if convErr != nil {
			return nil, errors.Wrap(convErr, "Bad input")
		}
		data = append(data, num)
	}
	if err = scanner.Err(); err != nil {
		return nil, errors.Wrap(err, "error reading input.txt")
	}
	return data, nil
}

func getTwoSumToTarget(data []int, target int) (int, int, error) {
	knownVals := map[int]bool{}
	for _, val := range data {
		wantedAddend := target - val
		if exists := knownVals[wantedAddend]; exists {
			return val, wantedAddend, nil
		}
		knownVals[val] = true
	}
	return 0, 0, errors.Errorf("No two entries match to target: %d", target)
}

func getThreeSumToTarget(data []int, target int) (int, int, int, error) {
	for _, val := range data {
		wantedTwoSum := target - val
		val1, val2, err := getTwoSumToTarget(data, wantedTwoSum)
		if err == nil {
			return val, val1, val2, nil
		}
	}
	return 0, 0, 0, errors.Errorf("No three entries match to target: %d", target)
}
