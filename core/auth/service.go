package auth

import (
	"github.com/raafly/webhook/constans"
	"github.com/raafly/webhook/utils"
)

type authService interface {
	insertOne(data *register) error
	login(data *login) (*loginResponse, error)
}

type authServiceImpl struct {
	port authRepository
	pass utils.Password
}

func NewAuthService(port authRepository, pass utils.Password) authService {
	return &authServiceImpl{
		port: port,
		pass: pass,
	}
}

func (s *authServiceImpl) insertOne(data *register) error {
	if err := s.port.insertOne(data); err != nil {
		return constans.NewBadRequestError("cannot insert data, please check again")
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

	token, exp, err := utils.NewGenerateToken().GenerateAccessToken(result.ID, result.Email, result.Username)
	refresh, exp, err := utils.NewGenerateToken().GenerateRefreshToken(result.ID, result.Email, result.Username)

	return &loginResponse{
		AccessToken: token,	
		AccessTokenExpired: exp,
		RefreshToken: refresh,
		RefreshTokenExpired: exp,
	}, nil
}