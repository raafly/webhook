package constans

type Response struct {
	Status  bool		`json:"status"`
	Code	int			`json:"code"`
	Message	string		`json:"massage"`
	Data 	interface{}	`json:"data"`
}

type ErrorResponse struct {
	Status  bool		`json:"status"`
	Code	int			`json:"code"`
	Message	string		`json:"massage"`
	Err 	interface{}	`json:"errors"`
}

func (r *ErrorResponse) Error() string {
	return r.Message
}

func (r *Response) Error() string {
	return r.Message
}

func NewSuccess(massage string, data any) error {
	return &Response{
		Message: massage,
		Status: true,
		Code: 200,
		Data: data,
	}
}

func NewCreated(massage string, data any) error {
	return &Response{
		Code: 201,
		Status: true,
		Message: massage,
		Data: data,
	}
}

func NewForbiddenError(message string) error {
	return &ErrorResponse{
		Status: false,
		Code: 403,
		Message: message,
	}
}

func NewNotFoundError(message string) error {
	return &ErrorResponse{
		Code: 404,
		Status: false,
		Message: message,
	}
}

func NewBadRequestError(message string) error {
	return &ErrorResponse{
		Code: 400,
		Status: false,
		Message: message,
	}
}

func NewUnauthorizedError(message string) error {
	return &ErrorResponse{
		Code: 401,
		Status: false,
		Message: message,
	}
}

func NewInternalServerError(message string) error {
	return &Response{
		Code: 500,
		Status: false,
		Message: message,
	}
}

func NewCustomError(code int, message string, data any) error {
	return &ErrorResponse{
		Status: false,
		Code: code,
		Message: message,
		Err: data,
	}
}