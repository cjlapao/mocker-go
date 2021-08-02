package executioncontext

import (
	"os"
	"strings"

	"github.com/cjlapao/common-go/helper"
	"github.com/cjlapao/common-go/log"
)

var globalContext *Context
var logger = log.Get()

// Context entity
type Context struct {
	MongoConnectionString string `json:"mongoConnectionString"`
	Database              string `json:"database"`
	ShowHelp              bool   `json:"help"`
	BackendEnabled        bool   `json:"backendEnabled"`
	ApiPrefix             string `json:"apiPrefix"`
	ApiPort               string `json:"apiPort"`
	CompanyNameSourceUrl  string `json:"companyNameSourceUrl"`
}

func Get() *Context {
	if globalContext != nil {
		return globalContext
	}

	logger.Debug("Creating Execution Context")
	globalContext = &Context{
		ShowHelp: helper.GetFlagSwitch("help", false),
	}

	globalContext.Getenv()

	return globalContext
}

// Getenv gets the environment variables for the entities
func (e *Context) Getenv() {

	e.Database = os.Getenv("MOCKER_DATABASENAME")
	e.MongoConnectionString = os.Getenv("MOCKER_MONGO_CONNECTION_STRING")
	e.ApiPrefix = os.Getenv("MOCKER_API_PREFIX")
	e.ApiPort = os.Getenv("MOCKER_API_PORT")
	e.CompanyNameSourceUrl = os.Getenv("MOCKER_COMPANY_NAME_SOURCE_URL")

	if e.MongoConnectionString == "" {
		e.BackendEnabled = false
	} else {
		e.BackendEnabled = true
	}

	if e.Database == "" {
		e.Database = "MOCKER_demo_db"
	}

	if e.ApiPort == "" {
		e.ApiPort = "80"
	}

	if e.CompanyNameSourceUrl == "" {
		e.CompanyNameSourceUrl = "https://raw.githubusercontent.com/notpeter/crunchbase-data/master/companies.csv"
	}

	if strings.HasSuffix(e.ApiPrefix, "/") {
		e.ApiPrefix = strings.TrimSuffix(e.ApiPrefix, "/")
	}
}
