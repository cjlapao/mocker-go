package controller

import (
	"fmt"
	"log"

	"net/http"

	commonLogger "github.com/cjlapao/common-go/log"
	"github.com/cjlapao/common-go/version"
	"github.com/cjlapao/mocker-go/database"
	"github.com/cjlapao/mocker-go/repositories"
	"github.com/cjlapao/mocker-go/serviceprovider"
	"github.com/gorilla/mux"
)

var logger = commonLogger.Get()
var versionSvc = version.Get()

// var router mux.Router
var databaseContext database.MongoFactory
var repo repositories.Repository
var serviceProvider = serviceprovider.New()

// Controller Controller structure
type Controller struct {
	Router     *mux.Router
	Repository *repositories.Repository
}

var globalController *Controller

// NewAPIController  Creates a new controller
func NewAPIController(router *mux.Router, repo repositories.Repository) *Controller {
	if globalController != nil {
		return globalController
	}

	controller := Controller{
		Router:     router,
		Repository: &repo,
	}

	controller.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/login", controller.Login).Methods("POST")
	controller.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/validate", controller.Validate).Methods("GET")
	controller.Router.HandleFunc(serviceProvider.Context.ApiPrefix+"/generator", controller.Generator).Methods("GET")

	globalController = &controller
	return globalController
}

func StartApi() {
	logger.Notice("Starting Mocker Rest API v%v", versionSvc.String())
	if serviceProvider.Context.BackendEnabled {
		logger.Info("Found MongoDB connection, enabling MongoDB backend...")
		databaseContext = database.NewFactory()
		repo = repositories.NewRepo(&databaseContext)
		pushTestData()
	}

	// Creating the default route maps
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Homepage!")
	fmt.Println("endpoint Hit: homepage")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(commonMiddleware)
	router.HandleFunc(serviceProvider.Context.ApiPrefix+"/", homePage)
	_ = NewAPIController(router, repo)
	logger.Info("Api listening on http://localhost:" + serviceProvider.Context.ApiPort + serviceProvider.Context.ApiPrefix)
	logger.Success("Finished Init")
	log.Fatal(http.ListenAndServe(":"+serviceProvider.Context.ApiPort, router))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerInfo := fmt.Sprintf("Handling [%v] %v", r.Method, r.RequestURI)
		logger.Info(handlerInfo)
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// Disabled for now
func pushTestData() {
	// importData := entities.ImportData{}
	// if !helper.FileExists("demo.json") {
	// 	logger.Error("File demo.json not found, not importing data")
	// }

	// data, err := ioutil.ReadFile("demo.json")

	// if err != nil {
	// 	logger.LogError(err)
	// 	return
	// }

	// json.Unmarshal(data, &importData)
	// if len(importData.Articles) > 0 {
	// 	logger.Info("Importing Demo articles")
	// 	repo.UpsertManyArticles(importData.Articles)
	// }
	// if len(importData.People) > 0 {
	// 	logger.Info("Importing Demo people")
	// 	repo.UpsertManyPersons(importData.People)
	// }
}
