package substring

import (
	"errors"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestFindRaw(t *testing.T){
	t.Run("Happy path", func(t *testing.T){
		testString := "This is a [[https://google.com][Link]]"
		expectedResult := "[[https://google.com][Link]]"

		resultString, err := FindRaw(testString, "[[", "]]")

		assert.Nil(t, err)
		assert.Equal(t, expectedResult, resultString)
	})

	t.Run("Start not found", func(t *testing.T) {
		testString := "This is a ]]"

		_, err := FindRaw(testString, "[[", "]]")

		assert.Error(t, err)
		assert.True(t, errors.Is(err, ErrorInvalidIndex))
	})

	t.Run("End not found", func(t *testing.T) {
		testString := "This is a [["

		_, err := FindRaw(testString, "[[", "]]")

		assert.Error(t, err)
		assert.True(t, errors.Is(err, ErrorInvalidIndex))
	})
}

func TestFind(t *testing.T){
	t.Run("Happy path", func(t *testing.T){
		testString := "This is a [[https://google.com][Link]]"
		expectedResult := "https://google.com][Link"

		resultString, err := Find(testString, "[[", "]]")

		assert.Nil(t, err)
		assert.Equal(t, expectedResult, resultString)
	})

	t.Run("Start not found", func(t *testing.T) {
		testString := "This is a ]]"

		_, err := Find(testString, "[[", "]]")

		assert.Error(t, err)
		assert.True(t, errors.Is(err, ErrorInvalidIndex))
	})

	t.Run("End not found", func(t *testing.T) {
		testString := "This is a [["

		_, err := Find(testString, "[[", "]]")

		assert.Error(t, err)
		assert.True(t, errors.Is(err, ErrorInvalidIndex))
	})
}
