package classin

import "fmt"

// Error 定义ClassIn接口error
type Error struct {
	// 错误代码
	Code uint32

	// 错误信息
	Msg string
}

func NewError(code uint32, msg string) *Error {
	return &Error{Code: code, Msg: msg}
}

// NewServerError 创建系统错误
// 主要是接口网络无法访问，拦截非ClassIn接口error
func NewServerError(msg string) *Error {
	return NewError(100001, msg)
}

// Error 实现error接口
func (e *Error) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.Code, e.Msg)
}

// GetCode 获取错误代码
func (e *Error) GetCode() uint32 {
	return e.Code
}

// GetMsg 获取错误信息
func (e *Error) GetMsg() string {
	return e.Msg
}
