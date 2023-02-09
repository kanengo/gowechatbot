package main

import (
	"fmt"
	"log"

	"gowechatbot/bot/job"
	"gowechatbot/bot/message"

	"github.com/eatmoreapple/openwechat"
	"github.com/robfig/cron/v3"
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

	c := cron.New()

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	//获取所有的好友
	//friends, err := self.Friends()
	//fmt.Println(friends, err)

	// 获取所有的群组
	groups, err := self.Groups()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, group := range groups {
			if group.NickName == message.MNLittleRoom {
				fmt.Println("群组:", group.ID(), group.NickName, group.RemarkName)
				_, err := c.AddFunc("0 8 * * *", func() {
					job.SendLoveWordMorning(group)
				})
				if err != nil {
					log.Fatalln(err)
				}
				_, err = c.AddFunc("55 21 * * * ", func() {
					job.TakeMedicine(group)
				})
				if err != nil {
					log.Fatalln(err)
				}
				_, err = c.AddFunc("0 */1 * * * ", func() {
					job.TakeMedicine(group)
				})
				if err != nil {
					log.Fatalln(err)
				}
				for _, member := range group.MemberList {
					fmt.Println(member.ID(), member.UserName, member.NickName)
				}
			}
		}
	}

	c.Start()

	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	if err := bot.Block(); err != nil {
		log.Fatalln(err)
	}
}
