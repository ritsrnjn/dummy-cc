package handlers

import (
	"encoding/json"
	"net/http"

	"ritsrnjn/dummy-cc/apiresponse"
	"ritsrnjn/dummy-cc/dto/request"
	"ritsrnjn/dummy-cc/service"
)

// add hello handler function
func HelloFunc(writer http.ResponseWriter, request *http.Request) {
	apiresponse.SendOK(writer, "dummy-cc project")
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	// get request body
	var accountCreateRequest request.CreateAccount
	err := json.NewDecoder(r.Body).Decode(&accountCreateRequest)
	if err != nil {
		apiresponse.SendBadRequest(w, "Invalid request body")
		return
	}

	// validate request body
	err = accountCreateRequest.Validate()
	if err != nil {
		apiresponse.SendBadRequest(w, err.Error())
		return
	}

	// call service layer
	createAccountResponse, err := service.CreateAccount(accountCreateRequest)
	if err != nil {
		apiresponse.SendInternalServerError(w, err.Error())
		return
	}

	// send response
	apiresponse.SendOK(w, createAccountResponse)
}
