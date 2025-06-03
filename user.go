package classin

import (
	"strconv"
)

// RegisterReq 注册用户请求参数
type RegisterReq struct {
	//注册手机号
	//格式为：00国家号-手机号；注意：中国大陆手机号不写国家号，手机号第一位不能为0
	//例如：美国手机号 1 (800) 643-7676 填成 001-8006437676；中国大陆手机号填成 15800000001
	//必填 与Email 2选1
	Telephone string

	//Email
	//必填 与Telephone 2选1
	Email string

	//昵称、姓名
	//最长24位字符，超过24字会自动截取为24字
	//填写该参数，则会将该参数作为老师或学生的昵称（显示在客户端）；
	//不填写昵称，手机号会作为老师或者学生的昵称（用户登录客户端，会弹窗让用户填写昵称）。
	//昵称显示在教室内用户摄像头下方,最大24个字符
	//非必填
	Nickname string

	//明文密码
	//6-20位，不符合会报错
	//必填 与 Md5Password 2选1
	Password string

	//MD5 加密密码
	//32位 MD5 加密
	//必填 与 Password 2选1
	Md5Password string

	//上传的用户头像，
	//头像会显示在 ClassIn 客户端；不上传头像，登录客户端时会弹出选择头像的弹窗
	//二进制流
	//非必填
	FileData string

	//是否加为机构成员（老师、学生）
	//addToSchoolMember等于1或者2时，用户会显示在 eeo.cn 后台的学生/教师管理页面，并且其 nickname 会被用来设置学生或者老师在 eeo.cn 后台显示的姓名
	//0 不加为机构成员；
	//1 加为机构学生；
	//2 加为机构老师；
	//其他值不加为机构成员。
	//不填默认为0。
	AddToSchoolMember uint8
}

// RegisterResp 注册用户响应参数
type RegisterResp struct {
	CommonRespV1

	//注册成功返回的用户 UID（此 UID 为该账号唯一 ID）
	Data int64 `json:"data"`
}

// AddSchoolStudentReq 添加学生请求参数
type AddSchoolStudentReq struct {
	// 学生账号
	// 手机号或者是邮箱，格式为：00国家号-手机号；注意：中国大陆手机号不写国家号
	// 例如：美国手机号 1 (800) 643-7676 填成 001-8006437676；中国大陆手机号填成 15800000001
	// 必填
	StudentAccount string

	// 学生名字
	// 会显示在 eeo.cn 管理后台的 学生管理 页面
	// 1-24字，不区分中英文，超过24会自动截取为24字
	StudentName string
}

// EditSchoolStudentReq 编辑学生请求参数
type EditSchoolStudentReq struct {
	// 用户 UID
	// 必填
	StudentUid int64

	// 学生的姓名
	// 1-24位字符，不区分中英文，超过24字会截取为24字
	// 必填
	StudentName string
}

// ModifyCourseStudentNickNameReq 同步学生班级昵称请求参数
type ModifyCourseStudentNickNameReq struct {
	StudentUids []int64 `json:"studentUids"`
}

// ModifyCourseStudentNickNameResp 同步学生班级昵称返回
type ModifyCourseStudentNickNameResp struct {
	CommonRespV2
	Data []struct {
		StudentUid int64  `json:"studentUid"`
		Code       int    `json:"code"`
		Msg        string `json:"msg"`
	} `json:"data"`
}

// AddTeacherReq 添加老师请求参数
type AddTeacherReq struct {
	// 老师的账号
	// 手机号或者是邮箱
	// 格式为：00国家号-手机号；注意：中国大陆手机号不写国家号
	// 例如：美国手机号 1 (800) 643-7676 填成 001-8006437676；中国大陆手机号填成 15800000001
	// 必填
	TeacherAccount string

	// 老师的姓名
	// 会显示在 eeo.cn 管理后台的 教师管理 页面
	// 1-24字，不区分中英文，超过24字会自动截取为24字
	// 必填
	TeacherName string

	// 老师的头像
	// 二进制流
	// 非必填
	FileData string
}

// AddTeacherResp 添加老师返回参数
type AddTeacherResp struct {
	CommonRespV1

	//机构和老师的关系的 ID，这个ID可以忽略
	Data int64 `json:"data"`
}

// EditTeacherReq 修改老师请求参数
type EditTeacherReq struct {
	// 老师的用户UID
	// 必填
	TeacherUid int64

	// 老师的姓名
	// 1-24位字符，不区分中英文，超过24字会截取为24字
	// 必填
	TeacherName string

	// 老师的头像
	// 非必填
	FileData string
}

