package service

import (
	"ritsrnjn/dummy-cc/dto/request"
	"ritsrnjn/dummy-cc/dto/response"
	"ritsrnjn/dummy-cc/models"
)

func CreateAccount(createAccountRequest request.CreateAccount) (response.Account, error) {
	// call model layer
	newAccount, err := models.CreateAccount(createAccountRequest)

	if err != nil {
		return response.Account{}, err
	}

	// return response
	return response.Account{
		AccountID:           newAccount.AccountID,
		CustomerID:          newAccount.CustomerID,
		AccountLimit:        newAccount.AccountLimit,
		PerTransactionLimit: newAccount.PerTransactionLimit,
	}, nil
}

func GetAccount(accountID int64) (response.Account, error) {

	// call model layer
	account, err := models.GetAccount(accountID)

	if err != nil {
		return response.Account{}, err
	}

	// return response
	return response.Account{
		AccountID:           account.AccountID,
		CustomerID:          account.CustomerID,
		AccountLimit:        account.AccountLimit,
		PerTransactionLimit: account.PerTransactionLimit,
	}, nil
}
