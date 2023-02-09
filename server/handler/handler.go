package handler

import (
	"crypto/sha1"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
)

const token = "KywET0KTeRw7TnabVNpCMq0aRdrPTDe"

type MpHandler struct {
}

func (m *MpHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	values := request.URL.Query()

	signature := values.Get("signature")
	timestamp := values.Get("timestamp")
	nonce := values.Get("nonce")

	log.Println("[Rev]:", signature, timestamp, nonce)

	keys := []string{token, timestamp, nonce}
	sort.Strings(keys)

	buf := strings.Builder{}

	for _, k := range keys {
		buf.WriteString(k)
	}

	tmpSign := fmt.Sprintf("%x", sha1.Sum([]byte(buf.String())))
	log.Println("[tmpSign]", tmpSign, signature)
	if tmpSign != signature {
		writer.WriteHeader(http.StatusForbidden)
	}
	log.Println("[接收成功]")
	writer.WriteHeader(http.StatusOK)
}
