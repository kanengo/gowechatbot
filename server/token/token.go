package token

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const tokenUrl = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wx5a7eee37d34af1f3&secret=5bf5091ce2d6d91203e34e3ccb3780ce"

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

var accessToken string

func GetAccessTokenRemote() (string, error) {
	resp, err := http.Get(tokenUrl)
	if err != nil {
		panic(err)
	}
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	resp.Body.Close()

	var token AccessToken
	_ = json.Unmarshal(buf, &token)

	return token.AccessToken, nil
}

func init() {
	accessToken, _ = GetAccessTokenRemote()
	go func() {
		for {
			time.Sleep(5 * time.Minute)
			accessToken, _ = GetAccessTokenRemote()
		}
	}()
}

func GetAccessToken() string {
	return accessToken
}
