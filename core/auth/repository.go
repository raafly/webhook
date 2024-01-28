package auth

import (
	"gorm.io/gorm"
)

type authRepository interface {
	insertOne(user *register) error
	findById(ID string) (*User, error)
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
		Phone: 	  user.Phone,
	}

	return r.db.Table("users").Create(data).Error
}

func (r *authRepositoryImpl) findById(ID string) (*User, error) {
	var user User
	err := r.db.Where("id = ?", ID).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}