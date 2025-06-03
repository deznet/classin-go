package classin

// ICommonResp 接口返回通用接口
type ICommonResp interface {
	// IsOK 接口访问/执行是否成功
	IsOK() bool

	// GetError 如果接口访问/执行不成功，获取错误信息
	GetError() *Error
}

// CommonRespV1 v1版本返回通用Response
type CommonRespV1 struct {
	ErrorInfo struct {
		ErrNo  uint32 `json:"errno"`
		ErrMsg string `json:"error"`
	} `json:"error_info"`
}

// CommonRespV2 v2版本返回通用Response
type CommonRespV2 struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// IsOK 请求是否成功
func (c *CommonRespV1) IsOK() bool {
	return c.ErrorInfo.ErrNo == 1
}

// GetError 返回错误
func (c *CommonRespV1) GetError() *Error {
	return NewError(c.ErrorInfo.ErrNo, c.ErrorInfo.ErrMsg)
}

// IsOK 请求是否成功
func (c *CommonRespV2) IsOK() bool {
	return c.Code == 1
}

// GetError 返回错误
func (c *CommonRespV2) GetError() *Error {
	return NewError(uint32(c.Code), c.Msg)
}
