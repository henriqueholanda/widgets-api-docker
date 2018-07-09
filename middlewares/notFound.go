package middlewares

import (
	"net/http"
	"github.com/henriqueholanda/widgets-api-docker/entities"
)

type NotFoundMiddleware struct {
	Next http.Handler
}

func (controller NotFoundMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	routes := entities.Routes

	for _, route := range routes {
		if route.Path == path {
			controller.Next.ServeHTTP(writer, request)
			return
		}
	}

	http.NotFound(writer, request)
}
