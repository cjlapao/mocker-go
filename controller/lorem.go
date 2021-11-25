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
	var random = mocker.New().Random()
	qtyKey := mux.Vars(r)["qty"]
	var qty int
	var err error

	if qtyKey != "" {
		qty, err = strconv.Atoi(qtyKey)
		if err != nil {
			qty = random.IntBetween(2, 10)
		}
	} else {
		queryQty := int(getRequestIntValue(r, "qty"))
		if queryQty <= 0 {
			queryQty = random.IntBetween(2, 10)
		}
		qty = queryQty
	}

	response.Query = "Fake Lorem Words"
	faker := mocker.New().Lorem()
	result := faker.Words(int(qty))

	response.Value = map[string]interface{}{"words": result}
	json.NewEncoder(w).Encode(response)
}

func (c *Controller) GetLoremSentence(w http.ResponseWriter, r *http.Request) {
	var response entities.MockerApiResponse
	var random = mocker.New().Random()
	qtyKey := mux.Vars(r)["qty"]
	var qty int
	var err error

	if qtyKey != "" {
		qty, err = strconv.Atoi(qtyKey)
		if err != nil {
			qty = random.IntBetween(2, 10)
		}
	} else {
		queryQty := int(getRequestIntValue(r, "word-qty"))
		if queryQty <= 0 {
			queryQty = random.IntBetween(2, 10)
		}
		qty = queryQty
	}

	response.Query = "Fake Lorem Sentence"
	faker := mocker.New().Lorem()
	result := faker.Sentence(qty)

	response.Value = map[string]interface{}{"sentence": result}
	json.NewEncoder(w).Encode(response)
}

func (c *Controller) GetLoremSentences(w http.ResponseWriter, r *http.Request) {
	var response entities.MockerApiResponse
	separator := getRequestStrValue(r, "separator")
	if separator == "" {
		separator = " "
	}

	var random = mocker.New().Random()
	qtyKey := mux.Vars(r)["qty"]
	var qty int
	var err error

	if qtyKey != "" {
		qty, err = strconv.Atoi(qtyKey)
		if err != nil {
			qty = random.IntBetween(2, 10)
		}
	} else {
		queryQty := int(getRequestIntValue(r, "qty"))
		if queryQty <= 0 {
			queryQty = random.IntBetween(2, 10)
		}
		qty = queryQty
	}

	response.Query = "Fake Lorem Sentences"
	faker := mocker.New().Lorem()
	result := faker.Sentences(int(qty), separator)

	response.Value = map[string]interface{}{"sentences": result}
	json.NewEncoder(w).Encode(response)
}

func (c *Controller) GetLoremSlug(w http.ResponseWriter, r *http.Request) {
	var response entities.MockerApiResponse
	var random = mocker.New().Random()
	qtyKey := mux.Vars(r)["qty"]
	var qty int
	var err error

	if qtyKey != "" {
		qty, err = strconv.Atoi(qtyKey)
		if err != nil {
			qty = random.IntBetween(2, 10)
		}
	} else {
		queryQty := int(getRequestIntValue(r, "qty"))
		if queryQty <= 0 {
			queryQty = random.IntBetween(2, 10)
		}
		qty = queryQty
	}

	response.Query = "Fake Lorem Slug"
	faker := mocker.New().Lorem()
	result := faker.Slug(int(qty))

	response.Value = map[string]interface{}{"slug": result}
	json.NewEncoder(w).Encode(response)
}

func (c *Controller) GetLoremParagraph(w http.ResponseWriter, r *http.Request) {
	var response entities.MockerApiResponse
	var random = mocker.New().Random()
	qtyKey := mux.Vars(r)["qty"]
	var qty int
	var err error

	if qtyKey != "" {
		qty, err = strconv.Atoi(qtyKey)
		if err != nil {
			qty = random.IntBetween(2, 10)
		}
	} else {
		queryQty := int(getRequestIntValue(r, "sentence-qty"))
		if queryQty <= 0 {
			queryQty = random.IntBetween(2, 10)
		}
		qty = queryQty
	}

	response.Query = "Fake Lorem Paragraph"
	faker := mocker.New().Lorem()
	result := faker.Paragraph(qty)

	response.Value = map[string]interface{}{"paragraph": result}
	json.NewEncoder(w).Encode(response)
}

