package controller

import (
	"encoding/json"
	"net/http"

	"github.com/cjlapao/mocker-go/entities"
	"github.com/cjlapao/mocker-go/mocker"
	"github.com/gorilla/mux"
)

func (c *Controller) GetAddressCity(w http.ResponseWriter, r *http.Request) {
	var response entities.MockerApiResponse

	response.Query = "Fake City"
	faker := mocker.New().Address()
	faker.Locale = getRequestLang(r)
	result := faker.CityName()

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

func (c *Controller) GetAddressPostcode(w http.ResponseWriter, r *http.Request) {
	var response entities.MockerApiResponse
	response.Query = "Fake Postcode"

	faker := mocker.New().Address()
	faker.Locale = getRequestLang(r)
	result := faker.PostCode()

	response.Value = map[string]interface{}{"postcode": result}
	json.NewEncoder(w).Encode(response)
}

func (c *Controller) GetAddressPostcodeByState(w http.ResponseWriter, r *http.Request) {
	var response entities.MockerApiResponse
	response.Query = "Fake Postcode by State"

	stateKey := mux.Vars(r)["state"]

	if stateKey != "" {
		faker := mocker.New().Address()
		faker.Locale = getRequestLang(r)
		result := faker.PostCodeByState(stateKey)
		response.Value = map[string]interface{}{"postcode": result}
	} else {
		response.Value = map[string]interface{}{"postcode": "State not found"}
	}

	json.NewEncoder(w).Encode(response)
}
func registerAddressRoutes() {
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/address/city", globalController.GetAddressCity).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/address/country", globalController.GetAddressCountry).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/address/postcode", globalController.GetAddressPostcode).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/address/postcode/{state:[a-zA-Z]+}", globalController.GetAddressPostcodeByState).Methods("GET")
}