// Register 注册用户
// @param req *RegisterReq 注册请求参数
// @result  int64 注册成功返回的用户 UID
// @result *Error 错误信息
func (c *Client) Register(req *RegisterReq) (int64, *Error) {
	formData := make(map[string]string)
	formData["telephone"] = req.Telephone
	formData["email"] = req.Email
	formData["nickname"] = req.Nickname
	formData["password"] = req.Password
	formData["md5pass"] = req.Md5Password
	formData["Filedata"] = req.FileData
	formData["addToSchoolMember"] = strconv.Itoa(int(req.AddToSchoolMember))
	uri := "/course.api.php?action=register"
	var result RegisterResp
	err := c.httpPostV1(uri, formData, &result)
	if err != nil {
		//如果手机号/邮箱已存在,直接返回数据
		if err.GetCode() == 135 || err.GetCode() == 461 {
			return result.Data, nil
		}
		return 0, err
	}
	return result.Data, nil
}

// AddSchoolStudent 添加学生（机构下添加学生）
// @param req *AddSchoolStudentReq 添加学生请求参数
// @result *Error 添加失败返回错误
func (c *Client) AddSchoolStudent(req *AddSchoolStudentReq) *Error {
	formData := make(map[string]string)
	formData["studentAccount"] = req.StudentAccount
	formData["studentName"] = req.StudentName
	uri := "/course.api.php?action=addSchoolStudent"
	var result CommonRespV1
	err := c.httpPostV1(uri, formData, &result)
	if err != nil {
		//如果已是机构学生
		if err.GetCode() == 133 {
			return nil
		}
		return err
	}
	return nil
}

// EditSchoolStudent 编辑学生信息
// @param req *EditSchoolStudentReq 编辑学生信息请求参数
// @result *Error 编辑失败返回错误信息
func (c *Client) EditSchoolStudent(req *EditSchoolStudentReq) *Error {
	formData := make(map[string]string)
	formData["studentUid"] = strconv.FormatInt(req.StudentUid, 10)
	formData["studentName"] = req.StudentName
	uri := "/course.api.php?action=editSchoolStudent"
	var result CommonRespV1
	err := c.httpPostV1(uri, formData, &result)
	if err != nil {
		return err
	}
	return nil
}

// ModifyCourseStudentNickName 同步学生班级昵称
// @param stuIds 要刷新的用户id集合，最多100个
// @result *ModifyCourseStudentNickNameResp
// @result *Error
func (c *Client) ModifyCourseStudentNickName(stuIds []int64) (*ModifyCourseStudentNickNameResp, *Error) {
	uri := "/schooluser/modifyCourseStudentNickName"
	signData := make(map[string]string, 0)
	var result ModifyCourseStudentNickNameResp
	data := ModifyCourseStudentNickNameReq{StudentUids: stuIds}
	err := c.httpPostV2(uri, data, signData, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// AddTeacher 添加教师
// @param req AddTeacherReq 添加教师参数
// @result *Error 添加失败返回错误
func (c *Client) AddTeacher(req *AddTeacherReq) (int64, *Error) {
	formData := make(map[string]string)
	formData["teacherAccount"] = req.TeacherAccount
	formData["teacherName"] = req.TeacherName
	formData["Filedata"] = req.FileData
	uri := "/course.api.php?action=addTeacher"
	var result AddTeacherResp
	err := c.httpPostV1(uri, formData, &result)
	if err != nil {
		//该帐户已经在机构下存在
		if err.GetCode() == 133 {
			return result.Data, nil
		}
		return 0, err
	}
	return result.Data, nil
}

// EditTeacher 修改老师
// @param req EditTeacherReq 修改教师参数
// @result *Error 修改失败返回错误
func (c *Client) EditTeacher(req *EditTeacherReq) *Error {
	formData := make(map[string]string)
	formData["teacherUid"] = strconv.FormatInt(req.TeacherUid, 10)
	formData["teacherName"] = req.TeacherName
	formData["Filedata"] = req.FileData
	uri := "/course.api.php?action=editTeacher"
	var result CommonRespV1
	err := c.httpPostV1(uri, formData, &result)
	if err != nil {
		return err
	}
	return nil
}

// StopUsingTeacher 停用老师
// 可通过此接口将老师停用，停用后机构不可给该老师排课，该老师在客户端看不到机构授权云盘资源等
// @param teacherUid int64 注册用户接口返回的用户 UID
// @result *Error 停用失败返回信息
func (c *Client) StopUsingTeacher(teacherUid int64) *Error {
	formData := make(map[string]string)
	formData["teacherUid"] = strconv.FormatInt(teacherUid, 10)
	uri := "/course.api.php?action=stopUsingTeacher"
	var result CommonRespV1
	err := c.httpPostV1(uri, formData, &result)
	if err != nil {
		return err
	}
	return nil
}

// RestartUsingTeacher 启用老师
// @param teacherUid int64 注册用户接口返回的用户 UID
// @result *Error 启用老师失败返回信息
func (c *Client) RestartUsingTeacher(teacherUid int64) *Error {
	formData := make(map[string]string)
	formData["teacherUid"] = strconv.FormatInt(teacherUid, 10)
	uri := "/course.api.php?action=restartUsingTeacher"
	var result CommonRespV1
	err := c.httpPostV1(uri, formData, &result)
	if err != nil {
		return err
	}
	return nil
}
