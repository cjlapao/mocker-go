package data

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/cjlapao/common-go/helper"
	"github.com/cjlapao/mocker-go/executioncontext"
)

type CompanyNameDatasource struct {
	IsLoaded bool     `json:"-"`
	Names    []string `json:"names"`
}

var context = executioncontext.Get()
var globalCompanyNameDatasource *CompanyNameDatasource

func NewCompanyNameDatasource() *CompanyNameDatasource {
	if globalCompanyNameDatasource != nil {
		return globalCompanyNameDatasource
	}

	globalCompanyNameDatasource = &CompanyNameDatasource{
		Names: make([]string, 0),
	}

	if !helper.FileExists(".\\company_names.json") {
		_ = globalCompanyNameDatasource.Generate()
	} else {
		globalCompanyNameDatasource.Load()
	}

	return globalCompanyNameDatasource
}

func GetCompanyNameDatasource() *CompanyNameDatasource {
	if globalCompanyNameDatasource == nil {
		return NewCompanyNameDatasource()
	}
	if !globalCompanyNameDatasource.IsLoaded {
		globalCompanyNameDatasource.Load()
	}
	return globalCompanyNameDatasource
}

func (c *CompanyNameDatasource) Generate() error {
	sp.Logger.Info("Generating local company names for mocking")
	resp, err := http.Get(context.CompanyNameSourceUrl)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	csvContentReader := csv.NewReader(resp.Body)
	records, err := csvContentReader.ReadAll()
	if err != nil {
		return err
	}

	for idx, record := range records {
		if idx > 0 {
			c.Names = append(c.Names, record[1])
		}
	}

	jsonContent, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	helper.WriteToFile(string(jsonContent), ".\\company_names.json")
	c.IsLoaded = true
	return nil
}

func (c *CompanyNameDatasource) Load() error {
	sp.Logger.Info("loading local company names for mocking")

	rawContent, err := ioutil.ReadFile(".\\company_names.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawContent, &c)
	if err != nil {
		return err
	}

	c.IsLoaded = true
	return nil
}
