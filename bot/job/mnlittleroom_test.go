package job

import (
	"fmt"
	"testing"
	"time"

	"github.com/robfig/cron/v3"
)

func TestParseNumberEmoji(t *testing.T) {
	fmt.Println(parseNumberEmoji(123))
}

type Foo struct {
	Id int
}

func increaseFoo(foo *Foo) {
	foo.Id += 1
	fmt.Println(foo.Id)
}

func TestCronJob(t *testing.T) {
	c := cron.New()

	foo := &Foo{}
	_, err := c.AddFunc("*/1 * * * *", func() {
		fmt.Println("=====")
		increaseFoo(foo)
	})
	if err != nil {
		t.Error(err)
		return
	}
	c.Start()
	time.Sleep(time.Hour)
}
