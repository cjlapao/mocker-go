package mocker

import "github.com/cjlapao/mocker-go/models"

type AddressGenerator struct {
	Mocker *Mocker
}

func NewAddress() *AddressGenerator {
	generator := AddressGenerator{}
	generator.Mocker = New()

	return &generator
}

func (g *AddressGenerator) City() string {
	hasPrefix := g.Mocker.Boolean().ChanceOfBool(10)
	hasSuffix := g.Mocker.Boolean().ChanceOfBool(25)
	result := ""
	if hasPrefix {
		result = g.Mocker.Random().RandomElement(models.CityPrefix[:])
	}
	result = result + " " + g.Mocker.Random().RandomElement(models.UsCityNames[:])
	if hasSuffix {
		result = result + g.Mocker.Random().RandomElement(models.CitySuffix[:])
	}

	return result
}

func (g *AddressGenerator) Country() string {
	return g.Mocker.Random().RandomElement(models.Country[:])
}
