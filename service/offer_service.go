package service

import (
	"errors"

	"ritsrnjn/dummy-cc/constants"
	"ritsrnjn/dummy-cc/dto/request"
	"ritsrnjn/dummy-cc/dto/response"
	"ritsrnjn/dummy-cc/models"
	"ritsrnjn/dummy-cc/utils"
)

func CreateOffer(createOfferRequest request.CreateOffer) (response.Offer, error) {
	// get account details
	account, err := models.GetAccount(createOfferRequest.AccountID)
	if err != nil {
		return response.Offer{}, err
	}

	// check if account exists
	if account.AccountID == 0 {
		return response.Offer{}, errors.New("Bad Request: Account does not exist")
	}

	// Validate if the offer is valid and can be created
	if !checkIfOfferCanBeCreated(createOfferRequest.LimitType, account.AccountLimit, account.PerTransactionLimit, createOfferRequest.NewLimit) {
		return response.Offer{}, errors.New("Bad Request: Offer cannot be created as new limit is not greater than current limit")
	}

	// if activation and expiration time was not provided in the request, set it to default values
	if createOfferRequest.ActivationTime == 0 && createOfferRequest.ExpirationTime == 0 {
		createOfferRequest.ActivationTime = utils.GetTmeStampInMs()
		createOfferRequest.ExpirationTime = createOfferRequest.ActivationTime + constants.DefaultExpirationPeriod
	}

	// call model layer
	newOffer, err := models.CreateOffer(createOfferRequest)

	if err != nil {
		return response.Offer{}, err
	}

	// return response
	return response.Offer{
		OfferID:        newOffer.OfferID,
		AccountID:      newOffer.AccountID,
		LimitType:      newOffer.LimitType,
		NewLimit:       newOffer.NewLimit,
		ActivationTime: newOffer.ActivationTime,
		ExpirationTime: newOffer.ExpirationTime,
		Status:         newOffer.Status,
	}, nil
}

func ListOffers(accountID, activeTime int64) ([]response.Offer, error) {

	// call model layer
	offers, err := models.ListOffers(accountID)

	if err != nil {
		return []response.Offer{}, err
	}

	// if activeTime was not present int the request, set it to current time
	if activeTime == 0 {
		activeTime = utils.GetTmeStampInMs()
	}

	// return response
	var offerResponses []response.Offer
	for _, offer := range offers {
		// add only if offer status PENDING and expiration time is greater than current time
		if offer.Status == constants.PendingStatus && offer.ExpirationTime >= activeTime && activeTime >= offer.ActivationTime {
			offerResponses = append(offerResponses, response.Offer{
				OfferID:        offer.OfferID,
				AccountID:      offer.AccountID,
				LimitType:      offer.LimitType,
				NewLimit:       offer.NewLimit,
				ActivationTime: offer.ActivationTime,
				ExpirationTime: offer.ExpirationTime,
				Status:         offer.Status,
			})
		}
	}

	return offerResponses, nil
}

func checkIfOfferCanBeCreated(limitType string, accountLimit, perTransactionLimit, newLimit int) bool {
	if limitType == constants.AccountLimit && accountLimit >= newLimit {
		return false
	} else if limitType == constants.PerTransactionLimit && perTransactionLimit >= newLimit {
		return false
	}
	return true
}
