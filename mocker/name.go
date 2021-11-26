package mocker

import (
	"strings"

	"github.com/cjlapao/common-go/language"
	models "github.com/cjlapao/mocker-go/models/name"
)

type NameGenerator struct {
	Mocker *Mocker
	Locale language.Locale
}

func (n *NameGenerator) Name() string {
	gender := n.Mocker.Boolean().BoolAsInt()
	if gender == 0 {
		return n.NameByGender("male")
	} else {
		return n.NameByGender("female")
	}
}

func (n *NameGenerator) NameByGender(gender string) string {
	var name models.Name
	var firstNames []string
	result := ""

	locale := name.Get(n.Locale)
	hasSuffix := n.Mocker.Boolean().ChanceOfBool(15)
	hasMiddleNames := n.Mocker.Boolean().ChanceOfBool(30)
	amountMiddleNames := n.Mocker.Random().IntBetween(1, 3)
	hasTwoNames := n.Mocker.Boolean().ChanceOfBool(15)

	if len(locale.FemaleFirstNames()) == 0 || len(locale.MaleFirstNames()) == 0 {
		firstNames = locale.FirstNames()
	} else {
		if strings.ToLower(gender) == "female" {
			firstNames = locale.FemaleFirstNames()
		} else {
			firstNames = locale.MaleFirstNames()
		}
	}

	result = n.Mocker.Random().RandomStrElement(firstNames)

	if hasTwoNames {
		result += " " + n.Mocker.Random().RandomStrElement(firstNames)
	}

	if hasMiddleNames {
		for i := 0; i < amountMiddleNames; i++ {
			result += " " + n.Mocker.Random().RandomStrElement(locale.LastNames())
		}
	}

	result += " " + n.Mocker.Random().RandomStrElement(locale.LastNames())

	if hasSuffix {

	}

	return result
}

func (n *NameGenerator) JobTitle() string {
	var name models.Name
	result := ""

	locale := name.Get(n.Locale)
	if len(locale.JobDescriptor()) == 0 && len(locale.JobLevel()) == 0 && len(locale.JobName()) == 0 {
		return ""
	}

	if len(locale.JobDescriptor()) > 0 {
		result += n.Mocker.Random().RandomStrElement(locale.JobDescriptor())
	}

	if len(locale.JobLevel()) > 0 {
		if len(result) > 0 {
			result += " "
		}
		result += n.Mocker.Random().RandomStrElement(locale.JobLevel())
	}

	if len(locale.JobName()) > 0 {
		if len(result) > 0 {
			result += " "
		}
		result += n.Mocker.Random().RandomStrElement(locale.JobName())
	}

	return result
}
