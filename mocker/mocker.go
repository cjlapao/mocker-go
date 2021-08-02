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

func (m Mocker) Company() Company {
	return Company{&m}
}

func (m Mocker) Boolean() Boolean {
	return Boolean{&m}
}

func New() (m Mocker) {
	seed := rand.NewSource(time.Now().Unix())
	m = NewWithSeed(seed)
	return
}

func NewWithSeed(src rand.Source) (f Mocker) {
	generator := rand.New(src)
	f = Mocker{Generator: generator}
	return
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

	return m.Generator.Intn(diff+1) + min
}
