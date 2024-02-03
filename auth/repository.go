package auth

import (
	"gorm.io/gorm"
)

type authRepository interface {
	insertOne(data *register) error
	findByEmail(ID string) (*User, error)
	resetPassword(data *resetPassword) error
}

type authRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) authRepository {
	return &authRepositoryImpl{db: db}
}

func (r *authRepositoryImpl) insertOne(data *register) error {
	user := &User{
		UUID: data.UUID,
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
	}

	return r.db.Create(user).Error
}

func (r *authRepositoryImpl) findByEmail(ID string) (*User, error) {
	var user User
	err := r.db.Where("email = ?", ID).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepositoryImpl) resetPassword(data *resetPassword) error {
	user := User {
		ID: data.UserID,
		Password: data.Password,
	}

	return r.db.Save(&user).Error
}