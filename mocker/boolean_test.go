package mocker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	// Arrange
	boolean := New().Boolean()

	// Act
	result := boolean.Bool()

	// Assert
	assert.Truef(t, result || !result, "The result should be true or false")
}

func TestBoolAsInt(t *testing.T) {
	// Arrange
	boolean := New().Boolean()

	// Act
	result := boolean.BoolAsInt()

	// Assert
	assert.Truef(t, result == 1 || result == 0, "The result should be true or false in binary mode")
}

func TestBoolAsString(t *testing.T) {
	// Arrange
	boolean := New().Boolean()

	// Act
	result := boolean.BoolAsString("true", "false")

	// Assert
	assert.Truef(t, result == "true" || result == "false", "The result should be true or false in binary mode")
}

func TestChangesOf(t *testing.T) {
	// Arrange
	boolean := New().Boolean()
	results := make([]bool, 0)

	// Act
	for i := 0; i < 10; i++ {
		results = append(results, boolean.ChanceOfBool(85))
	}

	// Assert
	amountOfTrue := 0
	amountOfFalse := 0
	for i := 0; i < 10; i++ {
		if results[i] {
			amountOfTrue += 1
		} else {
			amountOfFalse += 1
		}
	}
	assert.Truef(t, amountOfTrue > amountOfFalse, "The iteration should have more true than false")
}
