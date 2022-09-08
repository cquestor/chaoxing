package cx

import (
	"fmt"
	"testing"
)

func TestStart(t *testing.T) {
	result, err := Login("", "")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}
