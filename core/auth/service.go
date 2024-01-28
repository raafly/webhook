package auth

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/raafly/webhook/constans"
	"github.com/raafly/webhook/utils"
)

type authService interface {
	insertOne(data *register) error
	login(data *login) (*loginResponse, error)
}

type authServiceImpl struct {
	port 		authRepository
	pass 		utils.Password
	validate 	*validator.Validate
}

func NewAuthService(port authRepository, pass utils.Password, validate *validator.Validate) authService {
	return &authServiceImpl{
		port: port,
		pass: pass,
		validate: validate,
	}
}

func (s *authServiceImpl) insertOne(data *register) error {
	// err := s.validate.Struct(data)
	// if err != nil {
	// 	return constans.NewBadRequestError("missing field")
	// }

	hashPassword := s.pass.HashPassword(data.Password)
	uuid := uuid.NewString()

	data.Password = hashPassword
	data.ID = uuid

	if err := s.port.insertOne(data); err != nil {
		return constans.NewBadRequestError("duplicated data")
	}

	return nil
}

func (s *authServiceImpl) login(data *login) (*loginResponse, error) {
	result, err := s.port.findById(data.Username)
	if err != nil {
		return nil, constans.NewNotFoundError("ID user not found")
	}
	
	err = s.pass.ComparePassword(result.Password, data.Password)
	if err != nil {
		return nil, constans.NewBadRequestError("password not match")
	}

	token, tokenExp, _ := utils.NewGenerateToken().GenerateAccessToken(result.ID, result.Email, result.Username)
	refresh, refresExp, _ := utils.NewGenerateToken().GenerateRefreshToken(result.ID, result.Email, result.Username)

	return &loginResponse{
		AccessToken: token,	
		AccessTokenExpired: tokenExp,
		RefreshToken: refresh,
		RefreshTokenExpired: refresExp,
	}, nil
}