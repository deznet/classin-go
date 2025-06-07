package classin

import (
	"os"
	"strconv"
	"testing"
)

// 获取唤醒客户端并进入教室链接测试
func TestClient_GetLoginLinked(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	req := new(GetLoginLinkedReq)
	req.LifeTime = 86400
	req.Uid, _ = strconv.ParseInt(os.Getenv("CLASSIN_USERID"), 10, 64)
	req.CourseId, _ = strconv.ParseInt(os.Getenv("CLASSIN_COURSEID"), 10, 64)
	req.ClassId, _ = strconv.ParseInt(os.Getenv("CLASSIN_CLASSID"), 10, 64)
	req.DeviceType = 1
	link, err := client.GetLoginLinked(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(link)
}

// 获取课程直播/回放播放器地址测试
func TestClient_GetWebcastUrl(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	req := new(GetWebcastUrlReq)
	req.CourseId, _ = strconv.ParseInt(os.Getenv("CLASSIN_COURSEID"), 10, 64)
	req.ClassId, _ = strconv.ParseInt(os.Getenv("CLASSIN_CLASSID"), 10, 64)
	link, err := client.GetWebcastUrl(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(link)
}

// 直播聊天室免二次登录网址测试
func TestClient_GetNoLoginWebcastUrl(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	req := new(GetNoLoginWebcastUrlReq)
	req.CourseId, _ = strconv.ParseInt(os.Getenv("CLASSIN_COURSEID"), 10, 64)
	req.ClassId, _ = strconv.ParseInt(os.Getenv("CLASSIN_CLASSID"), 10, 64)
	req.Account = os.Getenv("CLASSIN_TESTMOBILE")
	req.Nickname = os.Getenv("CLASSIN_NICKNAME")
	link, err := client.GetNoLoginWebcastUrl(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(link)
}
