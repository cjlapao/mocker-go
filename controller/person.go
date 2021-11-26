package controller

import (
	"encoding/json"
	"net/http"

	"github.com/cjlapao/mocker-go/entities"
	"github.com/cjlapao/mocker-go/mocker"
	"github.com/gorilla/mux"
)

func (c *Controller) GetPerson(w http.ResponseWriter, r *http.Request) {
	var response entities.MockerApiResponse

	response.Query = "Fake Person"
	faker := mocker.New().Person()
	faker.Locale = getRequestLang(r)
	result := faker.Person()

	response.Value = map[string]interface{}{"person": result}
	json.NewEncoder(w).Encode(response)
}

func (c *Controller) GetPersonByGender(w http.ResponseWriter, r *http.Request) {
	var response entities.MockerApiResponse
	gender := mux.Vars(r)["gender"]
	response.Query = "Fake Person by Gender"
	faker := mocker.New().Person()
	faker.Locale = getRequestLang(r)
	result := faker.PersonByGender(gender)

	response.Value = map[string]interface{}{"person": result}
	json.NewEncoder(w).Encode(response)
}

func registerPersonRoutes() {
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/person", globalController.GetPerson).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/person/{gender:[a-zA-Z]+}", globalController.GetPersonByGender).Methods("GET")
}
