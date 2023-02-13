package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	bg, _ := os.ReadFile("./public/bg.jpg")
	http.HandleFunc("/resource/bg.jpg", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get image", r.URL)
		w.Write(bg)
	})
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "tpl/index.html")
	})
	err := http.ListenAndServe(":21119", nil)
	if err != nil {
		fmt.Println(err)
	}
}
