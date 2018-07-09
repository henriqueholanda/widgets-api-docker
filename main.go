package main

import (
	"net/http"
	"github.com/henriqueholanda/widgets-api-docker/handlers"
	"github.com/henriqueholanda/widgets-api-docker/middlewares"
	"github.com/henriqueholanda/widgets-api-docker/entities"
	"log"
)

func main() {
	http.Handle(
		entities.RootEndpoint,
		middlewares.NotFoundMiddleware{
			middlewares.NewRelicMiddleware{
				handlers.IndexHandler{},
			},
		},
	)
	http.Handle(
		entities.UsersEndpoint,
		middlewares.NotFoundMiddleware{
			middlewares.NewRelicMiddleware{
				middlewares.AuthenticationMiddleware{
					handlers.UsersHandler{},
				},
			},
		},
	)

	log.Fatal(http.ListenAndServe(":80", nil))
}
