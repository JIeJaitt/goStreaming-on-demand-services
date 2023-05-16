package defs

type Err struct {
	Error string `json:"error"`
	// 与状态码完全不一样, 是系统内部用来查每个 err 的方式
	// 例如 ErrorCode 001, ErrorCode 002
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{HttpSC: 400, Error: Err{Error: "Request body is not correct", ErrorCode: "001"}}
	ErrorNotAuthUser            = ErrorResponse{HttpSC: 401, Error: Err{Error: "User authentication failed.", ErrorCode: "002"}}
)
