package markov

type TrainingDatabase int64

const (
	SurnamesDatabase TrainingDatabase = iota
	UKSurnamesDatabase
	UkNamesDatabase
	CompanyNamesDatabase
)

func (s TrainingDatabase) String() string {
	switch s {
	case SurnamesDatabase:
		return "all_surnames.txt"
	case UKSurnamesDatabase:
		return "uk_surnames.txt"
	case UkNamesDatabase:
		return "uk_names.txt"
	case CompanyNamesDatabase:
		return "all_company_names.txt"
	}

	return "unknown"
}

func (s TrainingDatabase) ToMarkovModel() MarkovModel {
	switch s {
	case SurnamesDatabase:
		return SurnamesModel
	case UKSurnamesDatabase:
		return UKSurnamesModel
	case UkNamesDatabase:
		return UkNamesModel
	case CompanyNamesDatabase:
		return CompanyNamesModel
	}

	return SurnamesModel
}
