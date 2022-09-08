package cx

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"
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
