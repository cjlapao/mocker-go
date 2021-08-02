package serviceprovider

import (
	"github.com/cjlapao/common-go/log"
	"github.com/cjlapao/common-go/version"
	"github.com/cjlapao/mocker-go/executioncontext"
)

type ServiceProvider struct {
	Context *executioncontext.Context
	Version *version.Version
	Logger  *log.Logger
}

var globalProvider *ServiceProvider

func New() *ServiceProvider {
	if globalProvider != nil {
		return globalProvider
	}

	globalProvider = &ServiceProvider{}
	globalProvider.Context = executioncontext.Get()
	globalProvider.Logger = log.Get()
	globalProvider.Version = version.Get()

	return globalProvider
}

func Get() *ServiceProvider {
	if globalProvider == nil {
		return New()
	}
	return globalProvider
}
