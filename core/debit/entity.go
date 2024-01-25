package debit

type debit struct {
	Amount			int			
	Currency		int			
	Type 			string
	ChannelCode 	string
	Success 		string
	Failture 		string
	Reusability		string
	CustomerID		string	`gorm:"primaryKey"`
	Description		string		
}

func (d *debit) TableName() string {
	return "debit"
}