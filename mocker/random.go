package mocker

var _randomGenerator *RandomGenerator

type RandomGenerator struct {
	Mocker *Mocker
}

func NewRandomGenerator(mocker *Mocker) *RandomGenerator {
	if _loremGenerator != nil {
		return _randomGenerator
	}

	_randomGenerator = &RandomGenerator{}
	_randomGenerator.Mocker = mocker

	return _randomGenerator
}

func (m *RandomGenerator) RandomInt(max int) int {
	return m.IntBetween(0, max)
}

func (m *RandomGenerator) RandomStringElement(a []string) string {
	i := m.IntBetween(0, len(a)-1)
	return a[i]
}

func (m *RandomGenerator) RandomIntElement(a []int) int {
	i := m.IntBetween(0, len(a)-1)
	return a[i]
}

func (m *RandomGenerator) IntBetween(min, max int) int {
	diff := max - min

	if diff == 0 {
		return min
	}

	if diff < 0 {
		diff = min - max
		return m.Mocker.Generator.Intn(diff+1) + max
	}
	return m.Mocker.Generator.Intn(diff+1) + min
}

func (m *RandomGenerator) RandomElement(a []string) string {
	index := m.IntBetween(0, len(a)-1)

	return a[index]
}
