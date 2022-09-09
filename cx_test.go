package cx

import (
	"fmt"
	"testing"
)

func TestStart(t *testing.T) {
	cookie, err := Login("", "")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	user, err := GetUserInfo(cookie)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(user)
}
