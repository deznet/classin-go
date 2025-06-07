package classin

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strconv"
	"testing"
	"time"
)

// 创建课堂测试
func TestClient_CreateClass(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	req := new(CreateClassReq)
	req.CourseId, _ = strconv.ParseInt(os.Getenv("CLASSIN_COURSEID"), 10, 64)
	req.UnitId, _ = strconv.ParseInt(os.Getenv("CLASSIN_UNITID"), 10, 64)
	req.Name = "测试活动1"
	req.TeacherUid, _ = strconv.ParseInt(os.Getenv("CLASSIN_USERID"), 10, 64)
	req.StartTime = time.Now().Unix() + 3600
	req.EndTime = time.Now().Unix() + 7200
	req.AssistantUids = []int64{}
	req.CameraHide = 0
	req.IsAutoOnstage = 1
	req.SeatNum = 10
	req.IsHd = 0
	req.IsDc = 0
	req.RecordType = 0
	req.RecordState = 1
	req.LiveState = 1
	req.OpenState = 1
	req.IsAllowCheck = 0
	req.UniqueIdentity = ""
	resp, err := client.CreateClass(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

// 创建非课堂活动草稿测试
func TestClient_CreateActivityNoClass(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	req := new(CreateActivityNoClassReq)
	req.CourseId, _ = strconv.ParseInt(os.Getenv("CLASSIN_COURSEID"), 10, 64)
	req.UnitId, _ = strconv.ParseInt(os.Getenv("CLASSIN_UNITID"), 10, 64)
	req.ActivityType = 2
	req.Name = "测试1"
	req.TeacherUid, _ = strconv.ParseInt(os.Getenv("CLASSIN_USERID"), 10, 64)
	req.StartTime = time.Now().Unix() + 3600
	req.EndTime = time.Now().Unix() + 7200
	id, err := client.CreateActivityNoClass(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}

// 修改课堂测试
func TestClient_UpdateClass(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	req := new(UpdateClassReq)
	req.CourseId, _ = strconv.ParseInt(os.Getenv("CLASSIN_COURSEID"), 10, 64)
	req.ActivityId, _ = strconv.ParseInt(os.Getenv("CLASSIN_ACTIVITYID"), 10, 64)
	req.UnitId, _ = strconv.ParseInt(os.Getenv("CLASSIN_UNITID"), 10, 64)
	req.Name = "测试活动2"
	req.TeacherUid, _ = strconv.ParseInt(os.Getenv("CLASSIN_USERID"), 10, 64)
	req.StartTime = time.Now().Unix() + 7200
	req.EndTime = time.Now().Unix() + 9200
	req.AssistantUids = []int64{}
	req.CameraHide = 0
	req.IsAutoOnstage = 1
	req.SeatNum = 10
	req.IsHd = 0
	req.IsDc = 0
	req.RecordType = 0
	req.RecordState = 1
	req.LiveState = 1
	req.OpenState = 1
	req.IsAllowCheck = 0
	err := client.UpdateClass(req)
	if err != nil {
		t.Fatal(err)
	}
}

// 发布活动测试
func TestClient_ReleaseActivity(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	req := new(ReleaseActivityReq)
	req.ActivityId, _ = strconv.ParseInt(os.Getenv("CLASSIN_ACTIVITYID"), 10, 64)
	req.CourseId, _ = strconv.ParseInt(os.Getenv("CLASSIN_COURSEID"), 10, 64)
	err := client.ReleaseActivity(req)
	if err != nil {
		t.Fatal(err)
	}
}

// 添加活动成员
func TestClient_AddActivityStudents(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	req := new(AddActivityStudentsReq)
	req.ActivityId, _ = strconv.ParseInt(os.Getenv("CLASSIN_ACTIVITYID"), 10, 64)
	req.CourseId, _ = strconv.ParseInt(os.Getenv("CLASSIN_COURSEID"), 10, 64)
	req.StudentUids = []int64{28759840}
	resp, err := client.AddActivityStudents(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

// 删除活动成员测试
func TestClient_DeleteActivityStudents(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	req := new(DeleteActivityStudentsReq)
	req.ActivityId, _ = strconv.ParseInt(os.Getenv("CLASSIN_ACTIVITYID"), 10, 64)
	req.CourseId, _ = strconv.ParseInt(os.Getenv("CLASSIN_COURSEID"), 10, 64)
	req.StudentUids = []int64{28759840}
	err := client.DeleteActivityStudents(req)
	if err != nil {
		t.Fatal(err)
	}
}

// 删除活动测试
func TestClient_DeleteActivity(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	req := new(DeleteActivityReq)
	req.ActivityId, _ = strconv.ParseInt(os.Getenv("CLASSIN_ACTIVITYID"), 10, 64)
	req.CourseId, _ = strconv.ParseInt(os.Getenv("CLASSIN_COURSEID"), 10, 64)
	err := client.DeleteActivity(req)
	if err != nil {
		t.Fatal(err)
	}
}
