package template

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"
)

const tokenUrl = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wx5a7eee37d34af1f3&secret=5bf5091ce2d6d91203e34e3ccb3780ce"

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

func TestSendTempLate(t *testing.T) {
	resp, err := http.Get(tokenUrl)
	if err != nil {
		t.Error(err)
		return
	}
	buf, err := io.ReadAll(resp.Body)
	resp.Body.Close()

	var accessToken AccessToken
	_ = json.Unmarshal(buf, &accessToken)

	fmt.Println(accessToken.AccessToken, "+++", accessToken.ExpiresIn)

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", accessToken.AccessToken)
	data := map[string]any{
		"touser":      "oiQVG6Pu3nHiXQPupLc-54EF7K1E",
		"template_id": "g6aQaed0V-PQcsejWOt98US-J36A-pTyYhb8cJ-NOjw",
		"data": map[string]any{
			"name": map[string]string{
				"value": "leeka",
				"color": "#173177",
			},
		},
	}
	b, _ := json.Marshal(data)

	http.Post(url, "application/json", bytes.NewReader(b))
}

func TestSendTemplateMsg(t *testing.T) {
	time.Sleep(time.Second)
	err := SendTemplateMsg(&Model{
		ToUser:     "oiQVG6Pu3nHiXQPupLc-54EF7K1E",
		TemplateId: "bmuRJaJNhBsDA8cxZ2ElV4iCjYFog2lqEOYylluZd94",
		Data: map[string]Meta{
			"medicine": {
				Value: "小优",
				Color: "#e61025",
			},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
}
