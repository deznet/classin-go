package classin

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"strconv"
)

// GetLoginLinkedReq 获取唤醒客户端并进入教室链接请求参数
type GetLoginLinkedReq struct {
	// 用户 UID
	// 必填
	Uid int64

	//密钥有效时长（单位：秒）
	//默认为86400秒
	//非必填
	LifeTime int32

	//课程 ID
	//必填
	CourseId int64

	//课节 ID
	//必填
	ClassId int64

	//平台标志
	//1代表 Windows/Mac OS 端；2代表 iOS 移动端；3代表 Android
	//默认为1，代表 PC 端
	//非必填
	DeviceType int32
}

// GetLoginLinkedResp 获取唤醒客户端并进入教室链接返回参数
type GetLoginLinkedResp struct {
	CommonRespV1

	//可在网页登录ClassIn客户端的链接
	Data string `json:"data"`
}

// GetWebcastUrlReq 获取课程直播/回放播放器地址请求参数
type GetWebcastUrlReq struct {
	//课程 ID
	//必填
	CourseId int64

	//课节 ID
	//必填
	ClassId int64
}

// GetWebcastUrlResp 获取课程直播/回放播放器地址返回参数
type GetWebcastUrlResp struct {
	CommonRespV1

	//返回直播回放页面地址
	Data string `json:"data"`
}

// GetNoLoginWebcastUrlReq 机构直播聊天室免二次登录url请求参数
type GetNoLoginWebcastUrlReq struct {
	GetWebcastUrlReq

	//长度小于32个字符,手机号
	//必填
	Account string

	//昵称，长度小于16个字符
	//必填
	Nickname string
}

// GetLoginLinked 获取唤醒客户端并进入教室链接
// @see https://docs.eeo.cn/api/zh-hans/getLoginLinked.html
// @param req *GetLoginLinkedReq
// @result string 进入教室链接
// @result *Error
func (c *Client) GetLoginLinked(req *GetLoginLinkedReq) (string, *Error) {
	formData := make(map[string]string)
	formData["uid"] = strconv.FormatInt(req.Uid, 10)
	formData["lifeTime"] = strconv.FormatInt(int64(req.LifeTime), 10)
	formData["courseId"] = strconv.FormatInt(req.CourseId, 10)
	formData["classId"] = strconv.FormatInt(req.ClassId, 10)
	if req.DeviceType == 0 {
		req.DeviceType = 1
	}
	formData["deviceType"] = strconv.FormatInt(int64(req.DeviceType), 10)
	uri := "/course.api.php?action=getLoginLinked"
	var result GetLoginLinkedResp
	err := c.httpPostV1(uri, formData, &result)
	if err != nil {
		return "", err
	}
	return result.Data, nil
}

// GetWebcastUrl 获取课程直播/回放播放器地址
// 返回第1个参数为网址，第二个为courseKey，第三个为lessonid
// @see https://docs.eeo.cn/api/zh-hans/broadcast/getWebcastUrl.html
// @param req *GetWebcastUrlReq
// @result string 直播回放页面地址
// @result *Error
func (c *Client) GetWebcastUrl(req *GetWebcastUrlReq) (string, *Error) {
	formData := make(map[string]string)
	formData["courseId"] = strconv.FormatInt(req.CourseId, 10)
	formData["classId"] = strconv.FormatInt(req.ClassId, 10)
	uri := "/course.api.php?action=getWebcastUrl"
	var result GetWebcastUrlResp
	err := c.httpPostV1(uri, formData, &result)
	if err != nil {
		return "", err
	}
	return result.Data, nil
}

// GetNoLoginWebcastUrl 直播聊天室免二次登录网址
// @see https://docs.eeo.cn/api/zh-hans/broadcast/getWebcastUrl.html
func (c *Client) GetNoLoginWebcastUrl(req *GetNoLoginWebcastUrlReq) (string, *Error) {
	wu, err := c.GetWebcastUrl(&GetWebcastUrlReq{
		CourseId: req.CourseId,
		ClassId:  req.ClassId,
	})
	if err != nil {
		return "", err
	}
	wu1, err1 := url.Parse(wu)
	if err1 != nil {
		return "", NewServerError("返回网址无法解析")
	}
	m, err1 := url.ParseQuery(wu1.RawQuery)
	if err1 != nil {
		return "", NewServerError("无法解析网址参数")
	}
	str := c.Secret + m.Get("courseKey") + req.Account + req.Nickname
	h := md5.New()
	h.Write([]byte(str))
	code := hex.EncodeToString(h.Sum(nil))
	u := url.URL{}
	u.Scheme = "https"
	u.Host = "www.eeo.cn"
	u.Path = "/webcast_partner.html"
	values := u.Query()
	values.Add("courseKey", m.Get("courseKey"))
	values.Add("lessonid", m.Get("lessonid"))
	values.Add("account", req.Account)
	values.Add("nickname", req.Nickname)
	values.Add("checkCode", code)
	u.RawQuery = values.Encode()
	return u.String(), nil
}
