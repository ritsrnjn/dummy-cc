package request

import (
	"errors"
	"ritsrnjn/dummy-cc/constants"
	"ritsrnjn/dummy-cc/utils"
)

type CreateOffer struct {
	AccountID      int64  `json:"account_id"`
	LimitType      string `json:"limit_type"`
	NewLimit       int    `json:"new_limit"`
	ActivationTime int64  `json:"activation_time"`
	ExpirationTime int64  `json:"expiration_time"`
}

// Validate function to validate the request body
func (c *CreateOffer) Validate() error {
	if c.AccountID == 0 {
		return errors.New("account_id is required")
	}
	if c.LimitType != constants.PerTransactionLimit && c.LimitType != constants.AccountLimit {
		return errors.New("limit_type is invalid")
	}
	if c.ExpirationTime < c.ActivationTime {
		return errors.New("expiration_time cannot be less than activation_time")
	}
	if c.ExpirationTime != 0 && c.ExpirationTime < utils.GetTmeStampInMs() {
		return errors.New("expiration_time cannot be less than current time")
	}

	return nil
}
