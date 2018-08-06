package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"testing"
)

type A struct {
	Ax int `json:"ax"`
}

type User struct {
	UserId int `json:"user_id"`
	A
}

func Test_test(t *testing.T) {
	user := &User{
		UserId: 0,
		A: A{
			Ax: 0,
		},
	}
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	fmt.Println("user=", string(data))

	smsMes := &message.SmsMes{
		Content: "",
		User: message.User{
			UserId:     0,
			UserPwd:    "",
			UserName:   "",
			UserStatus: 0,
			Sex:        "",
		},
	}
	data, err = json.Marshal(smsMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	fmt.Println("smsMes=", string(data))
}
