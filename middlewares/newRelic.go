package middlewares

import (
	NewRelicService "github.com/henriqueholanda/widgets-api-docker/services/newrelic"
	"github.com/newrelic/go-agent"
	"net/http"
)

type NewRelicMiddleware struct {
	Next http.Handler
}

func (controller NewRelicMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if NewRelicService.Agent != nil {
		_, newRelicHandler := newrelic.WrapHandle(
			NewRelicService.Agent,
			request.URL.Path,
			http.Handler(controller.Next),
		)

		newRelicHandler.ServeHTTP(writer, request)

		return
	}

	controller.Next.ServeHTTP(writer, request)
}
