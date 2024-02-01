package auth

import (
	"gorm.io/gorm"
)

type authRepository interface {
	insertOne(data *register) (string, error)
	findByEmail(ID string) (*user, error)
}

type authRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) authRepository {
	return &authRepositoryImpl{db: db}
}

func (r *authRepositoryImpl) insertOne(data *register) (string, error) {
	user := &user{
		UUID: data.UUID,
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
	}

	return user.UUID ,r.db.Create(user).Error
}

func (r *authRepositoryImpl) findByEmail(ID string) (*user, error) {
	var user user
	err := r.db.Where("email = ?", ID).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}