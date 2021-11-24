package mocker

type BooleanGenerator struct {
	Mocker *Mocker
}

func (b BooleanGenerator) Bool() bool {
	return b.Mocker.Random().IntBetween(0, 100) > 50
}

func (b BooleanGenerator) BoolAsInt() int {
	return b.Mocker.Random().RandomIntElement([]int{0, 1})
}

func (b BooleanGenerator) BoolAsString(trueValue string, falseValue string) string {
	return b.Mocker.Random().RandomStrElement([]string{trueValue, falseValue})
}

func (b BooleanGenerator) ChanceOfBool(c int) bool {
	if c <= 0 {
		return false
	} else if c >= 100 {
		return true
	} else {
		return b.Mocker.Random().IntBetween(0, 100) < c
	}
}
