package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/cjlapao/mocker-go/entities"
	"github.com/cjlapao/mocker-go/markov"
	"github.com/cjlapao/mocker-go/mocker"
	"github.com/cjlapao/mocker-go/templates"
)

// Login Generate a token for a valid user
func (c *Controller) Generator(w http.ResponseWriter, r *http.Request) {
	mocker := mocker.New()
	queryType := r.URL.Query().Get("type")
	var message string
	switch strings.ToLower(queryType) {
	case "company.jobtitle":
		jobTitle := mocker.Company().JobTitle()
		message = fmt.Sprintf("[Fake Company.JobTitle], %v", jobTitle)
	case "company.name":
		companyName := mocker.Company().Name()
		message = fmt.Sprintf("[Fake Company.Name], %v", companyName)
	case "names.surname":
		surname := mocker.Names().Name()
		message = fmt.Sprintf("[Fake Names.Surname], %v", surname)
	case "boolean.chancesof":
		booleanResult := mocker.Boolean().ChanceOfBool(75)
		message = fmt.Sprintf("[Fake boolean.changesOf], %v", booleanResult)
	case "address.city":
		city := mocker.Address().City()
		message = fmt.Sprintf("[Fake address.country], %v", city)
	case "address.country":
		country := mocker.Address().Country()
		message = fmt.Sprintf("[Fake address.country], %v", country)
	case "template":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			json.NewEncoder(w).Encode("error")
			return
		}
		booleanResult := templates.JustDoIt(string(body))
		w.Write([]byte(booleanResult))
		return
	default:
		message = fmt.Sprint("Hello Faker")
	}

	response := entities.MockerApiResponse{
		Message: message,
	}

	json.NewEncoder(w).Encode(response)
}

func (c *Controller) TrainModel(w http.ResponseWriter, r *http.Request) {
	queryType := r.URL.Query().Get("type")
	var message string
	switch strings.ToLower(queryType) {
	case "companies.all":
		generator := markov.NewGenerator()
		generator.Train(markov.CompanyNamesDatabase, 5)
	case "surnames.all":
		generator := markov.NewGenerator()
		generator.Train(markov.SurnamesDatabase, 4)
	case "surnames.uk":
		generator := markov.NewGenerator()
		generator.Train(markov.UKSurnamesDatabase, 4)
	case "names.uk":
		generator := markov.NewGenerator()
		generator.Train(markov.UkNamesDatabase, 4)
	default:
		message = fmt.Sprint("Hello Faker")
	}

	response := entities.MockerApiResponse{
		Message: message,
	}

	json.NewEncoder(w).Encode(response)
}
