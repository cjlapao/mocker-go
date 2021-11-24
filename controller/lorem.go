package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cjlapao/mocker-go/entities"
	"github.com/cjlapao/mocker-go/mocker"
	"github.com/gorilla/mux"
)

func (c *Controller) GetLoremWord(w http.ResponseWriter, r *http.Request) {
	var response entities.MockerApiResponse

	response.Query = "Fake Lorem Word"
	faker := mocker.New().Lorem()
	result := faker.Word()

	response.Value = map[string]interface{}{"word": result}
	json.NewEncoder(w).Encode(response)
}

func (c *Controller) GetLoremWords(w http.ResponseWriter, r *http.Request) {
	var response entities.MockerApiResponse

	qtyKey := mux.Vars(r)["qty"]
	qty := -1
	var err error

	if qtyKey != "" {
		qty, err = strconv.Atoi(qtyKey)
		if err != nil {
			qty = -1
		}
	}

	response.Query = "Fake Lorem Words"
	faker := mocker.New().Lorem()
	result := faker.Words(int(qty))

	response.Value = map[string]interface{}{"words": result}
	json.NewEncoder(w).Encode(response)
}

func (c *Controller) GetLoremSentence(w http.ResponseWriter, r *http.Request) {
	var response entities.MockerApiResponse

	qty := getRequestIntValue(r, "word-qty")

	if qty <= 0 {
		qty = 1
	}

	response.Query = "Fake Lorem Sentence"
	faker := mocker.New().Lorem()
	result := faker.Sentence(int(qty))

	response.Value = map[string]interface{}{"sentence": result}
	json.NewEncoder(w).Encode(response)
}

func (c *Controller) GetLoremSentences(w http.ResponseWriter, r *http.Request) {
	var response entities.MockerApiResponse

	separator := getRequestStrValue(r, "separator")

	qtyKey := mux.Vars(r)["qty"]
	qty := -1
	var err error

	if qtyKey != "" {
		qty, err = strconv.Atoi(qtyKey)
		if err != nil {
			qty = -1
		}
	}

	if separator == "" {
		separator = " "
	}

	response.Query = "Fake Lorem Sentences"
	faker := mocker.New().Lorem()
	result := faker.Sentences(int(qty), separator)

	response.Value = map[string]interface{}{"sentences": result}
	json.NewEncoder(w).Encode(response)
}

func registerLoremRoutes() {
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/lorem/word", globalController.GetLoremWord).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/lorem/words/{qty:[0-9]+}", globalController.GetLoremWords).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/lorem/sentences", globalController.GetLoremSentence).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/lorem/sentences/{qty:[0-9]+}", globalController.GetLoremSentences).Methods("GET")
}
