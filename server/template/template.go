package template

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"gowechatbot/server/token"
)

type Meta struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type Model struct {
	ToUser     string          `json:"touser"`
	TemplateId string          `json:"template_id"`
	Data       map[string]Meta `json:"data"`
}

func SendTemplateMsg(tpl *Model) error {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", token.GetAccessToken())
	b, _ := json.Marshal(tpl)
	fmt.Println("data:", string(b))
	_, err := http.Post(url, "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}

	return nil
}
