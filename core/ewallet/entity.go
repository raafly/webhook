package ewallet

type ewallet struct {
	TransactionID string
	OrderID       string
	ExternalID    string
	Status        string
	Amount        string
	Actions       []struct {
		Name   string
		Method string
		URL    string
	}
	CreatedAt string
	UpdatedAt string
}
