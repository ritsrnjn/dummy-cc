package dbentity

type LimitOffer struct {
	OfferId        string `json:"offer_id"`
	AccountID      string `json:"account_id"`
	LimitType      string `json:"limit_type"`
	NewLimit       int    `json:"new_limit"`
	ActivationTime int    `json:"activation_time"`
	ExpirationTime int    `json:"expiration_time"`
}
