package response

const (
	NOT_FOUND_CODE    = 404
	BAD_REQUEST_CODE  = 400
	FORBIDDEN_CODE    = 403
	UNAUTHORIZED_CODE = 401
	SUCCESS_CODE      = 200
	UNDEFINED_CODE    = 500
)

const (
	NOT_FOUND_MSG     = "data not found"
	ALREADY_EXIST_MSG = "data already exist"
	UNAUTHORIZED_MSG  = "unauthorized"
	SUCCESS_MSG       = "ok"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *any   `json:"data,omitempty"`
}

var (
	NOT_FOUND     = new(NOT_FOUND_CODE, NOT_FOUND_MSG)
	ALREADY_EXIST = new(FORBIDDEN_CODE, ALREADY_EXIST_MSG)
	UNAUTHORIZED  = new(UNAUTHORIZED_CODE, UNAUTHORIZED_MSG)
)

func UNDEFINED(msg string) Response {
	return Response{
		Code:    UNDEFINED_CODE,
		Message: msg,
	}
}

func BAD_REQUEST(msg string) Response {
	return Response{
		Code:    BAD_REQUEST_CODE,
		Message: msg,
	}
}

func SUCCESS(data any) Response {
	return Response{
		Code:    SUCCESS_CODE,
		Message: SUCCESS_MSG,
		Data:    &data,
	}
}

func new(code int, msg string) Response {
	return Response{
		Code:    code,
		Message: msg,
	}
}

func (resp Response) Extract() (int, Response) {
	return resp.Code, resp
}
