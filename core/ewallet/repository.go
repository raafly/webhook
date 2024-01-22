package ewallet

import "gorm.io/gorm"

type ewalletRepository interface {
	FindOneByTransactionID(transactionID string) (*ewallet, error)
	InsertOne(ewallet *requestEwallet) error
	UpdateOne(ewallet *requestEwallet) error
	ChangeStatus(transactionID string, status string) error
}

type ewalletRepositoryImpl struct {
	db *gorm.DB
}

func NewEwalletRepository(db *gorm.DB) ewalletRepository {
	return &ewalletRepositoryImpl{db}
}

func (r *ewalletRepositoryImpl) FindOneByTransactionID(transactionID string) (*ewallet, error) {
	var ew ewallet
	err := r.db.Where("id = ?", transactionID).Take(&ew).Error
	if err != nil {
		return nil, err
	}
	return &ew, nil
}

func (r *ewalletRepositoryImpl) InsertOne(ewallet *requestEwallet) error {
	return r.db.Create(ewallet).Error
}

func (r *ewalletRepositoryImpl) UpdateOne(ewallet *requestEwallet) error {
	return r.db.Save(ewallet).Error
}

func (r *ewalletRepositoryImpl) ChangeStatus(transactionID string, status string) error {
	return r.db.Model(&ewallet{}).Where("id = ?", transactionID).Update("status", status).Error
}