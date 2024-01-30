package constans

type Response struct {
	Code	int			`json:"code"`
	Status  string		`json:"status"`
	Message	string		`json:"massage"`
	Result 	interface{}	`json:"data"`
}

func (r *Response) Error() string {
	return r.Message
}

func NewSuccess(massage string, data any) error {
	return &Response{
		Message: massage,
		Status: "OK",
		Code: 200,
		Result: data,
	}
}

func NewCreated(massage string) error {
	return &Response{
		Code: 201,
		Status: "Created",
		Message: massage,
	}
}

func NewNotFoundError(massage string) error {
	return &Response{
		Code: 404,
		Status: "NOT FOUND",
		Message: massage,
	}
}

func NewBadRequestError(massage string) error {
	return &Response{
		Code: 400,
		Status: "BAD REQUEST",
		Message: massage,
	}
}

func NewUnauthorizedError(massage string) error {
	return &Response{
		Code: 401,
		Status: "UNAUTHORIZED",
		Message: massage,
	}
}

func NewInternalServerError(massage string) error {
	return &Response{
		Code: 500,
		Status: "INTERNAL SERVER ERROR",
		Message: massage,
	}
}