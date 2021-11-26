package mocker

import (
	"strings"

	"github.com/cjlapao/common-go/language"
)

type Person struct {
	FullName    string `json:"fullName"`
	FirstName   string `json:"firstName"`
	MiddleNames string `json:"middleNames"`
	LastName    string `json:"lastNames"`
	JobTitle    string `json:"jobTitle"`
}

type PersonGenerator struct {
	Mocker *Mocker
	Locale language.Locale
}

func (p *PersonGenerator) Person() Person {
	gender := p.Mocker.Boolean().BoolAsInt()
	if gender == 0 {
		return p.PersonByGender("male")
	} else {
		return p.PersonByGender("female")
	}
}

func (p *PersonGenerator) PersonByGender(gender string) Person {
	person := Person{}
	name := p.Mocker.Name()
	name.Locale = p.Locale

	if gender != "" {
		person.FullName = name.NameByGender(gender)
	} else {
		person.FullName = name.Name()
	}
	nameSegments := strings.Split(person.FullName, " ")
	if len(nameSegments) > 0 {
		switch size := len(nameSegments); {
		case size == 1:
			person.FirstName = nameSegments[0]
		case size == 2:
			person.FirstName = nameSegments[0]
			person.LastName = nameSegments[1]
		case size > 2:
			person.FirstName = nameSegments[0]
			person.LastName = nameSegments[size-1]
			person.MiddleNames = strings.Join(nameSegments[1:size-1], " ")
		}
	}

	person.JobTitle = p.Mocker.Name().JobTitle()
	return person
}
