package middlewares

import (
	"net/http"

	"ritsrnjn/dummy-cc/apiresponse"
	"ritsrnjn/dummy-cc/constants"
)

// Middleware function to check if the request has a valid token
func AuthTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		tokenString := request.Header.Get("Authorization")

		if tokenString == constants.EmptyString {
			// return error if token is not valid
			apiresponse.SendUnauthorized(writer, "Token is required")
			return
		}
		// call the next handler
		next.ServeHTTP(writer, request)
	})
}
