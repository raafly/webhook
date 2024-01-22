package ewallet

type ewalletService interface {
	GetPaymentStatus(transactionID string) (*responePaymentStatus, error)
	InsertOne(ewallet *ewallet) error
	UpdateOne(ewallet *ewallet) error
	ChangeStatus(transactionID string, status string) error
}

type ewalletServiceImpl struct {
	ewalletRepository ewalletRepository
}

func NewEwalletService(ewalletRepository ewalletRepository) ewalletService {
	return &ewalletServiceImpl{ewalletRepository}
}

func (s *ewalletServiceImpl) GetPaymentStatus(transactionID string) (*responePaymentStatus, error) {
	res, err := s.ewalletRepository.FindOneByTransactionID(transactionID)
	if err != nil {
		return nil, err
	}

	return &responePaymentStatus{
		ID: res.ID,
		Status: res.Status,
		CustomerID: res.CustomerID,
		Created: res.Created,
		Updated: res.Update,
	}, nil
}

func (s *ewalletServiceImpl) InsertOne(ewallet *ewallet) error {
	return s.ewalletRepository.InsertOne(&requestEwallet{})
}

func (s *ewalletServiceImpl) UpdateOne(ewallet *ewallet) error {
	return s.ewalletRepository.UpdateOne(&requestEwallet{})
}

func (s *ewalletServiceImpl) ChangeStatus(transactionID string, status string) error {
	return s.ewalletRepository.ChangeStatus(transactionID, status)
}