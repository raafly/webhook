package ewallet

import "gorm.io/gorm"

type ewalletRepository interface {
	FindOneByTransactionID(transactionID string) (*ewalletStatusResponse, error)
	InsertOne(ewallet *ewalletStatusResponse) error
	UpdateOne(ewallet *ewalletStatusResponse) error
	ChangeStatus(transactionID string, status string) error
}

type ewalletRepositoryImpl struct {
	db *gorm.DB
}

func NewEwalletRepository(db *gorm.DB) ewalletRepository {
	return &ewalletRepositoryImpl{db}
}

func (r *ewalletRepositoryImpl) FindOneByTransactionID(transactionID string) (*ewalletStatusResponse, error) {
	var ew ewalletStatusResponse
	err := r.db.Where("transaction_id = ?", transactionID).First(&ew).Error
	if err != nil {
		return nil, err
	}
	return &ew, nil
}

func (r *ewalletRepositoryImpl) InsertOne(ewallet *ewalletStatusResponse) error {
	return r.db.Create(ewallet).Error
}

func (r *ewalletRepositoryImpl) UpdateOne(ewallet *ewalletStatusResponse) error {
	return r.db.Save(ewallet).Error
}

func (r *ewalletRepositoryImpl) ChangeStatus(transactionID string, status string) error {
	return r.db.Model(&ewallet{}).Where("transaction_id = ?", transactionID).Update("status", status).Error
}
