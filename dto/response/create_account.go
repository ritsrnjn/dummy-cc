package response

// CreateAccount is the response for the create account endpoint
type CreateAccount struct {
	AccountID           string `json:"account_id"`
	CustomerID          string `json:"customer_id"`
	AccountLimit        int    `json:"account_limit"`
	PerTransactionLimit int    `json:"per_transaction_limit"`
}
