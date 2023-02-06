package message

import (
	"github.com/eatmoreapple/openwechat"
	"log"
)

const (
	MNLittleRoom = "754116689"
)

func Handler(msg *openwechat.Message) {
	if msg.IsComeFromGroup() { //群组
		group, err := msg.Sender()
		if err != nil {
			log.Println("error:", err)
			return
		}
		switch group.ID() {
		case MNLittleRoom:
		default:

		}

	} else {

	}
	//if msg.IsText() && msg.Content == "ping" {
	msg.ReplyText("pong")
	//}
}
