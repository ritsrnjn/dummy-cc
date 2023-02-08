package dbentity

type Offer struct {
	OfferID        int64  `json:"offer_id"`
	AccountID      int64  `json:"account_id"`
	LimitType      string `json:"limit_type"`
	NewLimit       int    `json:"new_limit"`
	ActivationTime int64  `json:"activation_time"`
	ExpirationTime int64  `json:"expiration_time"`
	Status         string `json:"status"`
}

// write sql create table query for offer table, use accountId as Fk
