package classin

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/go-resty/resty/v2"
	"sort"
	"strconv"
	"strings"
	"time"
)

// ApiBaseURLV1 V1 Api网址
const ApiBaseURLV1 = "https://api.eeo.cn/partner/api"

// ApiBaseURLV2 V2 Api网址
const ApiBaseURLV2 = "https://api.eeo.cn"

// Client ClassIn接口请求客户端
type Client struct {
	// 密钥
	Secret string

	// 机构认证 ID
	SID string
}

func NewClient(sID string, secret string) *Client {
	return &Client{
		secret,
		sID,
	}
}

// httpPostV1 ClassIn v1请求接口post通用方法
func (c *Client) httpPostV1(uri string, formData map[string]string, resp ICommonResp) *Error {
	httpClient := resty.New()
	httpClient.BaseURL = ApiBaseURLV1
	request := httpClient.R()
	request.SetResult(&resp)
	formData["SID"] = c.SID
	timestamp, safeKey := c.genSignV1()
	formData["safeKey"] = safeKey
	formData["timeStamp"] = timestamp
	request.SetFormData(formData)
	_, err := request.Post(uri)
	if err != nil {
		return NewServerError(err.Error())
	}
	if resp.IsOK() {
		return nil
	}
	return resp.GetError()
}

// httpPostV2 ClassIn v2请求接口post通用方法
func (c *Client) httpPostV2(uri string, data interface{}, signMap map[string]string, resp ICommonResp) *Error {
	httpClient := resty.New()
	httpClient.BaseURL = ApiBaseURLV2
	request := httpClient.R()
	request.SetHeader("Accept", "application/json")
	timestamp, sign := c.genSignV2(signMap)
	request.SetHeader("X-EEO-SIGN", sign)
	request.SetHeader("X-EEO-UID", c.SID)
	request.SetHeader("X-EEO-TS", timestamp)
	request.SetResult(&resp)
	request.SetBody(data)
	_, err := request.Post(uri)
	if err != nil {
		return NewServerError(err.Error())
	}
	if resp.IsOK() {
		return nil
	}
	return resp.GetError()
}

// genSignV1 V1生成签名
// 返回：
// 第一个参数为当前参与签名的时间戳
// 第二个参数为签名
func (c *Client) genSignV1() (string, string) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	str := c.Secret + timestamp
	return timestamp, c.md5(str)
}

// genSignV2 V2生成签名
// 返回：
// 第一个参数为当前参与签名的时间戳
// 第二个参数为签名
func (c *Client) genSignV2(data map[string]string) (string, string) {
	data["timeStamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	data["sid"] = c.SID
	keys := make([]string, 0, len(data))
	for k, _ := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	strs := make([]string, 0, len(data))
	for _, key := range keys {
		strs = append(strs, key+"="+data[key])
	}
	str := strings.Join(strs, "&") + "&key=" + c.Secret
	return data["timeStamp"], c.md5(str)
}

// md5 生成32位md5字符串
func (c *Client) md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
