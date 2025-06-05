package classin

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strconv"
	"testing"
)

// 创建单元测试
func TestClient_CreateUnit(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	req := new(CreateUnitReq)
	req.CourseId, _ = strconv.ParseInt(os.Getenv("CLASSIN_COURSEID"), 10, 64)
	req.PublishFlag = 0
	req.Name = os.Getenv("CLASSIN_UNITNAME")
	req.Content = ""
	unitId, err := client.CreateUnit(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(unitId)
}

// 修改单元测试
func TestClient_UpdateUnit(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	req := new(UpdateUnitReq)
	req.CourseId, _ = strconv.ParseInt(os.Getenv("CLASSIN_COURSEID"), 10, 64)
	req.PublishFlag = 0
	req.Name = os.Getenv("CLASSIN_UNITNAME")
	req.Content = ""
	req.UnitId, _ = strconv.ParseInt(os.Getenv("CLASSIN_UNITID"), 10, 64)
	err := client.UpdateUnit(req)
	if err != nil {
		t.Fatal(err)
	}
}

// 删除单元测试
func TestClient_DeleteUnit(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	req := new(DeleteUnitReq)
	req.CourseId, _ = strconv.ParseInt(os.Getenv("CLASSIN_COURSEID"), 10, 64)
	req.UnitId, _ = strconv.ParseInt(os.Getenv("CLASSIN_UNITID"), 10, 64)
	err := client.DeleteUnit(req)
	if err != nil {
		t.Fatal(err)
	}
}
