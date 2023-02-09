package message

import (
	"fmt"
	"log"
	"strings"

	"gowechatbot/bot/job"

	"github.com/eatmoreapple/openwechat"
)

const (
	MNLittleRoom = "MN小房间"
)

func Handler(msg *openwechat.Message) {
	if msg.IsComeFromGroup() { //群组
		user, err := msg.Sender()
		if err != nil {
			log.Println("error:", err)
			return
		}
		group := &openwechat.Group{
			User: user,
		}
		fmt.Println("[MN小房间]", group.ID(), group.Uin, msg.Content)
		switch group.NickName {
		case MNLittleRoom:
			if msg.IsAt() {
				fmt.Println("[收到@消息]:", msg.Content)
			}
			if strings.Contains(msg.Content, "#testlove") {
				job.SendLoveWordMorning(group)
				job.TakeMedicine(group)
				job.DrinkWater(group)
			}
		default:

		}

	} else {
		if msg.IsText() && msg.Content == "ping" {
			msg.ReplyText("pong")
		}
	}
}
