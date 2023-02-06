package main

import (
	"fmt"
	"gowechatbot/bot/message"

	"github.com/eatmoreapple/openwechat"
)

func main() {
	bot := openwechat.DefaultBot(openwechat.Desktop)

	bot.MessageHandler = message.Handler
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	// defer reloadStorage.Close()

	// bot.HotLogin(reloadStorage)

	// 登陆
	if err := bot.Login(); err != nil {
		fmt.Println(err)
		return
	}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	//获取所有的好友
	friends, err := self.Friends()
	fmt.Println(friends, err)

	// 获取所有的群组
	groups, err := self.Groups()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, group := range groups {
			fmt.Println("群组:", group.ID(), group.NickName, group.RemarkName)
		}
	}

	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}
