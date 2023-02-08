package handlers

import (
	"encoding/json"
	"net/http"

	"ritsrnjn/dummy-cc/apiresponse"
	"ritsrnjn/dummy-cc/dto/request"
	"ritsrnjn/dummy-cc/service"
	"ritsrnjn/dummy-cc/utils"

	"github.com/go-chi/chi"
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

func GetAccount(w http.ResponseWriter, r *http.Request) {
	// get accountID from path params
	accountID, err := utils.StringToInt64(chi.URLParam(r, "accountID"))

	if err != nil {
		apiresponse.SendBadRequest(w, "Invalid accountID")
		return
	}

	// call service layer
	getAccountResponse, err := service.GetAccount(accountID)
	if err != nil {
		apiresponse.SendInternalServerError(w, err.Error())
		return
	}

	if getAccountResponse.AccountID == 0 {
		apiresponse.SendNotFound(w, "Account not found")
		return
	}

	// send response
	apiresponse.SendOK(w, getAccountResponse)
}
