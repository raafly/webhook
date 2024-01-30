package auth

import (
	"context"
	// "time"

	"github.com/google/uuid"
	"github.com/raafly/webhook/config"
	"github.com/raafly/webhook/core/mail"
	"github.com/raafly/webhook/utils"
	"github.com/raafly/webhook/utils/constans"
)

type authService interface {
	insertOne(user *register) error
	login(user *login) (*loginResponse, error)
	confirmPassword(ctx context.Context, email string) error
}

type authServiceImpl struct {
	port 		authRepository
	pass 		utils.Password
	validate 	constans.IValidation 	
}

func NewAuthService(port authRepository, pass utils.Password, validate constans.IValidation) authService {
	return &authServiceImpl{
		port: port,
		pass: pass,
		validate: validate,
	}
}

func (s *authServiceImpl) insertOne(user *register) error {
	err := s.validate.Validate(user)
	if err != nil {
		return s.validate.ValidationMessage(err)
	}

	// ctx, cancel := context.WithTimeout(context.Background(), 30 *time.Second)
	// defer cancel()
	// if err = s.confirmPassword(ctx, user.Email); err != nil {
	// 	return constans.NewBadRequestError("timeout for confirm password")
	// }

	hashPassword := s.pass.HashPassword(user.Password)
	uuid := uuid.NewString()

	user.Password = hashPassword
	user.ID = uuid

	if err := s.port.insertOne(user); err != nil {
		return constans.NewBadRequestError("account already exists")
	}

	return nil
}

func (s *authServiceImpl) login(user *login) (*loginResponse, error) {
	err := s.validate.Validate(user)
	if err != nil {
		return nil, s.validate.ValidationMessage(err)
	}

	result, err := s.port.findByEmail(user.Email)
	if err != nil {
		return nil, constans.NewNotFoundError("ID user not found")
	}
	
	err = s.pass.ComparePassword(result.Password, user.Password)
	if err != nil {
		return nil, constans.NewBadRequestError("password not match")
	}

	token, tokenExp, _ := utils.NewGenerateToken().GenerateAccessToken(result.ID, result.Email, result.Username)
	refresh, refresExp, _ := utils.NewGenerateToken().GenerateRefreshToken(result.ID, result.Email, result.Username)

	return &loginResponse{
		UserID: result.ID,
		AccessToken: token,	
		AccessTokenExpired: tokenExp,
		RefreshToken: refresh,
		RefreshTokenExpired: refresExp,
	}, nil
}

func (s *authServiceImpl) confirmPassword(ctx context.Context, email string) error {
	sender := mail.NewGmailSender(
		config.NewAppConfig().Email.Sender, 
		config.NewAppConfig().Email.Adderss, 
		config.NewAppConfig().Email.Password,
	)

	subject := "[POWERUP] Email confirmation"
	content := `
	<h1>Hello!</h1>
	<br><br>
	<p>
		A sign in attempt required further verifacation because we did not recognize your divice. To complate the sing in,
		enter the verifacation code on the unrecognized device.<br><br>

		Verifacation code: {token}
	</p>
	`
	to := []string{"tes@gmail.com"}
	
	return sender.SendEmail(subject, content, to, nil, nil)
}