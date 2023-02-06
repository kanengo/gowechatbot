package message

import "github.com/eatmoreapple/openwechat"

func Handler(msg *openwechat.Message) {
	if msg.IsText() && msg.Content == "ping" {
		msg.ReplyText("pong")
	}
}