func (c *Controller) GetLoremParagraphs(w http.ResponseWriter, r *http.Request) {
	var response entities.MockerApiResponse
	separator := getRequestStrValue(r, "separator")
	if separator == "" {
		separator = "\r\n"
	}

	var random = mocker.New().Random()
	qtyKey := mux.Vars(r)["qty"]
	var qty int
	var err error

	if qtyKey != "" {
		qty, err = strconv.Atoi(qtyKey)
		if err != nil {
			qty = random.IntBetween(2, 10)
		}
	} else {
		queryQty := int(getRequestIntValue(r, "qty"))
		if queryQty <= 0 {
			queryQty = random.IntBetween(2, 10)
		}
		qty = queryQty
	}

	response.Query = "Fake Lorem Paragraphs"
	faker := mocker.New().Lorem()
	result := faker.Paragraphs(int(qty), separator)

	response.Value = map[string]interface{}{"paragraphs": result}
	json.NewEncoder(w).Encode(response)
}

func (c *Controller) GetLoremLines(w http.ResponseWriter, r *http.Request) {
	var response entities.MockerApiResponse

	var random = mocker.New().Random()
	qtyKey := mux.Vars(r)["qty"]
	var qty int
	var err error

	if qtyKey != "" {
		qty, err = strconv.Atoi(qtyKey)
		if err != nil {
			qty = random.IntBetween(2, 10)
		}
	} else {
		queryQty := int(getRequestIntValue(r, "qty"))
		if queryQty <= 0 {
			queryQty = random.IntBetween(2, 10)
		}
		qty = queryQty
	}

	response.Query = "Fake Lorem Lines"
	faker := mocker.New().Lorem()
	result := faker.Lines(int(qty))

	response.Value = map[string]interface{}{"lines": result}
	json.NewEncoder(w).Encode(response)
}

func (c *Controller) GetLoremText(w http.ResponseWriter, r *http.Request) {
	var response entities.MockerApiResponse

	var random = mocker.New().Random()
	qtyKey := mux.Vars(r)["qty"]
	var qty int
	var err error

	if qtyKey != "" {
		qty, err = strconv.Atoi(qtyKey)
		if err != nil {
			qty = random.IntBetween(2, 10)
		}
	} else {
		queryQty := int(getRequestIntValue(r, "qty"))
		if queryQty <= 0 {
			queryQty = random.IntBetween(2, 10)
		}
		qty = queryQty
	}

	response.Query = "Fake Lorem Text"
	faker := mocker.New().Lorem()
	result := faker.Text(int(qty))

	response.Value = map[string]interface{}{"text": result}
	json.NewEncoder(w).Encode(response)
}

func registerLoremRoutes() {
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/lorem/word", globalController.GetLoremWord).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/lorem/words", globalController.GetLoremWords).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/lorem/words/{qty:[0-9]+}", globalController.GetLoremWords).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/lorem/sentence", globalController.GetLoremSentence).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/lorem/sentence/words/{qty:[0-9]+}", globalController.GetLoremSentence).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/lorem/sentences", globalController.GetLoremSentences).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/lorem/sentences/{qty:[0-9]+}", globalController.GetLoremSentences).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/lorem/slug", globalController.GetLoremSlug).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/lorem/slug/{qty:[0-9]+}", globalController.GetLoremSlug).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/lorem/paragraph", globalController.GetLoremParagraph).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/lorem/paragraph/sentence/{qty:[0-9]+}", globalController.GetLoremParagraph).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/lorem/paragraphs", globalController.GetLoremParagraphs).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/lorem/paragraphs/{qty:[0-9]+}", globalController.GetLoremParagraphs).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/lorem/lines", globalController.GetLoremLines).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/lorem/lines/{qty:[0-9]+}", globalController.GetLoremLines).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/lorem/text", globalController.GetLoremText).Methods("GET")
	globalController.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator/lorem/text/{qty:[0-9]+}", globalController.GetLoremText).Methods("GET")
}
