package classin

import (
	"encoding/json"
	"strconv"
)

// AddCourseReq 创建课程请求参数
type AddCourseReq struct {
	//课程名称
	//1-40位字符，不区分中英文，超过会自动截取为40字
	//必填
	CourseName string

	//可用资源文件夹 ID
	//非必填
	FolderId string

	//上传的课程封面图片
	//非必填
	FileData string

	//过期时间
	//非必填
	//0 永不过期
	ExpiryTime int64

	//班主任 UID
	//非必填
	MainTeacherUid int64

	//课程学科分类
	//非必填
	//0:空; 1:语文; 2:数学; 3:英语; 4:物理; 5:化学; 6:生物; 7:政治; 8:历史; 9:地理; 10:思想品德; 11:音乐; 12:体育; 13:美术; 14:通用技术; 15:信息技术; 16:科学; 99:其他学科
	SubjectId int8

	//课程简介 非必填
	CourseIntroduce string

	//教室设置 ID
	//非必填
	//默认0
	ClassroomSettingId int64

	//唯一标识
	//非必填
	//机构可传唯一标识，传入此值后，我们会检验已创建课程中是否有该唯一标识
	CourseUniqueIdentity string

	//是否允许班级成员在群里互相添加好友，
	//0=不允许，1=允许
	//非必填
	AllowAddFriend int8

	//是否允许学生在群里修改其班级昵称
	//0=不允许 1=允许
	//非必填
	AllowStudentModifyNickname int8
}

// AddCourseResp 创建课程响应参数
type AddCourseResp struct {
	//课程 ID
	Data int64 `json:"data"`

	CommonRespV1
}

// EditCourseReq 编辑课程请求参数
type EditCourseReq struct {
	//课程 ID
	//必填
	CourseId int64

	//原班主任是否加入教师列表
	//1加入，2不加入，默认为1
	//非必填
	Stamp int8

	AddCourseReq
}

// AddCourseStudentReq 课程下添加学生/旁听（单个）请求参数
type AddCourseStudentReq struct {
	//课程 ID
	//必填
	CourseId int64

	//学生和旁听的识别
	//1 为学生,2 为旁听
	//必填
	Identity uint8

	//学生 UID
	//注册用户接口返回的用户 UID
	//必填
	StudentUid int64

	//机构后台旁听生的姓名
	//1-24字，不区分中英文，超过24会自动截取为24字
	//仅用于当identity为2（旁听身份）时，才使用此参数。当identity为2时，如果没有传此参的话，则使用手机号码作为旁听生的名字
	//非必填
	StudentName string
}

// DelCourseStudentReq 课程下删除学生/旁听（单个）请求参数
type DelCourseStudentReq struct {
	//课程 ID
	//必填
	CourseId int64

	//学生和旁听的识别
	//1 为学生,2 为旁听
	//必填
	Identity uint8

	//学生 UID
	//注册用户接口返回的用户 UID
	//必填
	StudentUid int64
}

// AddCourseStudentMultiReq 课程下添加学生/旁听（多个）请求参数
type AddCourseStudentMultiReq struct {
	//课程 ID
	//必填
	CourseId int64

	//学生和旁听的识别
	//1 为学生,2 为旁听
	//必填
	Identity uint8

	//需要添加的帐号数组
	//旁听人数最多可添加20人
	//必填
	StudentJson []*CourseStuJson
}

// CourseStuJson 学生数组参数
type CourseStuJson struct {
	//学生 UID
	//注册用户接口返回的用户 UID
	//必填
	UID int64 `json:"uid"`

	//机构后台旁听生的姓名
	//1-24字，不区分中英文，超过24字会自动截取为24字
	//仅用于当identity为2（旁听身份）时，才使用此参数。当identity为2时，如果没有传此参的话，则使用手机号码作为旁听生的名字
	//非必填
	Name string `json:"name,omitempty"`

	//用户自定义标识
	//1-50字，不区分中英文，超过50会自动截取为50字
	//不为空则原样返回，为空则不返回该字段
	//非必填
	CustomColumn string `json:"customColumn,omitempty"`
}

// AddCourseStudentMultiResp 课程下添加学生/旁听（多个）返回参数
type AddCourseStudentMultiResp struct {
	CommonRespV1
	Data []struct {
		CustomColumn string `json:"customColumn"`
		ErrNo        uint32 `json:"errno"`
		ErrorMsg     string `json:"error"`
	} `json:"data"`
}

// DelCourseStudentMultiReq 课程下删除学生/旁听（多个）请求参数
type DelCourseStudentMultiReq struct {
	//课程 ID
	//必填
	CourseId int64

	//学生和旁听的识别
	//1 为学生,2 为旁听
	//必填
	Identity uint8

	//需要删除学生UID数组
	//数组如果不为空，长度至少为1，可多个UID
	//必填
	StudentUidJson []int64
}

// DelCourseStudentMultiResp 课程下删除学生/旁听（多个）返回参数
type DelCourseStudentMultiResp struct {
	CommonRespV1
	Data []struct {
		ErrNo    uint32 `json:"errno"`
		ErrorMsg string `json:"error"`
	} `json:"data"`
}

