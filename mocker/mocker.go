package mocker

import (
	"math/rand"

	"github.com/cjlapao/mocker-go/help"
	"github.com/cjlapao/mocker-go/serviceprovider"
)

var sp = serviceprovider.New()

type Mocker struct {
	Generator *rand.Rand
}

func (m Mocker) Company() *CompanyGenerator {
	return NewCompanyGenerator(New())
}

func (m Mocker) Names() *NameGenerator {
	return NewNameGenerator(New())
}

func (m Mocker) Boolean() Boolean {
	return Boolean{New()}
}

func (m Mocker) Address() *AddressGenerator {
	return NewAddress()
}

func (m Mocker) Date() *DateGenerator {
	return NewDateGenerator(New())
}

func (m Mocker) Lorem() *LoremGenerator {
	return NewLoremGenerator(New())
}

func (m Mocker) Random() *RandomGenerator {
	return NewRandomGenerator(New())
}

func New() *Mocker {
	var seed help.CryptoSource
	return NewWithSeed(seed)
}

func NewWithSeed(src rand.Source) *Mocker {
	generator := rand.New(src)
	m := Mocker{Generator: generator}
	return &m
}

func Rand() *rand.Rand {
	var seed help.CryptoSource
	return rand.New(seed)
}
