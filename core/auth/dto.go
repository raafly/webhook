package auth

type register struct {
	ID 			string	`json:"id"`
	Username	string	`json:"username" validate:"required, min=5"`
	Email 		string	`json:"email" validate:"required, email"`
	Phone 		string	`json:"phone" validate:"required, min=12, max=12"`
	Password 	string	`json:"password" validate:"required, min=8"`
}

type login struct {
	Username	string	`json:"username" validate:"required, min=5"`
	Email 		string	`json:"email" validate:"required, email"`
	Password 	string	`json:"password" validate:"required, min=8"`	
}

type loginResponse struct {
	AccessToken         string `json:"access_token"`
	AccessTokenExpired  int64  `json:"access_token_expired"`
	RefreshToken        string `json:"refresh_token"`
	RefreshTokenExpired int64  `json:"refresh_token_expired"`
}