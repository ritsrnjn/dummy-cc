package service

import (
	"ritsrnjn/dummy-cc/dto/request"
	"ritsrnjn/dummy-cc/dto/response"
)

func CreateAccount(createAccountRequest request.CreateAccount) (response.CreateAccount, error) {

	// return response
	return response.CreateAccount{
		AccountID:           createAccountRequest.CustomerID + "123",
		CustomerID:          createAccountRequest.CustomerID,
		AccountLimit:        createAccountRequest.AccountLimit,
		PerTransactionLimit: createAccountRequest.PerTransactionLimit,
	}, nil
}
