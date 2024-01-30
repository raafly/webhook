package auth

import (
	"gorm.io/gorm"
)

type authRepository interface {
	insertOne(user *register) error
	findByEmail(ID string) (*User, error)
}

type authRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) authRepository {
	return &authRepositoryImpl{db: db}
}

func (r *authRepositoryImpl) insertOne(user *register) error {
	data := &User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	return r.db.Table("users").Create(data).Error
}

func (r *authRepositoryImpl) findByEmail(ID string) (*User, error) {
	var user User
	err := r.db.Where("email = ?", ID).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}