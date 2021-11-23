package templates

import (
	"github.com/cjlapao/mocker-go/mocker"
)

type Lorem struct {
}

func (d Lorem) Word() string {
	return mocker.New().Lorem().Word()
}

func (d Lorem) Words(number int) string {
	return mocker.New().Lorem().Words(number)
}

func (d Lorem) Sentence(wordCount int) string {
	return mocker.New().Lorem().Sentence(wordCount)
}

func (d Lorem) Sentences(sentenceCount int, separator string) string {
	return mocker.New().Lorem().Sentences(sentenceCount, separator)
}

func (d Lorem) Slug(wordCount int) string {
	return mocker.New().Lorem().Slug(wordCount)
}

func (d Lorem) Paragraph(sentenceCount int) string {
	return mocker.New().Lorem().Paragraph(sentenceCount)
}

func (d Lorem) Paragraphs(paragraphCount int, separator string) string {
	return mocker.New().Lorem().Paragraphs(paragraphCount, separator)
}

func (d Lorem) Lines(lineCount int) string {
	return mocker.New().Lorem().Lines(lineCount)
}

func (d Lorem) Text(times int) string {
	return mocker.New().Lorem().Text(times)
}
