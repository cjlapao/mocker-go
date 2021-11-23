package mocker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomInt(t *testing.T) {
	// Arrange
	random := New().Random()

	// Act
	result := random.RandomInt(10)
	nextResult := random.RandomInt(10)

	assert.Truef(t, result >= 0 && result <= 10, "The number should be in the range of 0 and 10")
	assert.Truef(t, nextResult >= 0 && nextResult <= 10, "The number should be in the range of 0 and 10")
	assert.NotEqualf(t, nextResult, result, "The numbers should not be equal in the next run")
}
