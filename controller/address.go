package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cjlapao/mocker-go/entities"
	"github.com/cjlapao/mocker-go/mocker"
)

func (c *Controller) GetAddressCity(w http.ResponseWriter, r *http.Request) {
	var response entities.MockerApiResponse

	faker := mocker.New().Address()
	faker.Locale = getRequestLang(r)
	prefix := "[Fake City]"
	result := faker.City()
	fakedResult := fmt.Sprintf("%v: %v", prefix, result)

	response.Value = map[string]interface{}{"city": fakedResult}
	json.NewEncoder(w).Encode(response)
}

func (c *Controller) GetAddressCountry(w http.ResponseWriter, r *http.Request) {
	var response entities.MockerApiResponse

	faker := mocker.New().Address()
	faker.Locale = getRequestLang(r)
	prefix := "[Fake Country]"
	result := faker.Country()
	fakedResult := fmt.Sprintf("%v: %v", prefix, result)

	response.Value = map[string]interface{}{"city": fakedResult}
	json.NewEncoder(w).Encode(response)
}
