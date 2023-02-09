package main

import (
	"log"
	"net/http"

	"gowechatbot/server/handler"
)

func main() {
	http.Handle("/mp/v1", &handler.MpHandler{})
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalln(err)
	}
}
