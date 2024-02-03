package auth

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/raaafly/powerup-client-service-golang/config"
	"github.com/raaafly/powerup-client-service-golang/mail"
	"github.com/raaafly/powerup-client-service-golang/utils"
	"github.com/raaafly/powerup-client-service-golang/utils/constans"
)

type authService interface {
	insertOne(user *register) (string, error)
	login(user *login) (string, error)
	// confirmPassword(ctx context.Context, email string) error
	forgetPassword(email string) error
	resetPassword(req *resetPassword) error 
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

func (s *authServiceImpl) insertOne(user *register) (string, error) {
	err := s.validate.Validate(user)
	if err != nil {
		return "", s.validate.ValidationMessage(err)
	}

	/*
	ctx, cancel := context.WithTimeout(context.Background(), 30 *time.Second)
	defer cancel()
	if err = s.confirmPassword(ctx, user.Email); err != nil {
		return constans.NewBadRequestError("timeout for confirm password")
	}
	*/

	user.Password = s.pass.HashPassword(user.Password)
	user.UUID = uuid.NewString()

	err = s.port.insertOne(user)
	if err != nil {
		return "", constans.NewBadRequestError("account already exists")
	}  

	token := uuid.New().String()
	fmt.Println(token)

	return token, nil
}

func (s *authServiceImpl) login(user *login) (string, error) {
	err := s.validate.Validate(user)
	if err != nil {
		return "", s.validate.ValidationMessage(err)
	}

	result, err := s.port.findByEmail(user.Email)
	if err != nil {
		return  "", constans.NewNotFoundError("ID user not found")
	}
	
	err = s.pass.ComparePassword(result.Password, user.Password)
	if err != nil {
		return "", constans.NewBadRequestError("password not match")
	}

	token := uuid.New().String()

	return token, nil
}

/*
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
*/

func(s *authServiceImpl) forgetPassword(email string) error {
	sender := mail.NewGmailSender(
		config.NewAppConfig().Email.Sender, 
		config.NewAppConfig().Email.Adderss, 
		config.NewAppConfig().Email.Password,
	)

	to := []string{email}
	
	return sender.SendEmail(to, nil, nil)
}

func (s *authServiceImpl) resetPassword(req *resetPassword) error {
	err := s.port.resetPassword(req)
	if err != nil {
		return constans.NewInternalServerError("someathing when wrong")
	}
	return nil
}