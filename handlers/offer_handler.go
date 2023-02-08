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

// CreateOffer is a handler function to create an offer
func CreateOffer(w http.ResponseWriter, r *http.Request) {
	// get request body
	var createOfferRequest request.CreateOffer
	err := json.NewDecoder(r.Body).Decode(&createOfferRequest)
	if err != nil {
		apiresponse.SendBadRequest(w, "Invalid request body")
		return
	}

	// validate request body
	err = createOfferRequest.Validate()
	if err != nil {
		apiresponse.SendBadRequest(w, err.Error())
		return
	}

	// call service layer
	createOfferResponse, err := service.CreateOffer(createOfferRequest)
	if err != nil {
		apiresponse.SendInternalServerError(w, err.Error())
		return
	}

	// send response
	apiresponse.SendOK(w, createOfferResponse)
}

// ListOffers is a handler function to list all offers
func ListOffers(w http.ResponseWriter, r *http.Request) {
	// get accountID from path params
	accountID, err := utils.StringToInt64(chi.URLParam(r, "accountID"))

	if err != nil {
		apiresponse.SendBadRequest(w, "Invalid accountID")
		return
	}

	// get activeTime from query params
	activeTime, _ := utils.StringToInt64(r.URL.Query().Get("activeTime"))

	// call service layer
	listOffersResponse, err := service.ListOffers(accountID, activeTime)

	if err != nil && utils.Contains(err.Error(), "Bad Request") {
		apiresponse.SendBadRequest(w, err.Error())
		return
	}

	if err != nil {
		apiresponse.SendInternalServerError(w, err.Error())
		return
	}

	// send response
	apiresponse.SendOK(w, listOffersResponse)
}
