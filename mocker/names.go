package mocker

import (
	"github.com/cjlapao/mocker-go/markov"
)

var _nameGenerator *NameGenerator

type NameGenerator struct {
	Generator *markov.MarkovGenerator
	Mocker    *Mocker
}

func NewNameGenerator(mocker *Mocker) *NameGenerator {
	if _nameGenerator != nil {
		return _nameGenerator
	}

	_nameGenerator = &NameGenerator{}
	_nameGenerator.Mocker = mocker
	_nameGenerator.Generator = markov.NewGenerator()

	return _nameGenerator
}

func (c *NameGenerator) Name() string {
	surname := c.Generator.Generate((markov.UKSurnamesModel))
	name := c.Generator.Generate(markov.UkNamesModel)

	return name + " " + surname
}
