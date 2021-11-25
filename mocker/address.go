package mocker

import (
	"fmt"
	"strings"

	"github.com/cjlapao/common-go/cryptorand"
	"github.com/cjlapao/common-go/language"
	"github.com/cjlapao/mocker-go/models"
	regen "github.com/zach-klippenstein/goregen"
)

type AddressGenerator struct {
	Mocker *Mocker
	Locale language.Locale
}

func (g *AddressGenerator) CityName() string {
	hasPrefix := g.Mocker.Boolean().ChanceOfBool(10)
	hasSuffix := g.Mocker.Boolean().ChanceOfBool(25)
	result := ""

	var address models.Address

	locale := address.Get(g.Locale)

	if hasPrefix && len(locale.CityPrefixes()) > 0 {
		result = g.Mocker.Random().RandomStrElement(locale.CityPrefixes())
	}

	if len(result) > 0 {
		result = result + " "
	}
	result = result + g.Mocker.Random().RandomStrElement(locale.CityNames())

	if hasSuffix && len(locale.CitySuffixes()) > 0 {
		result = result + g.Mocker.Random().RandomStrElement(locale.CitySuffixes())
	}

	return result
}

func (g *AddressGenerator) CityPrefix() string {
	var address models.Address

	locale := address.Get(g.Locale)

	prefixes := locale.CityPrefixes()
	if len(prefixes) == 0 {
		return ""
	}

	return g.Mocker.Random().RandomStrElement(prefixes)
}

func (g *AddressGenerator) CitySuffix() string {
	var address models.Address

	locale := address.Get(g.Locale)

	prefixes := locale.CitySuffixes()
	if len(prefixes) == 0 {
		return ""
	}

	return g.Mocker.Random().RandomStrElement(prefixes)
}

func (g *AddressGenerator) Country() string {
	var address models.Address

	locale := address.Get(g.Locale)

	return g.Mocker.Random().RandomStrElement(locale.Countries())
}

func (g *AddressGenerator) PostCode() string {
	var address models.Address

	locale := address.Get(g.Locale)
	var src cryptorand.CryptoSource
	generator, err := regen.NewGenerator(g.Mocker.Random().RandomStrElement(locale.PostCode()), &regen.GeneratorArgs{RngSource: src})
	if err != nil {
		return ""
	}

	return generator.Generate()
}

func (g *AddressGenerator) PostCodeByState(state string) string {
	var address models.Address
	state = strings.ToUpper(state)
	locale := address.Get(g.Locale)
	if len(locale.StatePostcodes()) > 0 {
		statePostCodes := locale.StatePostcodes()[state]

		if len(statePostCodes.Patterns) > 0 {
			var src cryptorand.CryptoSource
			generator, err := regen.NewGenerator(g.Mocker.Random().RandomStrElement(statePostCodes.Patterns), &regen.GeneratorArgs{RngSource: src})
			if err != nil {
				return ""
			}
			return generator.Generate()
		} else if statePostCodes.Min > 0 {
			return fmt.Sprint(g.Mocker.Random().IntBetween(statePostCodes.Min, statePostCodes.Max))
		}
	} else {
		return g.PostCode()
	}

	return ""
}
