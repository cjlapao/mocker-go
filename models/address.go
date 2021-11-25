package models

import "github.com/cjlapao/common-go/language"

type AddressLocalData interface {
	CityPrefixes() []string
	CitySuffixes() []string
	CityNames() []string
	Countries() []string
	PostCode() []string
	State() []string
	StatePostcodes() map[string]StatePostcode
}

type AddressLocalCities interface {
}

type StatePostcode struct {
	Min      int
	Max      int
	Patterns []string
}

type Address struct{}

func (a Address) Get(locale language.Locale) AddressLocalData {
	switch locale {
	case language.English:
		return EnglishAddress{}
	case language.EnglishUnitedKingdom:
		return EnglishUnitedKingdom{}
	case language.EnglishUnitedStates:
		return EnglishAddress{}
	case language.PortuguesePortugal:
		return PortuguesePortugalAddress{}
	default:
		return EnglishAddress{}
	}
}
