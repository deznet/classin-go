package classin

import "strconv"

// ClassCommonReq 课堂添加修改公共请求参数
type ClassCommonReq struct {
	//班级（课程）ID
	CourseId int64 `json:"courseId"`

	//单元ID
	UnitId int64 `json:"unitId,omitempty"`

	//课堂活动名称,长度不超过50字
	Name string `json:"name"`

	//主讲教师UID
	TeacherUid int64 `json:"teacherUid"`

	//活动开始时间 可选择未来2年内的时间，Unix Epoch 时间戳（秒单位）
	StartTime int64 `json:"startTime"`

	//活动结束时间,Unix Epoch 时间戳（秒单位）
	EndTime int64 `json:"endTime"`

	//联席教师UID列表
	AssistantUids []int64 `json:"assistantUids"`

	//是否隐藏坐席区,
	//0 = 否（显示坐席区）
	//1 = 是（隐藏坐席区）。
	//默认为0，不传使用默认值，传错报参数错误；
	//当 cameraHide = 1 时，isAutoOnstage 会始终被设置为 0（也就是说，cameraHide = 1时，API 会忽略传参 isAutoOnstage 的值）
	CameraHide uint8 `json:"cameraHide"`

	//学生进入教室是否自动上台
	// 0 = 不自动，1 = 自动。
	// 默认为1，不传使用默认值，传错报参数错误；
	IsAutoOnstage uint8 `json:"isAutoOnstage"`

	//上台人数
	//包含老师，例：传1时，表示1v0，台上只显示老师头像。
	//默认为7（1v6），不传使用默认值，如机构最大台上人数配置小于1v6，会默认为机构最大台上人数。
	//传递的台上人数小于配置的最大台上人数，正常创建；传递的台上人数大于配置的最大台上人数，则会被重置为最大台上人数
	SeatNum uint8 `json:"seatNum"`

	//是否高清
	//0 = 非高清，1 = 高清，2 = 全高清。
	//默认为0，不传使用默认值，传错报参数错误。
	//目前仅支持 1V1 或 1V6 高清、全高清
	IsHd uint8 `json:"isHd"`

	//双摄模式，是否开启副摄像头
	//0 = 不开启，3 = 开启全高清副摄像头。
	//默认为0，不传使用默认值，传错报 参数错误
	//如果 isDc 等于3，课节的台上人数不是 1v1（即 seatNum 不等于2），则返回 该设置不支持双摄
	//如果 isDc 等于3，且 seatNum 等于2的话，则 isHd 一定会被设置为2（即这种情况 API 会忽略 isHd 的传参值）
	IsDc uint8 `json:"isDc"`

	//录课类型
	//0 = 录制教室，
	//1 = 录制现场，
	//2 = 两个都录。
	//默认为0，不传使用默认值，传错报参数错误
	RecordType uint8 `json:"recordType"`

	//是否开启录课
	//0 = 不录课（关闭），1 = 录课（开启）。
	//默认为0，不传使用默认值，传错报参数错误。
	//打开录课，ClassIn会将教室互动直播场景录制下来，可用于网页直播或者网页回放。
	//若需要网页直播或者网页回放，则必须选择录课，否则无法开启网页直播、网页回放
	RecordState uint8 `json:"recordState"`

	//是否开启直播
	//0 = 不直播（关闭），1 = 直播（开启）。
	//默认为0，不传使用默认值，传错报 参数错误。
	//若需要网页直播，则必须开启录课
	LiveState uint8 `json:"liveState"`

	//是否公开回放
	//0 = 不公开（关闭），1 = 公开（开启）。
	//默认为0，不传使用默认值，传错报 参数错误。
	//若需要网页回放，则必须开启录课
	OpenState uint8 `json:"openState"`

	//是否允许互相查看学习报告和评分
	//0 = 不允许，1 = 允许。
	//默认为1，不传使用默认值，传错报参数错误
	IsAllowCheck uint8 `json:"isAllowCheck"`

	//是否开启OMO站播 0-关闭，1-开启
	OmoStationBroadcast uint8 `json:"omoStationBroadcast"`
}

