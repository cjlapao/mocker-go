package mocker

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWord(t *testing.T) {
	// Arrange
	lorem := New().Lorem()

	// Act
	result := lorem.Word()
	nextResult := lorem.Word()

	assert.NotEqualf(t, "", result, "The word should not be empty")
	assert.NotEqualf(t, nextResult, result, "The words should not be equal on second round")
}

func TestWords(t *testing.T) {
	// Arrange
	lorem := New().Lorem()

	// Act
	result := lorem.Words(5)

	assert.Equalf(t, 5, len(strings.Split(result, " ")), "The amount of words should match")
}

func TestSentence(t *testing.T) {
	// Arrange
	lorem := New().Lorem()

	// Act
	result := lorem.Sentence(10)
	resultArray := strings.Split(result, " ")

	assert.Equalf(t, 10, len(resultArray), "The amount of words should match")
	assert.Truef(t, strings.HasSuffix(resultArray[len(resultArray)-1], "."), "The sentence should be finished with a .")
}

func TestSentences(t *testing.T) {
	// Arrange
	lorem := New().Lorem()

	// Act
	result := lorem.Sentences(2, ":")
	resultArray := strings.Split(result, ":")

	assert.Equal(t, 2, len(resultArray))
}

func TestSlug(t *testing.T) {
	// Arrange
	lorem := New().Lorem()

	// Act
	result := lorem.Slug(5)

	assert.Equalf(t, 5, len(strings.Split(result, "-")), "The amount of words should match")
}

func TestParagraph(t *testing.T) {
	// Arrange
	lorem := New().Lorem()

	// Act
	result := lorem.Paragraph(10)
	resultArray := strings.Split(result, ". ")

	assert.Equalf(t, 10, len(resultArray), "The amount of sentences should match")
}

func TestParagraphs(t *testing.T) {
	// Arrange
	lorem := New().Lorem()

	// Act
	result := lorem.Paragraphs(2, ":")
	resultArray := strings.Split(result, ":")

	assert.Equal(t, 2, len(resultArray))
}

func TestLines(t *testing.T) {
	// Arrange
	lorem := New().Lorem()

	// Act
	result := lorem.Lines(2)
	resultArray := strings.Split(result, "\n")

	assert.Equal(t, 2, len(resultArray))
}
