package mocker

import (
	"math/rand"

	"github.com/cjlapao/common-go/cryptorand"
	"github.com/cjlapao/common-go/language"
	"github.com/cjlapao/mocker-go/serviceprovider"
)

var sp = serviceprovider.New()

type Mocker struct {
	Generator *rand.Rand
}

func (m Mocker) Random() *RandomGenerator {
	return &RandomGenerator{New()}
}

func (m Mocker) Boolean() *BooleanGenerator {
	return &BooleanGenerator{New()}
}

func (m Mocker) Address() *AddressGenerator {
	return &AddressGenerator{
		Mocker: New(),
		Locale: language.English,
	}
}

func (m Mocker) Company() *CompanyGenerator {
	return NewCompanyGenerator(New())
}

func (m Mocker) Names() *NameGenerator {
	return NewNameGenerator(New())
}

func (m Mocker) Date() *DateGenerator {
	return NewDateGenerator(New())
}

func (m Mocker) Lorem() *LoremGenerator {
	return NewLoremGenerator(New())
}

func New() *Mocker {
	generator := cryptorand.Rand()
	m := Mocker{Generator: generator}
	return &m
}
