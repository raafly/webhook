package auth

type register struct {
	UUID            string `json:"uuid"`
	Email           string `json:"email" validate:"required,email"`
	Name            string `json:"name" validate:"required,min=3,max=50"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"eqfield=Password"`
}

type login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type forgetPassword struct {
	Email string `json:"email"`
}

type resetPassword struct {
	UserID			string `json:"user_id" valdate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"eqfield=Password"`	
}

type response struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    data   `json:"data"`
}

type data struct {
	Token string `json:"token"`
}
