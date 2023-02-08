package response

// Account is the response for the create and get account endpoint
type Account struct {
	AccountID           int64  `json:"account_id"`
	CustomerID          string `json:"customer_id"`
	AccountLimit        int    `json:"account_limit"`
	PerTransactionLimit int    `json:"per_transaction_limit"`
}
