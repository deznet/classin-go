package classin

import (
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
	"strconv"
	"testing"
)

// 添加课程班测试
// 281197321
func TestClient_AddCourse(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	req := new(AddCourseReq)
	req.CourseName = os.Getenv("CLASSIN_COURSENAME")
	courseId, err := client.AddCourse(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(courseId)
}

// 修改课程班测试
func TestClient_EditCourse(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	req := new(EditCourseReq)
	req.CourseName = os.Getenv("CLASSIN_COURSENAME")
	req.CourseId, _ = strconv.ParseInt(os.Getenv("CLASSIN_COURSEID"), 10, 64)
	err := client.EditCourse(req)
	if err != nil {
		log.Fatal(err)
	}
}

// 结束课程班测试
func TestClient_EndCourse(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	courseId, _ := strconv.ParseInt(os.Getenv("CLASSIN_COURSEID"), 10, 64)
	err := client.EndCourse(courseId)
	if err != nil {
		log.Fatal(err)
	}
}

// 课程添加学生测试
func TestClient_AddCourseStudent(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	req := new(AddCourseStudentReq)
	req.StudentUid, _ = strconv.ParseInt(os.Getenv("CLASSIN_USERID"), 10, 64)
	req.CourseId, _ = strconv.ParseInt(os.Getenv("CLASSIN_COURSEID"), 10, 64)
	req.StudentName = os.Getenv("CLASSIN_NICKNAME")
	req.Identity = 1
	err := client.AddCourseStudent(req)
	if err != nil {
		log.Fatal(err)
	}
}

// 课程删除学生测试
func TestClient_DelCourseStudent(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	req := new(DelCourseStudentReq)
	req.StudentUid, _ = strconv.ParseInt(os.Getenv("CLASSIN_USERID"), 10, 64)
	req.CourseId, _ = strconv.ParseInt(os.Getenv("CLASSIN_COURSEID"), 10, 64)
	req.Identity = 1
	err := client.DelCourseStudent(req)
	if err != nil {
		log.Fatal(err)
	}
}

// 测试课程下添加学生/旁听（多个）
func TestClient_AddCourseStudentMulti(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	courseId, _ := strconv.ParseInt(os.Getenv("CLASSIN_COURSEID"), 10, 64)
	students := make([]*CourseStuJson, 0)
	stu := new(CourseStuJson)
	stu.UID, _ = strconv.ParseInt(os.Getenv("CLASSIN_USERID"), 10, 64)
	stu.Name = os.Getenv("CLASSIN_NICKNAME")
	stu.CustomColumn = "1"
	students = append(students, stu)
	req := AddCourseStudentMultiReq{
		CourseId:    courseId,
		Identity:    1,
		StudentJson: students,
	}
	resp, err := client.AddCourseStudentMulti(&req)
	if err != nil {
		log.Fatal(err)
	}
	t.Log(resp)
}

// 课程下删除学生/旁听（多个）测试
func TestClient_DelCourseStudentMulti(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	courseId, _ := strconv.ParseInt(os.Getenv("CLASSIN_COURSEID"), 10, 64)
	stuId, _ := strconv.ParseInt(os.Getenv("CLASSIN_USERID"), 10, 64)
	req := DelCourseStudentMultiReq{
		CourseId:       courseId,
		Identity:       1,
		StudentUidJson: []int64{stuId},
	}
	resp, err := client.DelCourseStudentMulti(&req)
	if err != nil {
		log.Fatal(err)
	}
	t.Log(resp)
}

// 修改群成员的班级昵称测试
func TestClient_ModifyGroupMemberNickname(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	courseId, _ := strconv.ParseInt(os.Getenv("CLASSIN_COURSEID"), 10, 64)
	err := client.ModifyGroupMemberNickname(courseId)
	if err != nil {
		log.Fatal(err)
	}
}

// 添加课程教师测试
func TestClient_AddCourseTeacher(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	courseId, _ := strconv.ParseInt(os.Getenv("CLASSIN_COURSEID"), 10, 64)
	teacherUid, _ := strconv.ParseInt(os.Getenv("CLASSIN_USERID"), 10, 64)
	req := AddCourseTeacherReq{
		CourseId:    courseId,
		TeacherUids: []int64{teacherUid},
	}
	resp, err := client.AddCourseTeacher(&req)
	if err != nil {
		log.Fatal(err)
	}
	t.Log(resp)
}

// 移除课程老师测试
func TestClient_RemoveCourseTeacher(t *testing.T) {
	client := NewClient(os.Getenv("CLASSIN_SID"), os.Getenv("CLASSIN_SECRET"))
	courseId, _ := strconv.ParseInt(os.Getenv("CLASSIN_COURSEID"), 10, 64)
	teacherUid, _ := strconv.ParseInt(os.Getenv("CLASSIN_USERID"), 10, 64)
	req := RemoveCourseTeacherReq{
		CourseId:   courseId,
		TeacherUid: teacherUid,
	}
	err := client.RemoveCourseTeacher(&req)
	if err != nil {
		log.Fatal(err)
	}
}
