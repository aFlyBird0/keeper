package response

const SuccessCode = 200
const SuccessMsg = "success"

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(data any) *Response {
	return SuccessWithMsg(SuccessMsg, data)
}

func EmptySuccess() *Response {
	return Success(nil)
}

func SuccessWithMsg(msg string, data any) *Response {
	return &Response{
		Code: SuccessCode,
		Msg:  msg,
		Data: data,
	}
}

func Fail(bizError BizError) *Response {
	return FailRaw(bizError.Code, bizError.Msg)
}

func FailRaw(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
	}
}
