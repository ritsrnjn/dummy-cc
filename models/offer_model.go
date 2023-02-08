package models

import (
	"ritsrnjn/dummy-cc/constants"
	"ritsrnjn/dummy-cc/dto/dbentity"
	"ritsrnjn/dummy-cc/dto/request"
	"ritsrnjn/dummy-cc/sqldb"
)

// CreateOffer creates a new offer in the database.
func CreateOffer(createOfferRequest request.CreateOffer) (dbentity.Offer, error) {
	newOffer := dbentity.Offer{
		AccountID:      createOfferRequest.AccountID,
		LimitType:      createOfferRequest.LimitType,
		NewLimit:       createOfferRequest.NewLimit,
		ActivationTime: createOfferRequest.ActivationTime,
		ExpirationTime: createOfferRequest.ExpirationTime,
		Status:         constants.PendingStatus,
	}

	result, err := sqldb.Execute("INSERT INTO offers(account_id, limit_type, new_limit, activation_time, expiration_time, status) VALUES(?, ?, ?, ?, ?, ?)", newOffer.AccountID, newOffer.LimitType, newOffer.NewLimit, newOffer.ActivationTime, newOffer.ExpirationTime, newOffer.Status)
	if err != nil {
		return newOffer, err
	}

	newOffer.OfferID, err = result.LastInsertId()
	if err != nil {
		return newOffer, err
	}
	return newOffer, nil
}

// ListOffers list all offers for a given account_id
func ListOffers(accountID int64) ([]dbentity.Offer, error) {
	var offers []dbentity.Offer
	result, err := sqldb.Query("SELECT * FROM offers WHERE account_id = ?", accountID)
	if err != nil {
		return offers, err
	}

	for result.Next() {
		var offer dbentity.Offer
		err = result.Scan(&offer.OfferID, &offer.AccountID, &offer.LimitType, &offer.NewLimit, &offer.ActivationTime, &offer.ExpirationTime, &offer.Status)
		if err != nil {
			return offers, err
		}
		offers = append(offers, offer)
	}

	return offers, nil
}
