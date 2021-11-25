package controller

import (
	"encoding/json"
	"net/http"

	"github.com/cjlapao/mocker-go/entities"
	"github.com/cjlapao/mocker-go/mocker"
	"github.com/gorilla/mux"
)

func (c *Controller) GetName(w http.ResponseWriter, r *http.Request) {
	var response entities.MockerApiResponse

	response.Query = "Fake Name"
	faker := mocker.New().Name()
	faker.Locale = getRequestLang(r)
	result := faker.Name()

	response.Value = map[string]interface{}{"name": result}
	json.NewEncoder(w).Encode(response)
}

func (c *Controller) GetNameByGender(w http.ResponseWriter, r *http.Request) {
	var response entities.MockerApiResponse
	gender := mux.Vars(r)["gender"]
	response.Query = "Fake Name by Gender"
	faker := mocker.New().Name()
	faker.Locale = getRequestLang(r)
	result := faker.NameByGender(gender)

	response.Value = map[string]interface{}{"name": result}
	json.NewEncoder(w).Encode(response)
}

func registerNameRoutes() {
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/name", globalController.GetName).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/name/{gender:[a-zA-Z]+}", globalController.GetNameByGender).Methods("GET")
}
