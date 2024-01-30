package auth

type register struct {
	ID 					string	`json:"id"`
	Username			string	`json:"username" validate:"required,min=3,max=50"`
	Email 				string	`json:"email" validate:"required,email"`
	Password        	string `json:"password" validate:"required,min=8"`
    ConfirmPassword 	string `json:"confirm_password" validate:"eqfield=Password"`
}

type login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type response struct {
	Code		int 	`json:"code"`
	Status		string	`json:"status"`
	Message		string	`json:"message"`
	Data 		payload	`json:"data"`
}

type payload struct {
	UserID		string	`json:"user_id"`
	Username	string	`json:"username"`
	Token		string	`json:"token"`
}

type loginResponse struct {
	UserID 				string `json:"user_id"`
	AccessToken         string `json:"access_token"`
	AccessTokenExpired  int64  `json:"access_token_expired"`
	RefreshToken        string `json:"refresh_token"`
	RefreshTokenExpired int64  `json:"refresh_token_expired"`
}