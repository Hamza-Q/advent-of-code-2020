package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	passwords, err := loadInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	numValid, err := getNumValidPasswords(passwords)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", numValid)
}

func loadInput() ([]password, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	passwords := make([]password, 2000)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		password, passErr := passwordFromLine(scanner.Text())
		if passErr != nil {
			return nil, errors.Wrap(passErr, "Bad input")
		}
		passwords = append(passwords, password)
	}
	if err = scanner.Err(); err != nil {
		return nil, errors.Wrap(err, "error reading input.txt")
	}
	return passwords, nil
}

func passwordFromLine(line string) (password, error) {
	vals := strings.Split(line, " ")
	policy, letter, passwordVal := vals[0], vals[1], vals[2]

	minCount, maxCount := parsePolicy(policy)

	return password{
		value:    passwordVal,
		letter:   string(letter[0]),
		minCount: minCount,
		maxCount: maxCount,
	}, nil
}

func parsePolicy(policy string) (minCount, maxCount int) {
	split := strings.Split(policy, "-")
	minCountStr, maxCountStr := split[0], split[1]

	minCount, _ = strconv.Atoi(minCountStr)
	maxCount, _ = strconv.Atoi(maxCountStr)

	return minCount, maxCount
}

func getNumValidPasswords(passwords []password) (int, error) {
	numValid := 0
	for _, pass := range passwords {
		if pass.validate() {
			numValid++
		}
	}
	return numValid, nil
}

type password struct {
	minCount int
	maxCount int
	letter   string
	// The actual password
	value string
}

func (p password) validate() bool {
	count := strings.Count(p.value, p.letter)
	return count >= p.minCount && count <= p.maxCount
}
