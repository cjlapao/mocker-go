package models

import "github.com/cjlapao/common-go/language"

type AddressLocalData interface {
	Prefixes() []string
	Suffixes() []string
	Names() []string
	Countries() []string
}

type AddressLocalCities interface {
}

type Address struct{}

func (a Address) Get(locale language.Locale) AddressLocalData {
	switch locale {
	case language.English:
		return English{}
	case language.EnglishUnitedKingdom:
		return English{}
	case language.EnglishUnitedStates:
		return English{}
	case language.PortuguesePortugal:
		return PortuguesePortugal{}
	default:
		return English{}
	}
}
