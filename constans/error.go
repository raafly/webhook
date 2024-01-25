package constans

type Response struct {
	Code	int			`json:"code"`
	Message	string		`json:"massage"`
	Result 	interface{}	`json:"result"`
}

func (r *Response) Error() string {
	return r.Message
}

func NewSuccess(massage string, data any) error {
	return &Response{
		Message: massage,
		Code: 200,
		Result: data,
	}
}

func NewCreated(massage string) error {
	return &Response{
		Message: massage,
		Code: 201,
	}
}

func NewNotFoundError(massage string) error {
	return &Response{
		Message: massage,
		Code: 404,
	}
}

func NewBadRequestError(massage string) error {
	return &Response{
		Message: massage,
		Code: 400,
	}
}

func NewUnauthorizedError(massage string) error {
	return &Response{
		Message: massage,
		Code: 401,
	}
}

func NewInternalServerError(massage string) error {
	return &Response{
		Message: massage,
		Code: 500,
	}
}