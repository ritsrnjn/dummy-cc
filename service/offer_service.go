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

func UpdateOffer(updateOfferRequest request.UpdateOffer) error {

	// check if status is valid
	if updateOfferRequest.Status != constants.AcceptedStatus && updateOfferRequest.Status != constants.RejectedStatus {
		return errors.New("Bad Request: Invalid status")
	}

	// get offer details
	offer, err := models.GetOffer(updateOfferRequest.OfferID)
	if err != nil {
		return err
	}

	// check if offer exists
	if offer.OfferID == 0 {
		return errors.New("Bad Request: Offer does not exist")
	}

	// check if offer is already accepted or rejected
	if offer.Status != constants.PendingStatus {
		return errors.New("Bad Request: Offer is already accepted or rejected")
	}

	// check if offer is expired
	if offer.ExpirationTime < utils.GetTmeStampInMs() {
		return errors.New("Bad Request: Offer is expired")
	}

	// check if offer is active
	if offer.ActivationTime > utils.GetTmeStampInMs() {
		return errors.New("Bad Request: Offer is not active yet")
	}

	// get account details
	account, err := models.GetAccount(offer.AccountID)
	if err != nil {
		return err
	}

	// check if valid offer according to current account limit
	if !checkIfOfferCanBeCreated(offer.LimitType, account.AccountLimit, account.PerTransactionLimit, offer.NewLimit) {
		return errors.New("Bad Request: Offer cannot be accepted as new limit is not greater than current limit")
	}

	// call model layer to update offer status
	err = models.UpdateOffer(updateOfferRequest)
	if err != nil {
		return err
	}

	// call model layer to update limit in account table if the offer was accepted
	if updateOfferRequest.Status == constants.AcceptedStatus && offer.LimitType == constants.AccountLimit {
		err = models.UpdateAccountLimit(offer.AccountID, offer.NewLimit)
	} else if updateOfferRequest.Status == constants.AcceptedStatus && offer.LimitType == constants.PerTransactionLimit {
		err = models.UpdatePerTransactionLimit(offer.AccountID, offer.NewLimit)
	}

	if err != nil {
		return err
	}

	return nil
}
