package request

import (
	"errors"

	"ritsrnjn/dummy-cc/constants"
)

type CreateAccount struct {
	CustomerID          string `json:"customer_id"`
	AccountLimit        int    `json:"account_limit"`
	PerTransactionLimit int    `json:"per_transaction_limit"`
}

// Validate function to validate the request body
func (c *CreateAccount) Validate() error {
	if c.CustomerID == constants.EmptyString {
		return errors.New("customer_id is required")
	}
	return nil
}
