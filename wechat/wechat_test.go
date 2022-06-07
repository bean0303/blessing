package wechat

import (
	"testing"
)

func TestWechat_SendImageMessage(t *testing.T) {

	imgname := "../out.png"

	tests := []struct {
		name    string
		imgname string
		wantErr bool
	}{
		{"test", imgname, false},
	}
	for _, tt := range tests {
		weChat := &Wechat{}
		weChat.Login()
		friend := weChat.Search("鹏飞")

		weChat.SendImageMessage(friend, tt.imgname)

	}
}
