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
	numValid, err := getNumValidPasswords(passwords, validatePart1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", numValid)

	numValid, err = getNumValidPasswords(passwords, validatePart2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 2:", numValid)
}

func loadInput() ([]password, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	passwords := make([]password, 0, 1000)
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
	vals := strings.Split(line, ":")
	policyStr, passwordVal := vals[0], vals[1]

	data := parsePolicyData(policyStr)

	return password{
		value:  passwordVal,
		policy: data,
	}, nil
}

func parsePolicyData(policyStr string) policyData {
	fmt.Println(policyStr)
	policySplit := strings.Split(policyStr, " ")
	values, letter := policySplit[0], policySplit[1]

	split := strings.Split(values, "-")
	val1Str, val2Str := string(split[0]), string(split[1])

	val1, _ := strconv.Atoi(val1Str)
	val2, _ := strconv.Atoi(val2Str)

	return policyData{val1, val2, letter}
}

type validator func(password) bool

func getNumValidPasswords(passwords []password, validate validator) (int, error) {
	numValid := 0
	for _, pass := range passwords {
		if validate(pass) {
			numValid++
		}
	}
	return numValid, nil
}

type password struct {
	policy policyData
	// The actual password
	value string
}

type policyData struct {
	val1   int
	val2   int
	letter string
}

type passwordPolicy interface {
	validate(password string)
}

type policyPart1 struct {
	policyData
}

func validatePart1(pwd password) bool {
	policy := pwd.policy
	minCount, maxCount := policy.val1, policy.val2
	count := strings.Count(pwd.value, policy.letter)
	return count >= minCount && count <= maxCount
}

func validatePart2(pwd password) bool {
	policy := pwd.policy
	pos1 := pwd.value[policy.val1]
	pos2 := pwd.value[policy.val2]
	if pos1 == pos2 {
		return false
	}
	if string(pos1) != policy.letter && string(pos2) != policy.letter {
		return false
	}
	return true
}
