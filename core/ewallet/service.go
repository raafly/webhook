package ewallet

type ewalletService interface {
	GetPaymentStatus(transactionID string) (*ewalletStatusResponse, error)
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

func (s *ewalletServiceImpl) GetPaymentStatus(transactionID string) (*ewalletStatusResponse, error) {
	return s.ewalletRepository.FindOneByTransactionID(transactionID)
}

func (s *ewalletServiceImpl) InsertOne(ewallet *ewallet) error {
	return s.ewalletRepository.InsertOne(&ewalletStatusResponse{})
}

func (s *ewalletServiceImpl) UpdateOne(ewallet *ewallet) error {
	return s.ewalletRepository.UpdateOne(&ewalletStatusResponse{})
}

func (s *ewalletServiceImpl) ChangeStatus(transactionID string, status string) error {
	return s.ewalletRepository.ChangeStatus(transactionID, status)
}
