package handlers

import (
	"net/http"

	"ritsrnjn/dummy-cc/apiresponse"
)

// add hello handler function
func HelloFunc(writer http.ResponseWriter, request *http.Request) {
	apiresponse.SendOK(writer, "dummy-cc project")
}
