// Package patterns contains various methods for pattern matching
package patterns

import (
	"errors"
	"strings"
)

// ErrInvalidMatch error for when the supplied start or end strings isn't in the input string.
var ErrInvalidMatch = errors.New("match is invalid")

// FindRaw Returns a substring matching the supplied arguments or error; also contains the matched characters.
func FindRaw(input string, start string, end string) (string, error) {
	if !strings.Contains(input, start) || !strings.Contains(input, end) {
		return "", ErrInvalidMatch
	}

	s := strings.Index(input, start)

	input = input[s:]

	e := strings.Index(input, end) + len(end)

	return input[:e], nil
}

// Find Returns a substring matching the supplied arguments or error; doesn't contains the matched characters.
func Find(input string, start string, end string) (string, error) {
	result, err := FindRaw(input, start, end)
	if err != nil {
		return "", err
	}

	s := len(start)
	e := len(result) - len(end)

	return result[s:e], err
}

// FindAndSplit Returns a splitted substring based on params.
func FindAndSplit(input, start, end, split string) ([]string, error) {
	result, err := Find(input, start, end)
	if err != nil {
		return make([]string, 0), err
	}

	return strings.Split(result, split), nil
}
