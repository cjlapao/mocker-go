package controller

import (
	"encoding/json"
	"net/http"

	"github.com/cjlapao/mocker-go/entities"
	"github.com/cjlapao/mocker-go/mocker"
)

func (c *Controller) GetBoolean(w http.ResponseWriter, r *http.Request) {
	var response entities.MockerApiResponse

	response.Query = "Fake Boolean"
	faker := mocker.New().Boolean()
	result := faker.Bool()

	response.Value = map[string]interface{}{"bool": result}
	json.NewEncoder(w).Encode(response)
}

func registerBooleanRoutes() {
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/boolean", globalController.GetBoolean).Methods("GET")
}
