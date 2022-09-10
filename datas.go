package cx

const (
	URL_LOGIN       = "http://passport2.chaoxing.com/fanyalogin"
	URL_USER_INFO   = "http://passport2.chaoxing.com/mooc/accountManage"
	URL_COURSE_LIST = "http://mooc1-1.chaoxing.com/visit/courselistdata"
)

const (
	KEY_DES = "u2oh6Vu^HWe40fj"
)

const (
	MALE   = 1
	FEMALE = 0
)

type Sex struct {
	Value int
	Text  string
}

type User struct {
	Name   string
	Id     string
	Sex    Sex
	Phone  string
	School string
	Sno    string
}

type Course struct {
	CourseId    string
	ClazzId     string
	PersonId    string
	CourseName  string
	TeacherName string
	ClazzName   string
}
