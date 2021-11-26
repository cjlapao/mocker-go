package mocker

import (
	"strings"

	"github.com/cjlapao/mocker-go/help"
	models "github.com/cjlapao/mocker-go/models/lorem"
)

const NEW_LINE = "\n"

var _loremGenerator *LoremGenerator
var lorem models.Lorem

type LoremGenerator struct {
	Mocker *Mocker
}

func NewLoremGenerator(mocker *Mocker) *LoremGenerator {
	if _loremGenerator != nil {
		return _loremGenerator
	}

	_loremGenerator = &LoremGenerator{}
	_loremGenerator.Mocker = mocker

	return _loremGenerator
}

func (l *LoremGenerator) Word() string {
	return l.Mocker.Random().RandomStrElement(lorem.Words())
}

func (l *LoremGenerator) Words(number int) string {
	words := ""
	for i := 0; i < number; i++ {
		word := l.Mocker.Random().RandomStrElement(lorem.Words())
		if len(words) == 0 {
			words = word
		} else {
			words = words + " " + word
		}
	}

	return words
}

func (l *LoremGenerator) Sentence(wordCount int) string {
	sentence := ""
	for i := 0; i < wordCount; i++ {
		word := l.Mocker.Random().RandomStrElement(lorem.Words())
		if len(sentence) == 0 {
			sentence = strings.ToUpper(string(word[0])) + word[1:]
		} else {
			sentence = sentence + " " + word
		}
	}

	return sentence + "."
}

func (l *LoremGenerator) Sentences(sentenceCount int, separator string) string {
	sentences := ""
	if separator == "" {
		separator = " "
	}

	for i := 0; i < sentenceCount; i++ {
		wordCount := l.Mocker.Random().IntBetween(3, 10)
		if i == 0 {
			sentences = l.Sentence(wordCount)
		} else {
			sentences = sentences + separator + l.Sentence(wordCount)
		}
	}

	return sentences
}

func (l *LoremGenerator) Slug(wordCount int) string {
	var words = l.Words(wordCount)
	return help.Slugify(words)
}

func (l *LoremGenerator) Paragraph(sentenceCount int) string {
	paragraph := ""

	for i := 0; i < sentenceCount; i++ {
		wordCount := l.Mocker.Random().IntBetween(3, 10)
		if i == 0 {
			paragraph = l.Sentence(wordCount)
		} else {
			paragraph = paragraph + " " + l.Sentence(wordCount)
		}
	}

	return paragraph
}

func (l *LoremGenerator) Paragraphs(paragraphCount int, separator string) string {
	paragraphs := ""
	if separator == "" {
		separator = NEW_LINE
	}

	for i := 0; i < paragraphCount; i++ {
		sentenceCount := l.Mocker.Random().IntBetween(3, 10)
		if i == 0 {
			paragraphs = l.Paragraph(sentenceCount)
		} else {
			paragraphs = paragraphs + separator + l.Paragraph(sentenceCount)
		}
	}

	return paragraphs
}

func (l *LoremGenerator) Lines(lineCount int) string {
	return l.Sentences(lineCount, "\n")
}

func (l *LoremGenerator) Text(times int) string {
	methods := [...]string{
		"sentence",
		"sentences",
		"paragraph",
		"paragraphs",
	}

	result := ""
	for i := 0; i < times; i++ {
		randomMethod := l.Mocker.Random().RandomStrElement(methods[:])
		randomCount := l.Mocker.Random().IntBetween(3, 10)
		switch randomMethod {
		case "sentence":
			result = result + l.Sentence(randomCount)
		case "sentences":
			result = result + l.Sentences(randomCount, " ")
		case "paragraph":
			result = result + l.Paragraph(randomCount)
		case "paragraphs":
			result = result + l.Paragraphs(randomCount, NEW_LINE)
		}
	}

	return result
}
