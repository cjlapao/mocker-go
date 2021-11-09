package markov

type MarkovModel int64

const (
	SurnamesModel MarkovModel = iota
	UKSurnamesModel
	UkNamesModel
	CompanyNamesModel
)

func (s MarkovModel) String() string {
	switch s {
	case SurnamesModel:
		return "all_surnames.json"
	case UKSurnamesModel:
		return "uk_surnames.json"
	case UkNamesModel:
		return "uk_names.json"
	case CompanyNamesModel:
		return "all_company_names.json"
	}

	return "unknown"
}
