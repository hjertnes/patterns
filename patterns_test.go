package patterns

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testString1 = "This is a [[https://google.com][Link]]"

func TestFindRaw(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		expectedResult := "[[https://google.com][Link]]"

		resultString, err := FindRaw(testString1, "[[", "]]")

		assert.Nil(t, err)
		assert.Equal(t, expectedResult, resultString)
	})

	t.Run("Start not found", func(t *testing.T) {
		testString := "This is a ]]"

		_, err := FindRaw(testString, "[[", "]]")

		assert.Error(t, err)
		assert.True(t, errors.Is(err, ErrInvalidMatch))
	})

	t.Run("End not found", func(t *testing.T) {
		testString := "This is a [["

		_, err := FindRaw(testString, "[[", "]]")

		assert.Error(t, err)
		assert.True(t, errors.Is(err, ErrInvalidMatch))
	})
}

func TestFind(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		expectedResult := "https://google.com][Link"

		resultString, err := Find(testString1, "[[", "]]")

		assert.Nil(t, err)
		assert.Equal(t, expectedResult, resultString)
	})

	t.Run("Start not found", func(t *testing.T) {
		testString := "This is a ]]"

		_, err := Find(testString, "[[", "]]")

		assert.Error(t, err)
		assert.True(t, errors.Is(err, ErrInvalidMatch))
	})

	t.Run("End not found", func(t *testing.T) {
		testString := "This is a [["

		_, err := Find(testString, "[[", "]]")

		assert.Error(t, err)
		assert.True(t, errors.Is(err, ErrInvalidMatch))
	})
}

func TestFindAndSplit(t *testing.T) {
	t.Run("Org just url link", func(t *testing.T) {
		result, err := FindAndSplit(testString1, "[[", "]]", "][")

		assert.Nil(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, "https://google.com", result[0])
		assert.Equal(t, "Link", result[1])
	})

	t.Run("Full org link", func(t *testing.T) {
		testString := "This is a [[https://google.com]]"

		result, err := FindAndSplit(testString, "[[", "]]", "][")

		assert.Nil(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, "https://google.com", result[0])
	})
}