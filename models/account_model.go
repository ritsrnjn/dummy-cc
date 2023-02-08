package models

import (
	"ritsrnjn/dummy-cc/dto/dbentity"
	"ritsrnjn/dummy-cc/dto/request"
	"ritsrnjn/dummy-cc/sqldb"
	"ritsrnjn/dummy-cc/utils"
)

// CreateAccount creates a new account in the database.
func CreateAccount(createAccountRequest request.CreateAccount) (dbentity.Account, error) {
	currentTimestamp := utils.GetTmeStampInMs()

	newAccount := dbentity.Account{
		CustomerID:                    createAccountRequest.CustomerID,
		AccountLimit:                  createAccountRequest.AccountLimit,
		PerTransactionLimit:           createAccountRequest.PerTransactionLimit,
		LastAccountLimit:              createAccountRequest.AccountLimit,
		LastPerTransactionLimit:       createAccountRequest.PerTransactionLimit,
		AccountLimitUpdateTime:        currentTimestamp,
		PerTransactionLimitUpdateTime: currentTimestamp,
	}

	result, err := sqldb.Execute("INSERT INTO accounts(customer_id, account_limit, per_transaction_limit, last_account_limit, last_per_transaction_limit, account_limit_update_time, per_transaction_limit_update_time) VALUES(?, ?, ?, ?, ?, ?, ?)", newAccount.CustomerID, newAccount.AccountLimit, newAccount.PerTransactionLimit, newAccount.LastAccountLimit, newAccount.LastPerTransactionLimit, newAccount.AccountLimitUpdateTime, newAccount.PerTransactionLimitUpdateTime)
	if err != nil {
		return newAccount, err
	}

	newAccount.AccountID, err = result.LastInsertId()
	if err != nil {
		return newAccount, err
	}
	return newAccount, nil
}

// GetAccount gets an account from the database.
func GetAccount(accountID int64) (dbentity.Account, error) {
	var account dbentity.Account
	result, err := sqldb.Query("SELECT * FROM accounts WHERE account_id = ?", accountID)
	if err != nil {
		return account, err
	}

	if result.Next() {
		err = result.Scan(&account.AccountID, &account.CustomerID, &account.AccountLimit, &account.PerTransactionLimit, &account.LastAccountLimit, &account.LastPerTransactionLimit, &account.AccountLimitUpdateTime, &account.PerTransactionLimitUpdateTime)
		if err != nil {
			return account, err
		}
	}

	return account, nil
}

// Update the account limit in the database.
func UpdateAccountLimit(accountID int64, newAccountLimit int) error {
	_, err := sqldb.Execute("UPDATE accounts SET account_limit = ?, last_account_limit = account_limit, account_limit_update_time = ? WHERE account_id = ?", newAccountLimit, utils.GetTmeStampInMs(), accountID)
	if err != nil {
		return err
	}
	return nil
}

// Update the per transaction limit in the database.
func UpdatePerTransactionLimit(accountID int64, newPerTransactionLimit int) error {
	_, err := sqldb.Execute("UPDATE accounts SET per_transaction_limit = ?, last_per_transaction_limit = per_transaction_limit, per_transaction_limit_update_time = ? WHERE account_id = ?", newPerTransactionLimit, utils.GetTmeStampInMs(), accountID)
	if err != nil {
		return err
	}
	return nil
}
