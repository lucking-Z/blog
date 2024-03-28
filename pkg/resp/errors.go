package resp

// Error 错误类
type Error struct {
	ErrNo  int    `json:"code"`
	ErrMsg string `json:"msg"`
}

func (e *Error) GetErrMsg() string {
	return e.ErrMsg
}

func (e *Error) String() string {
	return e.ErrMsg
}

func (e *Error) Error() string {
	return e.ErrMsg
}

//在这里添加业务处理错误码
var (
	ErrSuccess       = &Error{ErrNo: 0, ErrMsg: "成功"}
	ErrParams        = &Error{ErrNo: 1, ErrMsg: "参数错误"}
	ErrFailed        = &Error{ErrNo: 2, ErrMsg: "服务器内部错误"}
	ErrPassword      = &Error{ErrNo: 1000, ErrMsg: "账号或密码错误"}
	ErrAccountExists = &Error{ErrNo: 1001, ErrMsg: "账户已存在"}
	ErrLogin         = &Error{ErrNo: 1002, ErrMsg: "请登录"}
	//ERR_VERIFY_SIG = Error{ErrNo: 201, ErrMsg: "验证签名失败"}
)
