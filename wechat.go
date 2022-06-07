package wechat

import (
	"fmt"
	"log"
	"os"

	"github.com/eatmoreapple/openwechat"
)

type Wechat struct {
	bot  *openwechat.Bot
	self *openwechat.Self
}

func (w *Wechat) Login() {
	messageHandler := func(msg *openwechat.Message) {
		fmt.Println(msg)
	}
	w.bot = openwechat.DefaultBot()
	//w.bot = openwechat.DefaultBot(openwechat.Desktop)

	// 注册消息处理函数
	w.bot.MessageHandler = messageHandler
	// 设置默认的登录回调
	// 可以设置通过该uuid获取到登录的二维码
	w.bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	// 登录
	if err := w.bot.Login(); err != nil {
		fmt.Println(err)
		return
	}
	var err error
	w.self, err = w.bot.GetCurrentUser()
	if err != nil {
		log.Fatalln(err)
	}
}

func (w *Wechat) SendMessage(friend *openwechat.Friend, message string) {
	if friend == nil {
		return
	}
	_, _ = friend.SendText(message)
}

func (w *Wechat) SendImageMessage(friend *openwechat.Friend, imgname string) {
	img, _ := os.Open(imgname)
	defer img.Close()
	if friend == nil {
		return
	}
	_, _ = friend.SendImage(img)
}

func (w *Wechat) Search(name string) *openwechat.Friend {
	friends, err := w.self.Friends()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	results := friends.Search(-1, func(friend *openwechat.Friend) bool {
		return friend.User.NickName == name || friend.User.RemarkName == name
	})
	//log.Printf("friend：%v", results.First().NickName)
	if results.Count() > 0 {
		return results.First()
	}
	return nil
}
