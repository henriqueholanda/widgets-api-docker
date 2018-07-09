package newrelic

import (
	"github.com/newrelic/go-agent"
	"net/http"
)

// Agent New Relic Agent
var Agent newrelic.Application

func init() {
	License := ""
	AppName := ""
	if License == "" {
		return
	}
	config := newrelic.NewConfig(AppName, License)
	config.ErrorCollector.IgnoreStatusCodes = []int{
		http.StatusNotFound,     // 404
		http.StatusUnauthorized, // 401
	}
	app, err := newrelic.NewApplication(config)
	if err != nil {
		return
	}

	Agent = app
}

func PopulateTransaction(writer http.ResponseWriter, request *http.Request) {
	if transaction, ok := writer.(newrelic.Transaction); ok {
		for key, value := range request.URL.Query() {
			transaction.AddAttribute(key, value[0])
		}
		for key, value := range request.Header {
			transaction.AddAttribute(key, value[0])
		}
	}
}
