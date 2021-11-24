package models

var english English

type EnglishUnitedKingdom struct {
}

func (l EnglishUnitedKingdom) Prefixes() []string {
	return english.Prefixes()
}

func (l EnglishUnitedKingdom) Suffixes() []string {
	return english.Suffixes()
}

func (l EnglishUnitedKingdom) Names() []string {
	return english.Names()
}

func (l EnglishUnitedKingdom) Countries() []string {
	return english.Countries()
}
