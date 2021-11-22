package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

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
	case "date.between":
		dateInbetween := mocker.Date().Between(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2021, 10, 31, 23, 59, 59, 999, time.Local))
		message = fmt.Sprintf("[Fake date.between], %v", dateInbetween)
	case "date.recent":
		dateRecent := mocker.Date().Recent(90)
		message = fmt.Sprintf("[Fake date.recent], %v", dateRecent)
	case "date.soon":
		dateSoon := mocker.Date().Recent(30)
		message = fmt.Sprintf("[Fake date.soon], %v", dateSoon)
	case "date.month":
		dateMonth := mocker.Date().Month()
		message = fmt.Sprintf("[Fake date.month], %v", dateMonth)
	case "date.weekday":
		result := mocker.Date().Weekday()
		message = fmt.Sprintf("[Fake date.weekday], %v", result)
	case "date.past":
		result := mocker.Date().Past(5)
		message = fmt.Sprintf("[Fake date.past], %v", result)
	case "date.future":
		result := mocker.Date().Future(3)
		message = fmt.Sprintf("[Fake date.future], %v", result)
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

	if strings.Contains(queryType, "lorem.") {
		message = loremGenerator(mocker, queryType)
	}

	response := entities.MockerApiResponse{
		Message: message,
	}

	json.NewEncoder(w).Encode(response)
}

func loremGenerator(mocker *mocker.Mocker, query string) string {
	switch strings.ToLower(query) {
	case "lorem.word":
		result := mocker.Lorem().Word()
		return result
	case "lorem.words":
		result := mocker.Lorem().Words(10)
		return result
	case "lorem.sentence":
		result := mocker.Lorem().Sentence(10)
		return result
	default:
		return "Unknown query"
	}
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