// AddCourseTeacherReq 添加课程教师请求参数
type AddCourseTeacherReq struct {
	//班级（课程）ID
	CourseId int64 `json:"courseId"`

	//教师UID数组
	TeacherUids []int64 `json:"teacherUids"`
}

// AddCourseTeacherResp  添加课程教师返回参数
type AddCourseTeacherResp struct {
	CommonRespV2
	Data []struct {
		TeacherUid int64  `json:"teacherUid"`
		Code       int64  `json:"code"`
		Msg        string `json:"msg"`
	} `json:"data"`
}

// RemoveCourseTeacherReq 移除课程老师请求参数
type RemoveCourseTeacherReq struct {
	// 课程 ID
	CourseId int64

	//老师 UID
	TeacherUid int64
}

// AddCourse 创建课程
// @see https://docs.eeo.cn/api/zh-hans/classroom/addCourse.html
// @param req *AddCourseReq 创建课程请求参数
// @result int64 创建成功返回的课程 ID
// @result *Error 创建失败返回的错误
func (c *Client) AddCourse(req *AddCourseReq) (int64, *Error) {
	formData := make(map[string]string)
	formData["courseName"] = req.CourseName
	formData["folderId"] = req.FolderId
	formData["Filedata"] = req.FileData
	formData["expiryTime"] = strconv.FormatInt(req.ExpiryTime, 10)
	if req.MainTeacherUid > 0 {
		formData["mainTeacherUid"] = strconv.FormatInt(req.MainTeacherUid, 10)
	}
	formData["subjectId"] = strconv.Itoa(int(req.SubjectId))
	formData["courseIntroduce"] = req.CourseIntroduce
	formData["classroomSettingId"] = strconv.FormatInt(req.ClassroomSettingId, 10)
	formData["courseUniqueIdentity"] = req.CourseUniqueIdentity
	formData["allowAddFriend"] = strconv.Itoa(int(req.AllowAddFriend))
	formData["allowStudentModifyNickname"] = strconv.Itoa(int(req.AllowStudentModifyNickname))
	uri := "/course.api.php?action=addCourse"
	var result AddCourseResp
	err := c.httpPostV1(uri, formData, &result)
	if err != nil {
		return 0, err
	}
	return result.Data, nil
}

// EditCourse 编辑课程
// @see https://docs.eeo.cn/api/zh-hans/classroom/editCourse.html
// @param req *EditCourseReq
// @result *Error 编辑失败返回错误信息
func (c *Client) EditCourse(req *EditCourseReq) *Error {
	formData := make(map[string]string)
	formData["courseId"] = strconv.FormatInt(req.CourseId, 10)
	formData["folderId"] = req.FolderId
	formData["courseName"] = req.CourseName
	formData["Filedata"] = req.FileData
	formData["expiryTime"] = strconv.FormatInt(req.ExpiryTime, 10)
	if req.MainTeacherUid > 0 {
		formData["mainTeacherUid"] = strconv.FormatInt(req.MainTeacherUid, 10)
	}
	formData["subjectId"] = strconv.Itoa(int(req.SubjectId))
	formData["courseIntroduce"] = req.CourseIntroduce
	formData["classroomSettingId"] = strconv.FormatInt(req.ClassroomSettingId, 10)
	formData["allowAddFriend"] = strconv.Itoa(int(req.AllowAddFriend))
	formData["allowStudentModifyNickname"] = strconv.Itoa(int(req.AllowStudentModifyNickname))
	formData["stamp"] = strconv.Itoa(int(req.Stamp))
	uri := "/course.api.php?action=editCourse"
	var result CommonRespV1
	err := c.httpPostV1(uri, formData, &result)
	if err != nil {
		return err
	}
	return nil
}

// EndCourse 结束课程
// @see https://docs.eeo.cn/api/zh-hans/classroom/endCourse.html
// @param courseId int64
// @result *Error 失败返回错误信息
func (c *Client) EndCourse(courseId int64) *Error {
	formData := make(map[string]string)
	formData["courseId"] = strconv.FormatInt(courseId, 10)
	uri := "/course.api.php?action=endCourse"
	var result CommonRespV1
	err := c.httpPostV1(uri, formData, &result)
	if err != nil {
		return err
	}
	return nil
}

// AddCourseStudent 课程下添加学生/旁听（单个）
// @see https://docs.eeo.cn/api/zh-hans/classroom/addCourseStudent.html
// @param req *AddCourseStudentReq
// @result *Error 失败返回错误信息
func (c *Client) AddCourseStudent(req *AddCourseStudentReq) *Error {
	formData := make(map[string]string)
	formData["courseId"] = strconv.FormatInt(req.CourseId, 10)
	formData["identity"] = strconv.Itoa(int(req.Identity))
	formData["studentUid"] = strconv.FormatInt(req.StudentUid, 10)
	formData["studentName"] = req.StudentName
	uri := "/course.api.php?action=addCourseStudent"
	var result CommonRespV1
	err := c.httpPostV1(uri, formData, &result)
	if err != nil {
		return err
	}
	return nil
}

