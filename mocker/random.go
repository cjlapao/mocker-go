package mocker

import "github.com/cjlapao/common-go/cryptorand"

type RandomGenerator struct {
	Mocker *Mocker
}

func (m *RandomGenerator) IntBetween(min, max int) int {
	diff := max - min

	if diff == 0 {
		return min
	}

	if diff < 0 {
		diff = min - max
		return cryptorand.Rand().Intn(diff+1) + max
	}
	return cryptorand.Rand().Intn(diff+1) + min
}

func (m *RandomGenerator) RandomInt(max int) int {
	return m.IntBetween(0, max)
}

func (m *RandomGenerator) RandomStrElement(a []string) string {
	if len(a) == 0 {
		return ""
	}
	i := m.IntBetween(0, len(a)-1)
	return a[i]
}

func (m *RandomGenerator) RandomIntElement(a []int) int {
	i := m.IntBetween(0, len(a)-1)
	return a[i]
}

func (m *RandomGenerator) RandomElement(a []interface{}) interface{} {
	index := m.IntBetween(0, len(a)-1)

	return a[index]
}
