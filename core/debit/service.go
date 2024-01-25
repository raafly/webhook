package debit

type debitService interface {
	createOne()
}

type debitServiceImpl struct {
}

func NewDebitService() debitService {
	return &debitServiceImpl{}
}


func (d *debitServiceImpl) createOne() {
	
}