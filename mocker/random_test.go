package mocker

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInBetween(t *testing.T) {
	// Arrange
	random := New().Random()

	// Act
	result := random.IntBetween(5, 500)
	nextResult := random.IntBetween(500, 1000)

	// Assert
	assert.Truef(t, result >= 5 && result <= 500, "The number should be in the range of 0 and 10")
	assert.Truef(t, nextResult >= 500 && nextResult <= 1000, "The number should be in the range of 0 and 10")
}

func TestRandomInt(t *testing.T) {
	// Arrange
	random := New().Random()

	// Act
	result := random.RandomInt(10)
	nextResult := random.RandomInt(10)

	// Assert
	assert.Truef(t, result >= 0 && result <= 10, "The number should be in the range of 0 and 10")
	assert.Truef(t, nextResult >= 0 && nextResult <= 10, "The number should be in the range of 0 and 10")
	assert.NotEqualf(t, nextResult, result, "The numbers should not be equal in the next run")
}

func TestRandomStringElement(t *testing.T) {
	// Arrange
	random := New().Random()
	array := [...]string{
		"a",
		"b",
		"c",
		"d",
	}

	// Act
	result := random.RandomStrElement(array[:])
	nextResult := random.RandomStrElement(array[:])

	// Assert
	assert.NotEmptyf(t, result, "String result should not be empty")
	assert.NotEmptyf(t, nextResult, "String nextResult should not be empty")
	assert.Containsf(t, array, result, "Result should be a subset of the array")
	assert.Containsf(t, array, nextResult, "NextResult should be a subset of the array")
}

func TestRandomIntElement(t *testing.T) {
	// Arrange
	random := New().Random()
	array := [...]int{
		1,
		10,
		100,
		1000,
	}

	// Act
	result := random.RandomIntElement(array[:])
	nextResult := random.RandomIntElement(array[:])

	// Assert
	assert.NotEmptyf(t, result, "String result should not be empty")
	assert.NotEmptyf(t, nextResult, "String nextResult should not be empty")
	assert.Containsf(t, array, result, "Result should be a subset of the array")
	assert.Containsf(t, array, nextResult, "NextResult should be a subset of the array")
}

func TestRandomElement(t *testing.T) {
	// Arrange
	strArray := [4]interface{}{
		"a",
		"b",
		"c",
		"d",
	}

	intArray := [4]interface{}{
		1,
		10,
		100,
		1000,
	}

	var tests = []struct {
		collection   []interface{}
		expectedType string
	}{
		{strArray[:], "string"},
		{intArray[:], "int"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%v %v", "Element should be of type", tt.expectedType)
		t.Run(testName, func(t *testing.T) {
			random := New().Random()

			// Act
			result := random.RandomElement(tt.collection[:])
			resultType := reflect.TypeOf(result)
			nextResult := random.RandomElement(tt.collection[:])
			nextResultType := reflect.TypeOf(result)

			// Assert
			assert.NotEmptyf(t, result, "String result should not be empty")
			assert.NotEmptyf(t, nextResult, "String nextResult should not be empty")
			assert.Containsf(t, tt.collection, result, "Result should be a subset of the array")
			assert.Containsf(t, tt.collection, nextResult, "NextResult should be a subset of the array")
			assert.Equalf(t, resultType.String(), tt.expectedType, "Result should be a subset of the array")
			assert.Equalf(t, nextResultType.String(), tt.expectedType, "NextResult should be a subset of the array")
		})
	}
}
