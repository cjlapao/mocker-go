package models

import "github.com/cjlapao/common-go/language"

type NameLocalData interface {
	FirstNames() []string
	FemaleFirstNames() []string
	MaleFirstNames() []string
	LastNames() []string
}

type Name struct{}

func (a Name) Get(locale language.Locale) NameLocalData {
	switch locale {
	case language.English:
		return EnglishName{}
	case language.EnglishUnitedKingdom:
		return EnglishName{}
	case language.EnglishUnitedStates:
		return EnglishName{}
	case language.PortuguesePortugal:
		return PortuguesePortugalName{}
	default:
		return EnglishName{}
	}
}
