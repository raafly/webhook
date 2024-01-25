package auth

import "gorm.io/gorm"

type authRepository interface {
	insertOne(data *register) error
	findById(ID string) (*User, error)
}

type authRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) authRepository {
	return &authRepositoryImpl{db: db}
}

func (r *authRepositoryImpl) insertOne(data *register) error {
	return r.db.Create(data).Error
}

func (r *authRepositoryImpl) findById(ID string) (*User, error) {
	var user User
	err := r.db.Where("id = ?", ID).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}