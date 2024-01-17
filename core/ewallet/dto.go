package ewallet

type ewalletStatusResponse struct {
	TransactionID string `json:"transaction_id"`
	OrderID       string `json:"order_id"`
	ExternalID    string `json:"external_id"`
	Status        string `json:"status"`
	Amount        string `json:"amount"`
	Actions       []struct {
		Name   string `json:"name"`
		Method string `json:"method"`
		URL    string `json:"url"`
	} `json:"actions"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
