package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/cjlapao/common-go/helper"
	"github.com/cjlapao/common-go/log"
	"github.com/cjlapao/common-go/version"
	"github.com/cjlapao/mocker-go/controller"
	"github.com/cjlapao/mocker-go/serviceprovider"
)

var serviceProvider = serviceprovider.New()
var logger = log.Get()
var versionSvc = version.Get()

func main() {
	versionSvc.Minor = 1
	versionSvc.Name = "Faker API"
	versionSvc.Author = "Carlos Lapao"
	versionSvc.License = "MIT"
	getVersion := helper.GetFlagSwitch("version", false)
	if getVersion {
		format := helper.GetFlagValue("o", "json")
		switch strings.ToLower(format) {
		case "json":
			fmt.Println(versionSvc.PrintVersion(int(version.JSON)))
		case "yaml":
			fmt.Println(versionSvc.PrintVersion(int(version.JSON)))
		default:
			fmt.Println("Please choose a valid format, this can be either json or yaml")
		}
		os.Exit(0)
	}

	versionSvc.PrintAnsiHeader()

	defer func() {
	}()

	controller.StartApi()
}
