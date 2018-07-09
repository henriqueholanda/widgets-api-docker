package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/newrelic/go-agent"
	"net/http"
	"strings"
)

type AuthenticationMiddleware struct {
	Next http.Handler
}

func (controller AuthenticationMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	authHeader := request.Header.Get("Authorization")
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})

	if err == nil && token.Valid {
		controller.Next.ServeHTTP(writer, request)
		return
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if transaction, ok := writer.(newrelic.Transaction); ok {
			transaction.AddAttribute("Unauthorized-Reason", ve.Error())
		}
		http.Error(writer, ve.Error(), http.StatusUnauthorized)
		return
	}

	if transaction, ok := writer.(newrelic.Transaction); ok {
		transaction.AddAttribute("Unauthorized-Reason", "Unknown")
	}
	http.Error(writer, "Unknown", http.StatusUnauthorized)
	return
}
