package models

type EnglishAddress struct{}

func (l EnglishAddress) PostCode() []string {
	return []string{
		`[0-9]{5}`,
		`[0-9]{5}\-[0-9]{3}`,
	}
}
