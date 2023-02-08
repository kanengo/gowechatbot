package message

import (
	"fmt"
	"log"
	"strings"

	"gowechatbot/bot/job"

	"github.com/eatmoreapple/openwechat"
)

const (
	MNLittleRoom = "754116689"
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
		switch group.ID() {
		case MNLittleRoom:
			if msg.IsAt() {
				fmt.Println("[收到@消息]:", msg.Content)
			}
			if strings.Contains(msg.Content, "#testlove") {
				job.SendLoveWordMorning(group)
			}
		default:

		}

	} else {
		if msg.IsText() && msg.Content == "ping" {
			msg.ReplyText("pong")
		}
	}
}
