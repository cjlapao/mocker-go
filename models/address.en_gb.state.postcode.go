package models

func (l EnglishUnitedKingdom) StatePostcodes() map[string]StatePostcode {
	stateMap := make(map[string]StatePostcode)
	stateMap["BEDFORDSHIRE"] = StatePostcode{
		Patterns: []string{
			`LU[1-7]{1} [0-9]{1}[A-Z]{2}`,
			`MK17 [0-9]{1}[A-Z]{2}`,
			`MK4[0-5]{1} [0-9]{1}[A-Z]{2}`,
			`NN10 [0-9]{1}[A-Z]{2}`,
			`NN29 [0-9]{1}[A-Z]{2}`,
			`PE18 [0-9]{1}[A-Z]{2}`,
			`PE19 [0-9]{1}[A-Z]{2}`,
			`SG5 [0-9]{1}[A-Z]{2}`,
			`SG1[5-9]{1} [0-9]{1}[A-Z]{2}`,
		},
	}

	return stateMap
}
