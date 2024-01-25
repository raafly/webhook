package debit

type debitRequest struct {
	Amount			int				`json:"amount"`
	Currency		int				`json:"currency"`
	Method 			paymentMethod	`json:"payment_method"`
	CustomerID		string			`json:"customer_id"`
	Description		string			`json:"description"`
}

type paymentMethod struct {
	Type 		string		`json:"type"`
	DirecDebit	direcDebit	`json:"direct_debit"`	
}

type direcDebit struct {
	ChannelCode 		string				`json:"channel_code"`
	ChannelProperties	channelProperties	`json:"channel_properties"`
	Reusability			string				`json:"reusability"`
}

type channelProperties struct {
	Success 	string	`json:"success_return_url"`
	Failture	string	`json:"failture_return_url"`
}