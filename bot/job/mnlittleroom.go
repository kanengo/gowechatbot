package job

import (
	"fmt"
	"time"

	"github.com/eatmoreapple/openwechat"
)

var numberEmoji = []string{
	"0️⃣",
	"1️⃣",
	"2️⃣",
	"3️⃣",
	"4️⃣",
	"5️⃣",
	"6️⃣",
	"7️⃣",
	"8️⃣",
	"9️⃣",
}

func parseNumberEmoji(number int64) string {
	emojiText := make([]string, 0, 3)

	for {
		n := number % 10
		emojiText = append(emojiText, numberEmoji[n])
		number /= 10
		if number == 0 {
			break
		}
	}

	text := ""

	for i := len(emojiText) - 1; i >= 0; i-- {
		text += emojiText[i]
	}

	return text

}

func SendLoveWord(group *openwechat.Group) {
	meetDayUnix := int64(1665072000)
	loveDayUnix := int64(1668787200)

	now := time.Now().Unix()
	meetDayDuration := (now-meetDayUnix)/(3600*24) + 1
	loveDayaDuration := (now - loveDayUnix) / (3600 * 24)

	msg := fmt.Sprintf("💖💘💗早上好啊M宝,今天是N宝和你相识的第%s天，已经相爱了%s天啦！！要一直走下去哦~💗💘💖",
		parseNumberEmoji(meetDayDuration), parseNumberEmoji(loveDayaDuration))

	group.SendText(msg)

}
