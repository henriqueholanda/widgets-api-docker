package handlers

import "net/http"

type IndexHandler struct {
	handler http.Handler
}

// ServeHTTP Welcome Endpoint
// @Title Welcome Endpoint
// @Description Return a Welcome Message
// @Success 200 {object} string "Welcome Message"
// @router /index [get]
func (controller IndexHandler) ServeHTTP(writer http.ResponseWriter, response *http.Request) {
	writer.Write([]byte("Welcome to the Jungle!!!"))
}
