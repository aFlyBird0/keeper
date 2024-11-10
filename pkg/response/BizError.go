package response

type BizError struct {
	Code int
	Msg  string
}

var (
	ExampleError = BizError{Code: 10000, Msg: "示例错误（仅仅用于教学）"}
	ParamError   = BizError{Code: 10001, Msg: "参数错误"}
	NoAuthority  = BizError{Code: 10002, Msg: "无权限"}
	DatabseError = BizError{Code: 10003, Msg: "数据库错误"}
)
