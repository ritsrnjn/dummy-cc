package dbentity

type Account struct {
	AccountID                     string `json:"account_id"`
	CustomerID                    string `json:"customer_id"`
	AccountLimit                  int    `json:"account_limit"`
	PerTransactionLimit           int    `json:"per_transaction_limit"`
	LastAccountLimit              int    `json:"last_account_limit"`
	LastPerTransactionLimit       int    `json:"last_per_transaction_limit"`
	AccountLimitUpdateTime        int    `json:"account_limit_update_time"`
	PerTransactionLimitUpdateTime int    `json:"per_transaction_limit_update_time"`
}
