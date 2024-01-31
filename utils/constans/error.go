package constans

type Response struct {
	Status  bool		`json:"status"`
	Code	int			`json:"code"`
	Message	string		`json:"massage"`
	Data 	interface{}	`json:"data"`
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

func NewCreated(massage string) error {
	return &Response{
		Code: 201,
		Status: true,
		Message: massage,
	}
}

func NewForbiddenError(message string) error {
	return &Response{
		Status: false,
		Code: 403,
		Message: message,
	}
}

func NewNotFoundError(message string) error {
	return &Response{
		Code: 404,
		Status: false,
		Message: message,
	}
}

func NewBadRequestError(message string) error {
	return &Response{
		Code: 400,
		Status: false,
		Message: message,
	}
}

func NewUnauthorizedError(message string) error {
	return &Response{
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