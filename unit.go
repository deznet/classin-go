package classin

import "strconv"

// CreateUnitReq 创建单元请求参数
type CreateUnitReq struct {
	//班级（课程）ID
	CourseId int64 `json:"courseId"`

	//单元名称,长度不超过50字。注：课程下不支持创建同名单元
	Name string `json:"name"`

	//是否发布 0 = 草稿，2 = 已发布（显示）
	PublishFlag uint8 `json:"publishFlag"`

	//单元介绍,不传默认为空
	Content string `json:"content"`
}

// CreateUnitResp 创建单元返回参数
type CreateUnitResp struct {
	CommonRespV2
	Data struct {
		Name   string `json:"name"`   //单元名
		UnitId int64  `json:"unitId"` //单元ID
	} `json:"data"`
}

// UpdateUnitReq 编辑单元请求参数
type UpdateUnitReq struct {
	CreateUnitReq

	//单元ID
	UnitId int64 `json:"unitId"`
}

// DeleteUnitReq 删除单元请求参数
type DeleteUnitReq struct {
	//单元ID
	UnitId int64 `json:"unitId"`

	//班级（课程）ID
	CourseId int64 `json:"courseId"`
}

// CreateUnit 创建单元
// 创建单元的时候，不支持创建名称重复的单元。如课程下有同名单元，不管是否草稿状态，均会报 单元已存在 的错误。
// @see https://docs.eeo.cn/api/zh-hans/LMS/createUnit.html
// @param req *CreateUnitReq
// @result int64 返回单元id
// @result *Error
func (c *Client) CreateUnit(req *CreateUnitReq) (int64, *Error) {
	uri := "/lms/unit/create"
	signData := make(map[string]string, 0)
	signData["courseId"] = strconv.FormatInt(req.CourseId, 10)
	signData["name"] = req.Name
	signData["content"] = req.Content
	signData["publishFlag"] = strconv.FormatUint(uint64(req.PublishFlag), 10)
	var result CreateUnitResp
	err := c.httpPostV2(uri, req, signData, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.UnitId, nil
}

// UpdateUnit 编辑单元
// @see https://docs.eeo.cn/api/zh-hans/LMS/updateUnit.html
// @param req *UpdateUnitReq
// @result *Error
func (c *Client) UpdateUnit(req *UpdateUnitReq) *Error {
	uri := "/lms/unit/update"
	signData := make(map[string]string, 0)
	signData["courseId"] = strconv.FormatInt(req.CourseId, 10)
	signData["name"] = req.Name
	signData["content"] = req.Content
	signData["publishFlag"] = strconv.FormatUint(uint64(req.PublishFlag), 10)
	signData["unitId"] = strconv.FormatInt(req.UnitId, 10)
	var result CommonRespV2
	err := c.httpPostV2(uri, req, signData, &result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUnit 删除单元
// @see https://docs.eeo.cn/api/zh-hans/LMS/deleteUnit.html
// @param req *DeleteUnitReq
// @result *Error
func (c *Client) DeleteUnit(req *DeleteUnitReq) *Error {
	uri := "/lms/unit/delete"
	signData := make(map[string]string, 0)
	signData["courseId"] = strconv.FormatInt(req.CourseId, 10)
	signData["unitId"] = strconv.FormatInt(req.UnitId, 10)
	var result CommonRespV2
	err := c.httpPostV2(uri, req, signData, &result)
	if err != nil {
		return err
	}
	return nil
}
