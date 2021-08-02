package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/cjlapao/mocker-go/entities"
	"github.com/cjlapao/mocker-go/mocker"
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
	default:
		message = fmt.Sprint("Hello Faker")
	}

	response := entities.MockerApiResponse{
		Message: message,
	}

	json.NewEncoder(w).Encode(response)
}