// DelCourseStudent 课程下删除学生/旁听（单个）
// @see https://docs.eeo.cn/api/zh-hans/classroom/delCourseStudent.html
// @param req *DelCourseStudentReq
// @result *Error 失败返回错误信息
func (c *Client) DelCourseStudent(req *DelCourseStudentReq) *Error {
	formData := make(map[string]string)
	formData["courseId"] = strconv.FormatInt(req.CourseId, 10)
	formData["identity"] = strconv.Itoa(int(req.Identity))
	formData["studentUid"] = strconv.FormatInt(req.StudentUid, 10)
	uri := "/course.api.php?action=delCourseStudent"
	var result CommonRespV1
	err := c.httpPostV1(uri, formData, &result)
	if err != nil {
		return err
	}
	return nil
}

// AddCourseStudentMulti 课程下添加学生/旁听（多个）
// @see https://docs.eeo.cn/api/zh-hans/classroom/addCourseStudentMultiple.html
// @param req *AddCourseStudentMultiReq
// @result *Error 失败返回错误信息
func (c *Client) AddCourseStudentMulti(req *AddCourseStudentMultiReq) (*AddCourseStudentMultiResp, *Error) {
	formData := make(map[string]string)
	formData["courseId"] = strconv.FormatInt(req.CourseId, 10)
	formData["identity"] = strconv.Itoa(int(req.Identity))
	studentJson, err := json.Marshal(req.StudentJson)
	if err != nil {
		return nil, NewServerError(err.Error())
	}
	formData["studentJson"] = string(studentJson)
	uri := "/course.api.php?action=addCourseStudentMultiple"
	var result AddCourseStudentMultiResp
	err1 := c.httpPostV1(uri, formData, &result)
	if err1 != nil {
		return nil, err1
	}
	return &result, nil
}

// DelCourseStudentMulti 课程下删除学生/旁听（多个）
// @see https://docs.eeo.cn/api/zh-hans/classroom/delCourseStudentMultiple.html
// @param req *DelCourseStudentMultiReq
// @result *DelCourseStudentMultiResp
// @result *Error 失败返回错误信息
func (c *Client) DelCourseStudentMulti(req *DelCourseStudentMultiReq) (*DelCourseStudentMultiResp, *Error) {
	formData := make(map[string]string)
	formData["courseId"] = strconv.FormatInt(req.CourseId, 10)
	formData["identity"] = strconv.Itoa(int(req.Identity))
	studentUidJson, err := json.Marshal(req.StudentUidJson)
	if err != nil {
		return nil, NewServerError(err.Error())
	}
	formData["studentUidJson"] = string(studentUidJson)
	uri := "/course.api.php?action=delCourseStudentMultiple"
	var result DelCourseStudentMultiResp
	err1 := c.httpPostV1(uri, formData, &result)
	if err1 != nil {
		return nil, err1
	}
	return &result, nil
}

// ModifyGroupMemberNickname 修改群成员的班级昵称
// @see https://docs.eeo.cn/api/zh-hans/group/modifyGroupMemberNickname.html
// @param courseId int64 课程 ID
// @result *Error 失败返回错误信息
func (c *Client) ModifyGroupMemberNickname(courseId int64) *Error {
	formData := make(map[string]string)
	formData["courseId"] = strconv.FormatInt(courseId, 10)
	uri := "/course.api.php?action=modifyGroupMemberNickname"
	var result CommonRespV1
	err := c.httpPostV1(uri, formData, &result)
	if err != nil {
		return err
	}
	return nil
}

// AddCourseTeacher 添加课程教师
// @see https://docs.eeo.cn/api/zh-hans/classroom/addCourseTeacher.html
// @param req *AddCourseTeacherReq
// @result *AddCourseTeacherResp
// @result *Error
func (c *Client) AddCourseTeacher(req *AddCourseTeacherReq) (*AddCourseTeacherResp, *Error) {
	uri := "/course/addCourseTeacher"
	signData := make(map[string]string, 0)
	signData["courseId"] = strconv.FormatInt(req.CourseId, 10)
	var result AddCourseTeacherResp
	err := c.httpPostV2(uri, req, signData, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// RemoveCourseTeacher 移除课程老师
// @see https://docs.eeo.cn/api/zh-hans/classroom/removeCourseTeacher.html
// @param req *RemoveCourseTeacherReq
// @result *Error
func (c *Client) RemoveCourseTeacher(req *RemoveCourseTeacherReq) *Error {
	formData := make(map[string]string)
	formData["courseId"] = strconv.FormatInt(req.CourseId, 10)
	formData["teacherUid"] = strconv.FormatInt(req.TeacherUid, 10)
	uri := "/course.api.php?action=removeCourseTeacher"
	var result CommonRespV1
	err := c.httpPostV1(uri, formData, &result)
	if err != nil {
		return err
	}
	return nil
}
