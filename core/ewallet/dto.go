package ewallet

import "time"

type responePaymentStatus struct {
	ID          string    `json:"id"`
	Status      string    `json:"status"`
	CustomerID  string    `json:"customer_id"`
	ReferenceID string    `json:"reference_id"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

type requestEwallet struct {
	ID			string	`json:"id"`
	Status		string	`json:"status"`
	CustomerID	string	`json:"customer_id"`
}