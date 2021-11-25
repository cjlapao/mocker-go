package models

func (l EnglishAddress) StatePostcodes() map[string]StatePostcode {
	stateMap := make(map[string]StatePostcode)
	stateMap["AK"] = StatePostcode{
		Min: 99501,
		Max: 99950,
	}
	stateMap["AL"] = StatePostcode{
		Min: 35004,
		Max: 36925,
	}
	stateMap["AR"] = StatePostcode{
		Min: 71601,
		Max: 72959,
	}
	stateMap["AZ"] = StatePostcode{
		Min: 85001,
		Max: 86556,
	}
	stateMap["CA"] = StatePostcode{
		Min: 90001,
		Max: 96162,
	}
	stateMap["CO"] = StatePostcode{
		Min: 80001,
		Max: 81658,
	}
	stateMap["CT"] = StatePostcode{
		Min: 6001,
		Max: 6389,
	}
	stateMap["DC"] = StatePostcode{
		Min: 20001,
		Max: 20039,
	}
	stateMap["DE"] = StatePostcode{
		Min: 19701,
		Max: 19980,
	}
	stateMap["FL"] = StatePostcode{
		Min: 32004,
		Max: 34997,
	}
	stateMap["GA"] = StatePostcode{
		Min: 30001,
		Max: 31999,
	}
	stateMap["HI"] = StatePostcode{
		Min: 96701,
		Max: 96898,
	}
	stateMap["IA"] = StatePostcode{
		Min: 50001,
		Max: 52809,
	}
	stateMap["ID"] = StatePostcode{
		Min: 83201,
		Max: 83876,
	}
	stateMap["IL"] = StatePostcode{
		Min: 60001,
		Max: 62999,
	}
	stateMap["IN"] = StatePostcode{
		Min: 46001,
		Max: 47997,
	}
	stateMap["KS"] = StatePostcode{
		Min: 66002,
		Max: 67954,
	}
	stateMap["KY"] = StatePostcode{
		Min: 40003,
		Max: 42788,
	}
	stateMap["LA"] = StatePostcode{
		Min: 70001,
		Max: 71232,
	}
	stateMap["MA"] = StatePostcode{
		Min: 1001,
		Max: 2791,
	}
	stateMap["MD"] = StatePostcode{
		Min: 20331,
		Max: 20331,
	}
	stateMap["ME"] = StatePostcode{
		Min: 3901,
		Max: 4992,
	}
	stateMap["MI"] = StatePostcode{
		Min: 48001,
		Max: 49971,
	}
	stateMap["MN"] = StatePostcode{
		Min: 55001,
		Max: 56763,
	}
	stateMap["MO"] = StatePostcode{
		Min: 63001,
		Max: 65899,
	}
	stateMap["MS"] = StatePostcode{
		Min: 38601,
		Max: 39776,
	}
	stateMap["MT"] = StatePostcode{
		Min: 59001,
		Max: 59937,
	}
	stateMap["NC"] = StatePostcode{
		Min: 27006,
		Max: 28909,
	}
	stateMap["ND"] = StatePostcode{
		Min: 58001,
		Max: 58856,
	}
	stateMap["NE"] = StatePostcode{
		Min: 68001,
		Max: 68118,
	}
	stateMap["NH"] = StatePostcode{
		Min: 3031,
		Max: 3897,
	}
	stateMap["NJ"] = StatePostcode{
		Min: 7001,
		Max: 8989,
	}
	stateMap["NM"] = StatePostcode{
		Min: 87001,
		Max: 88441,
	}
	stateMap["NV"] = StatePostcode{
		Min: 88901,
		Max: 89883,
	}
	stateMap["NY"] = StatePostcode{
		Min: 6390,
		Max: 6390,
	}
	stateMap["OH"] = StatePostcode{
		Min: 43001,
		Max: 45999,
	}
	stateMap["OK"] = StatePostcode{
		Min: 73001,
		Max: 73199,
	}
	stateMap["OR"] = StatePostcode{
		Min: 97001,
		Max: 97920,
	}
	stateMap["PA"] = StatePostcode{
		Min: 15001,
		Max: 19640,
	}
	stateMap["PR"] = StatePostcode{
		Min: 0,
		Max: 0,
	}
	stateMap["RI"] = StatePostcode{
		Min: 2801,
		Max: 2940,
	}
	stateMap["SC"] = StatePostcode{
		Min: 29001,
		Max: 29948,
	}
	stateMap["SD"] = StatePostcode{
		Min: 57001,
		Max: 57799,
	}
	stateMap["TN"] = StatePostcode{
		Min: 37010,
		Max: 38589,
	}
	stateMap["TX"] = StatePostcode{
		Min: 75503,
		Max: 79999,
	}
	stateMap["UT"] = StatePostcode{
		Min: 84001,
		Max: 84784,
	}
	stateMap["VA"] = StatePostcode{
		Min: 20040,
		Max: 20041,
	}
	stateMap["VT"] = StatePostcode{
		Min: 5001,
		Max: 5495,
	}
	stateMap["WA"] = StatePostcode{
		Min: 98001,
		Max: 99403,
	}
	stateMap["WI"] = StatePostcode{
		Min: 53001,
		Max: 54990,
	}
	stateMap["WV"] = StatePostcode{
		Min: 24701,
		Max: 26886,
	}
	stateMap["WY"] = StatePostcode{
		Min: 82001,
		Max: 83128,
	}

	return stateMap
}