// CreateClassReq 创建课堂请求参数
type CreateClassReq struct {
	ClassCommonReq

	//唯一标识
	//例如：45s8d5a6asaa1ssf。
	//1-32 位字符，不符合规则的值报参数错误；
	//可传唯一标识，用于校验机构下是否已存在此唯一标识，防止因网络原因导致的重复创建
	UniqueIdentity string `json:"uniqueIdentity"`
}

// CreateClassResp 创建课堂返回参数
type CreateClassResp struct {
	CommonRespV2

	Data struct {
		//活动ID
		ActivityId int64 `json:"activityId"`

		//课堂ID
		ClassId int64 `json:"classId"`

		//课堂名称
		Name string `json:"name"`

		LiveInfo struct {
			RTMP string `json:"RTMP"` //RTMP协议的拉流地址
			HLS  string `json:"HLS"`  //HLS协议的拉流地址
			FLV  string `json:"FLV"`  //FLV协议的拉流地址
		} `json:"live_info"`

		//课节直播回放页面
		LiveUrl string `json:"live_url"`
	} `json:"data"`
}

// UpdateClassReq 编辑课堂请求参数
type UpdateClassReq struct {
	ClassCommonReq

	//课堂活动ID
	ActivityId int64 `json:"activityId"`
}

// CreateActivityNoClassReq 创建非课堂活动草稿请求参数
type CreateActivityNoClassReq struct {
	//班级（课程）ID
	CourseId int64 `json:"courseId"`

	//单元ID
	UnitId int64 `json:"unitId"`

	//活动类型
	//2=作业；3=测验；4=录播课；5=学习资料；6=讨论；7=答题卡；8=打卡
	ActivityType uint8 `json:"activityType"`

	//活动名称
	//长度不超过50字。注：课程下不支持创建同名单元
	Name string `json:"name"`

	//活动教师ID
	//老师必须在机构下，另外，如果老师不在班级下，会自动进班
	TeacherUid int64 `json:"teacherUid"`

	//活动开始时间
	//可选择未来3年内的时间，Unix Epoch 时间戳（秒单位）
	StartTime int64 `json:"startTime"`

	//活动结束时间
	//Unix Epoch 时间戳（秒单位）
	EndTime int64 `json:"endTime"`
}

// CreateActivityNoClassResp 创建非课堂活动草稿返回
type CreateActivityNoClassResp struct {
	CommonRespV2

	Data struct {
		//活动ID
		ActivityId int64 `json:"activityId"`

		//活动名
		Name string `json:"name"`
	} `json:"data"`
}

// ReleaseActivityReq 发布活动请求参数
type ReleaseActivityReq struct {
	// 班级（课程）ID
	CourseId int64 `json:"courseId"`

	//活动ID
	ActivityId int64 `json:"activityId"`
}

// DeleteActivityReq 删除活动请求参数
type DeleteActivityReq struct {
	// 班级（课程）ID
	CourseId int64 `json:"courseId"`

	//活动ID
	ActivityId int64 `json:"activityId"`
}

// MoveActivityReq 移动活动请求参数
type MoveActivityReq struct {
	// 班级（课程）ID
	CourseId int64 `json:"courseId"`

	//源单元ID
	UnitId int64 `json:"unitId"`

	//目标单元ID
	ToUnitId int64 `json:"toUnitId"`
}

// AddActivityStudentsReq 添加活动成员请求参数
type AddActivityStudentsReq struct {
	// 班级（课程）ID
	CourseId int64 `json:"courseId"`

	//活动创建完成后返回的活动ID
	ActivityId int64 `json:"activityId"`

	//学生UID数组
	StudentUids []int64 `json:"studentUids"`
}

// AddActivityStudentsResp 添加活动成员返回参数
type AddActivityStudentsResp struct {
	CommonRespV2

	Data []struct {
		StudentUid int64  `json:"studentUid"`
		Code       int    `json:"code"`
		Msg        string `json:"msg"`
	} `json:"data"`
}

// DeleteActivityStudentsReq 删除活动成员请求参数
type DeleteActivityStudentsReq struct {
	// 班级（课程）ID
	CourseId int64 `json:"courseId"`

	//活动创建完成后返回的活动ID
	ActivityId int64 `json:"activityId"`

	//学生UID数组
	StudentUids []int64 `json:"studentUids"`
}

