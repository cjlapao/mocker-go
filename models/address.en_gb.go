package models

var english EnglishAddress

type EnglishUnitedKingdom struct {
}

func (l EnglishUnitedKingdom) CityPrefixes() []string {
	return english.CityPrefixes()
}

func (l EnglishUnitedKingdom) CitySuffixes() []string {
	return english.CitySuffixes()
}

func (l EnglishUnitedKingdom) CityNames() []string {
	return english.CityNames()
}

func (l EnglishUnitedKingdom) Countries() []string {
	return english.Countries()
}

func (l EnglishUnitedKingdom) PostCode() []string {
	return []string{
		`[A-Z]{1}[0-9]{1} [0-9]{1}[A-Z]{2}`,
		`[A-Z]{2}[0-9]{2} [0-9]{1}[A-Z]{2}`,
		`[A-Z]{2}[0-9]{1}[A-Z]{1} [0-9]{1}[A-Z]{2}`,
	}
}
