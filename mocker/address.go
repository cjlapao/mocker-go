package mocker

import (
	"github.com/cjlapao/common-go/language"
	"github.com/cjlapao/mocker-go/models"
)

type AddressGenerator struct {
	Mocker *Mocker
	Locale language.Locale
}

func (g *AddressGenerator) City() string {
	hasPrefix := g.Mocker.Boolean().ChanceOfBool(10)
	hasSuffix := g.Mocker.Boolean().ChanceOfBool(25)
	result := ""

	var address models.Address

	locale := address.Get(g.Locale)

	if hasPrefix && len(locale.Prefixes()) > 0 {
		result = g.Mocker.Random().RandomStrElement(locale.Prefixes())
	}

	if len(result) > 0 {
		result = result + " "
	}
	result = result + g.Mocker.Random().RandomStrElement(locale.Names())

	if hasSuffix && len(locale.Suffixes()) > 0 {
		result = result + g.Mocker.Random().RandomStrElement(locale.Suffixes())
	}

	return result
}

func (g *AddressGenerator) Country() string {
	var address models.Address

	locale := address.Get(g.Locale)

	return g.Mocker.Random().RandomStrElement(locale.Countries())
}
