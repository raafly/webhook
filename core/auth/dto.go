package auth

type register struct {
	Username	string	`json:"username"`
	Email 		string	`json:"email"`
	Phone 		string	`json:"phone"`
	Password 	string	`json:"password"`
}

type login struct {
	Username	string	`json:"username"`
	Email 		string	`json:"email"`
	Password 	string	`json:"password"`	
}

type loginResponse struct {
	AccessToken         string `json:"access_token"`
	AccessTokenExpired  int64  `json:"access_token_expired"`
	RefreshToken        string `json:"refresh_token"`
	RefreshTokenExpired int64  `json:"refresh_token_expired"`
}