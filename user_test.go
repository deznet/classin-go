package classin

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strconv"
	"testing"
)

// 用户注册测试
func TestClient_Register(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	req := RegisterReq{
		Telephone:         os.Getenv("CLASSIN_TESTMOBILE"),
		Email:             "",
		Nickname:          os.Getenv("CLASSIN_NICKNAME"),
		Password:          os.Getenv("CLASSIN_PASSWORD"),
		Md5Password:       "",
		FileData:          "",
		AddToSchoolMember: 2,
	}
	userId, err := client.Register(&req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(userId)
}

// 学生注册测试
func TestClient_AddSchoolStudent(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	req := AddSchoolStudentReq{
		StudentAccount: os.Getenv("CLASSIN_TESTMOBILE"),
		StudentName:    os.Getenv("CLASSIN_NICKNAME"),
	}
	err := client.AddSchoolStudent(&req)
	if err != nil {
		t.Fatal(err)
	}
}

// 编辑学生测试
func TestClient_EditSchoolStudent(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	userId, _ := strconv.ParseInt(os.Getenv("CLASSIN_USERID"), 10, 64)
	req := EditSchoolStudentReq{
		StudentUid:  userId,
		StudentName: os.Getenv("CLASSIN_NICKNAME"),
	}
	err := client.EditSchoolStudent(&req)
	if err != nil {
		t.Fatal(err)
	}
}

// 同步学生昵称测试
func TestClient_ModifyCourseStudentNickName(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	userId, _ := strconv.ParseInt(os.Getenv("CLASSIN_USERID"), 10, 64)
	uids := []int64{userId}
	result, err := client.ModifyCourseStudentNickName(uids)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

// 添加老师测试
func TestClient_AddTeacher(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	req := AddTeacherReq{
		TeacherAccount: os.Getenv("CLASSIN_TESTMOBILE"),
		TeacherName:    os.Getenv("CLASSIN_NICKNAME"),
		FileData:       "",
	}
	result, err := client.AddTeacher(&req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

// 修改老师测试
func TestClient_EditTeacher(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	teacherUid, _ := strconv.ParseInt(os.Getenv("CLASSIN_USERID"), 10, 64)
	req := EditTeacherReq{
		TeacherUid:  teacherUid,
		TeacherName: os.Getenv("CLASSIN_NICKNAME"),
		FileData:    "",
	}
	err := client.EditTeacher(&req)
	if err != nil {
		t.Fatal(err)
	}
}

// 停用老师测试
func TestClient_StopUsingTeacher(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	teacherUid, _ := strconv.ParseInt(os.Getenv("CLASSIN_USERID"), 10, 64)
	err := client.StopUsingTeacher(teacherUid)
	if err != nil {
		t.Fatal(err)
	}
}

// 启用老师测试
func TestClient_RestartUsingTeacher(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	teacherUid, _ := strconv.ParseInt(os.Getenv("CLASSIN_USERID"), 10, 64)
	err := client.RestartUsingTeacher(teacherUid)
	if err != nil {
		t.Fatal(err)
	}
}
