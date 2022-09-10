package cx

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
)

func Login(uname, password string) (map[string]string, error) {
	password, err := desEncrypt([]byte(password), []byte(KEY_DES))
	if err != nil {
		return nil, err
	}
	payload := "fid=-1&uname=" + uname + "&password=" + password + "&refer=http%253A%252F%252Fi.chaoxing.com&t=true&forbidotherlogin=0&validate=&doubleFactorLogin=0&independentId=0"
	req, _ := http.NewRequest(http.MethodPost, URL_LOGIN, strings.NewReader(payload))
	addRequestHeader(req, nil, "application/x-www-form-urlencoded; charset=UTF-8")
	client := &http.Client{Timeout: time.Second * 10, CheckRedirect: disallowRedirect}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var respJson struct {
		Url    string `json:"url"`
		Status bool   `json:"status"`
		Msg    string `json:"msg2"`
	}
	json.Unmarshal(body, &respJson)
	if !respJson.Status {
		return nil, errors.New(respJson.Msg)
	}
	cookie := make(map[string]string)
	for _, v := range resp.Cookies() {
		cookie[v.Name] = v.Value
	}
	return cookie, nil
}

func GetUserInfo(cookie map[string]string) (*User, error) {
	req, _ := http.NewRequest(http.MethodGet, URL_USER_INFO, nil)
	addRequestHeader(req, cookie, "")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	userName := htmlquery.InnerText(htmlquery.FindOne(doc, `//p[@id="messageName"]`))
	userID := htmlquery.InnerText(htmlquery.FindOne(doc, `//p[@id="uid"]`))
	userSex, _ := strconv.Atoi(htmlquery.SelectAttr(htmlquery.FindOne(doc, `//p[contains(@class, "sex")]//i[contains(@class, "checked")]`), "value"))
	var userSexString string
	if userSex == MALE {
		userSexString = "男"
	} else {
		userSexString = "女"
	}
	userPhone := htmlquery.InnerText(htmlquery.FindOne(doc, `//span[@id="messagePhone"]`))
	schoolName := strings.TrimSpace(strings.Split(htmlquery.InnerText(htmlquery.FindOne(doc, `//ul[@id="messageFid"]`)), "\n")[1])
	schoolSno := strings.Split(htmlquery.InnerText(htmlquery.FindOne(doc, `//p[@class="xuehao"]`)), ":")[1]
	user := User{Name: userName, Id: userID, Sex: Sex{Value: userSex, Text: userSexString}, Phone: userPhone, School: schoolName, Sno: schoolSno}
	return &user, nil
}

func GetCourseList(cookie map[string]string) ([]Course, error) {
	payload := "courseType=1&courseFolderId=0&baseEducation=0&superstarClass=&courseFolderSize=0"
	req, _ := http.NewRequest(http.MethodPost, URL_COURSE_LIST, strings.NewReader(payload))
	addRequestHeader(req, cookie, "application/x-www-form-urlencoded; charset=UTF-8")
	client := &http.Client{Timeout: time.Second * 10, CheckRedirect: disallowRedirect}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	courseListDome := htmlquery.Find(doc, `//ul[@id="courseList"]/li[contains(@class, "course")]`)
	var courseList []Course
	for _, v := range courseListDome {
		courseId := htmlquery.SelectAttr(v, "courseid")
		clazzId := htmlquery.SelectAttr(v, "clazzid")
		personId := htmlquery.SelectAttr(v, "personid")
		courseName := htmlquery.SelectAttr(htmlquery.FindOne(v, `.//span[contains(@class, "course-name")]`), "title")
		teacherName := htmlquery.SelectAttr(htmlquery.FindOne(v, `.//p[@class="line2"]`), "title")
		clazzName := strings.Split(htmlquery.InnerText(htmlquery.FindOne(v, `.//p[@class="overHidden1"]`)), "：")[1]
		courseList = append(courseList, Course{CourseId: courseId, ClazzId: clazzId, PersonId: personId, CourseName: courseName, TeacherName: teacherName, ClazzName: clazzName})
	}
	return courseList, nil
}

func disallowRedirect(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}

func addRequestHeader(req *http.Request, cookie map[string]string, contentType string) {
	if cookie != nil {
		cookieString := ""
		for k, v := range cookie {
			cookieString += k + "=" + v + ";"
		}
		req.Header.Add("Cookie", cookieString)
	}
	if contentType != "" {
		req.Header.Add("Content-Type", contentType)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
}
