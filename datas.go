package cx

const (
	URL_LOGIN     = "http://passport2.chaoxing.com/fanyalogin"
	URL_USER_INFO = "http://passport2.chaoxing.com/mooc/accountManage"
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
	Name   string //姓名
	Id     string //ID
	Sex    Sex    //性别
	Phone  string //手机号
	School string //学校
	Sno    string //学号
}
