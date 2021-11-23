package mocker

type Boolean struct {
	Mocker *Mocker
}

func (b Boolean) Bool() bool {
	return b.Mocker.Random().IntBetween(0, 100) > 50
}

func (b Boolean) BoolAsInt() int {
	return b.Mocker.Random().RandomIntElement([]int{0, 1})
}

func (b Boolean) BoolAsString(trueValue string, falseValue string) string {
	return b.Mocker.Random().RandomStringElement([]string{trueValue, falseValue})
}

func (b Boolean) ChanceOfBool(c int) bool {
	if c <= 0 {
		return false
	} else if c >= 100 {
		return true
	} else {
		return b.Mocker.Random().IntBetween(0, 100) < c
	}
}
