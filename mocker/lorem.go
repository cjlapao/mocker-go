package mocker

import (
	"strings"

	"github.com/cjlapao/mocker-go/models"
)

var _loremGenerator *LoremGenerator

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
	return l.Mocker.RandomElement(models.LoremWords[:])
}

func (l *LoremGenerator) Words(number int) string {
	words := ""
	for i := 0; i < number; i++ {
		word := l.Mocker.RandomElement(models.LoremWords[:])
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
		word := l.Mocker.RandomElement(models.LoremWords[:])
		if len(sentence) == 0 {
			sentence = strings.ToUpper(string(word[0])) + word[1:]
		} else {
			sentence = sentence + " " + word
		}
	}

	return sentence
}