// CreateClass 创建课堂
// @see https://docs.eeo.cn/api/zh-hans/LMS/createClassroom.html
// @param req *CreateClassReq
// @result *CreateClassResp
// @result *Error
func (c *Client) CreateClass(req *CreateClassReq) (*CreateClassResp, *Error) {
	uri := "/lms/activity/createClass"
	signData := make(map[string]string, 0)
	signData["courseId"] = strconv.FormatInt(req.CourseId, 10)
	if req.UnitId > 0 {
		signData["unitId"] = strconv.FormatInt(req.UnitId, 10)
	}
	signData["name"] = req.Name
	signData["teacherUid"] = strconv.FormatInt(req.TeacherUid, 10)
	signData["startTime"] = strconv.FormatInt(req.StartTime, 10)
	signData["endTime"] = strconv.FormatInt(req.EndTime, 10)
	signData["cameraHide"] = strconv.FormatUint(uint64(req.CameraHide), 10)
	signData["isAutoOnstage"] = strconv.FormatUint(uint64(req.IsAutoOnstage), 10)
	signData["seatNum"] = strconv.FormatUint(uint64(req.SeatNum), 10)
	signData["isHd"] = strconv.FormatUint(uint64(req.IsHd), 10)
	signData["isDc"] = strconv.FormatUint(uint64(req.IsDc), 10)
	signData["recordType"] = strconv.FormatUint(uint64(req.RecordType), 10)
	signData["recordState"] = strconv.FormatUint(uint64(req.RecordState), 10)
	signData["liveState"] = strconv.FormatUint(uint64(req.LiveState), 10)
	signData["openState"] = strconv.FormatUint(uint64(req.OpenState), 10)
	signData["isAllowCheck"] = strconv.FormatUint(uint64(req.IsAllowCheck), 10)
	signData["uniqueIdentity"] = req.UniqueIdentity
	signData["omoStationBroadcast"] = strconv.FormatUint(uint64(req.OmoStationBroadcast), 10)
	var result CreateClassResp
	err := c.httpPostV2(uri, req, signData, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateClass 编辑课堂
// @see https://docs.eeo.cn/api/zh-hans/LMS/updateClassroom.html
// @param req *UpdateClassReq
// @result *Error
func (c *Client) UpdateClass(req *UpdateClassReq) *Error {
	uri := "/lms/activity/updateClass"
	signData := make(map[string]string, 0)
	signData["courseId"] = strconv.FormatInt(req.CourseId, 10)
	if req.UnitId > 0 {
		signData["unitId"] = strconv.FormatInt(req.UnitId, 10)
	}
	signData["name"] = req.Name
	signData["teacherUid"] = strconv.FormatInt(req.TeacherUid, 10)
	signData["startTime"] = strconv.FormatInt(req.StartTime, 10)
	signData["endTime"] = strconv.FormatInt(req.EndTime, 10)
	signData["cameraHide"] = strconv.FormatUint(uint64(req.CameraHide), 10)
	signData["isAutoOnstage"] = strconv.FormatUint(uint64(req.IsAutoOnstage), 10)
	signData["seatNum"] = strconv.FormatUint(uint64(req.SeatNum), 10)
	signData["isHd"] = strconv.FormatUint(uint64(req.IsHd), 10)
	signData["isDc"] = strconv.FormatUint(uint64(req.IsDc), 10)
	signData["recordType"] = strconv.FormatUint(uint64(req.RecordType), 10)
	signData["recordState"] = strconv.FormatUint(uint64(req.RecordState), 10)
	signData["liveState"] = strconv.FormatUint(uint64(req.LiveState), 10)
	signData["openState"] = strconv.FormatUint(uint64(req.OpenState), 10)
	signData["isAllowCheck"] = strconv.FormatUint(uint64(req.IsAllowCheck), 10)
	signData["activityId"] = strconv.FormatInt(req.ActivityId, 10)
	signData["omoStationBroadcast"] = strconv.FormatUint(uint64(req.OmoStationBroadcast), 10)
	var result CommonRespV2
	err := c.httpPostV2(uri, req, signData, &result)
	if err != nil {
		return err
	}
	return nil
}

// CreateActivityNoClass 创建非课堂活动草稿
// @see https://docs.eeo.cn/api/zh-hans/LMS/createActivityNoClass.html
// @param req *CreateActivityNoClassReq
// @result int64 活动ID
// @result *Error
func (c *Client) CreateActivityNoClass(req *CreateActivityNoClassReq) (int64, *Error) {
	uri := "/lms/activity/createActivityNoClass"
	signData := make(map[string]string, 0)
	signData["courseId"] = strconv.FormatInt(req.CourseId, 10)
	signData["unitId"] = strconv.FormatInt(req.UnitId, 10)
	signData["activityType"] = strconv.FormatUint(uint64(req.ActivityType), 10)
	signData["name"] = req.Name
	signData["teacherUid"] = strconv.FormatInt(req.TeacherUid, 10)
	signData["startTime"] = strconv.FormatInt(req.StartTime, 10)
	signData["endTime"] = strconv.FormatInt(req.EndTime, 10)
	var result CreateActivityNoClassResp
	err := c.httpPostV2(uri, req, signData, &result)
	if err != nil {
		return 0, err
	}
	return result.Data.ActivityId, nil
}

// ReleaseActivity 发布活动
// @see https://docs.eeo.cn/api/zh-hans/LMS/releaseActivity.html
// @param req *ReleaseActivityReq
// @result *Error
func (c *Client) ReleaseActivity(req *ReleaseActivityReq) *Error {
	uri := "/lms/activity/release"
	signData := make(map[string]string, 0)
	signData["courseId"] = strconv.FormatInt(req.CourseId, 10)
	signData["activityId"] = strconv.FormatInt(req.ActivityId, 10)
	var result CommonRespV2
	err := c.httpPostV2(uri, req, signData, &result)
	//活动已发布直接返回成功
	if err.GetCode() == 29184 {
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}

// DeleteActivity 删除活动
// @see https://docs.eeo.cn/api/zh-hans/LMS/deleteActivity.html
// @param req *DeleteActivityReq
// @result *Error
func (c *Client) DeleteActivity(req *DeleteActivityReq) *Error {
	uri := "/lms/activity/delete"
	signData := make(map[string]string, 0)
	signData["courseId"] = strconv.FormatInt(req.CourseId, 10)
	signData["activityId"] = strconv.FormatInt(req.ActivityId, 10)
	var result CommonRespV2
	err := c.httpPostV2(uri, req, signData, &result)
	if err != nil {
		return err
	}
	return nil
}

// MoveActivity 移动活动
// @see https://docs.eeo.cn/api/zh-hans/LMS/moveActivity.html
// @param req *MoveActivityReq
// @result *Error
func (c *Client) MoveActivity(req *MoveActivityReq) *Error {
	uri := "/lms/unit/move"
	signData := make(map[string]string, 0)
	signData["courseId"] = strconv.FormatInt(req.CourseId, 10)
	signData["unitId"] = strconv.FormatInt(req.UnitId, 10)
	signData["toUnitId"] = strconv.FormatInt(req.ToUnitId, 10)
	var result CommonRespV2
	err := c.httpPostV2(uri, req, signData, &result)
	if err != nil {
		return err
	}
	return nil
}

// AddActivityStudents 添加活动成员
// @see https://docs.eeo.cn/api/zh-hans/LMS/addStudent.html
// @param req *AddActivityStudentsReq
// @result *AddActivityStudentsResp
// @result *Error
func (c *Client) AddActivityStudents(req *AddActivityStudentsReq) (*AddActivityStudentsResp, *Error) {
	uri := "/lms/activity/addStudent"
	signData := make(map[string]string, 0)
	signData["courseId"] = strconv.FormatInt(req.CourseId, 10)
	signData["activityId"] = strconv.FormatInt(req.ActivityId, 10)
	var result AddActivityStudentsResp
	err := c.httpPostV2(uri, req, signData, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteActivityStudents 删除活动成员
// @see https://docs.eeo.cn/api/zh-hans/LMS/deleteStudent.html
// @param req *DeleteActivityStudentsReq
// @result *Error
func (c *Client) DeleteActivityStudents(req *DeleteActivityStudentsReq) *Error {
	uri := "/lms/activity/deleteStudent"
	signData := make(map[string]string, 0)
	signData["courseId"] = strconv.FormatInt(req.CourseId, 10)
	signData["activityId"] = strconv.FormatInt(req.ActivityId, 10)
	var result CommonRespV2
	err := c.httpPostV2(uri, req, signData, &result)
	if err != nil {
		return err
	}
	return nil
}
