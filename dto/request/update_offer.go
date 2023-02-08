package request

import (
	"errors"
	"ritsrnjn/dummy-cc/constants"
)

type UpdateOffer struct {
	OfferID int64  `json:"offer_id"`
	Status  string `json:"status"`
}

// Validate function to validate the request body
func (u *UpdateOffer) Validate() error {
	if u.OfferID == 0 {
		return errors.New("offer_id is required")
	}
	if u.Status != constants.AcceptedStatus && u.Status != constants.RejectedStatus {
		return errors.New("status is invalid")
	}
	return nil
}
