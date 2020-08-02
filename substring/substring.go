package substring

import (
	"errors"
	"strings"
)

var ErrorInvalidIndex = errors.New("index is invalid")

func FindRaw(input string, start string, end string) (string, error) {
	if !strings.Contains(input, start) || !strings.Contains(input, end){
		return "", ErrorInvalidIndex
	}

	s := strings.Index(input, start)
	e := strings.Index(input, end) + len(end)

	return input[s:e], nil
}

func Find(input string, start string, end string) (string, error) {
	result, err := FindRaw(input, start, end)
	if err != nil{
		return "", err
	}

	s := len(start)
	e := len(result) - len(end)

	return result[s:e], err


}