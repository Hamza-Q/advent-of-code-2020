package utils

import (
	"bufio"
	"os"

	"github.com/pkg/errors"
)

// Readlines returns a slice of lines from the given
// text file.
func Readlines(filename string) ([]string, error) {
	f, err := os.Open(filename)
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
