package controller

import (
	"encoding/json"
	"net/http"

	"github.com/cjlapao/mocker-go/entities"
	"github.com/cjlapao/mocker-go/mocker"
)

func (c *Controller) GetAddressCity(w http.ResponseWriter, r *http.Request) {
	var response entities.MockerApiResponse

	response.Query = "Fake City"
	faker := mocker.New().Address()
	faker.Locale = getRequestLang(r)
	result := faker.City()

	response.Value = map[string]interface{}{"city": result}
	json.NewEncoder(w).Encode(response)
}

func (c *Controller) GetAddressCountry(w http.ResponseWriter, r *http.Request) {
	var response entities.MockerApiResponse
	response.Query = "Fake Country"

	faker := mocker.New().Address()
	faker.Locale = getRequestLang(r)
	result := faker.Country()

	response.Value = map[string]interface{}{"country": result}
	json.NewEncoder(w).Encode(response)
}

func registerAddressRoutes() {
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/address/city", globalController.GetAddressCity).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/address/country", globalController.GetAddressCountry).Methods("GET")
}
