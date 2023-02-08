package dbentity

type LimitOffer struct {
	OfferId        string `json:"offer_id"`
	AccountID      string `json:"account_id"`
	LimitType      string `json:"limit_type"`
	NewLimit       int    `json:"new_limit"`
	ActivationTime string `json:"activation_time"`
	ExpirationTime string `json:"expiration_time"`
}
