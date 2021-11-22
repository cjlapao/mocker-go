package mocker

import (
	"math/rand"
	"time"

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

func New() *Mocker {
	seed := rand.NewSource(time.Now().Unix())
	return NewWithSeed(seed)
}

func NewWithSeed(src rand.Source) *Mocker {
	generator := rand.New(src)
	m := Mocker{Generator: generator}
	return &m
}

func (m *Mocker) RandomStringElement(a []string) string {
	i := m.IntBetween(0, len(a)-1)
	return a[i]
}

func (m *Mocker) RandomIntElement(a []int) int {
	i := m.IntBetween(0, len(a)-1)
	return a[i]
}

func (m *Mocker) IntBetween(min, max int) int {
	diff := max - min

	if diff == 0 {
		return min
	}

	if diff < 0 {
		diff = min - max
		return m.Generator.Intn(diff+1) + max
	}
	return m.Generator.Intn(diff+1) + min
}

func (m *Mocker) RandomElement(a []string) string {
	index := m.IntBetween(0, len(a)-1)

	return a[index]
}
