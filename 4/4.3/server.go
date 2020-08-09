package main

import (
	"fmt"
	"net/http"
)

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body) // 上で作ったbodyに、requestのbodyを読み取る
	// POSTのデータはFormValueとかFormFileで読み取れるから、普通はそちらを使うか
	fmt.Fprintln(w, string(body))
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/body", body)
	server.ListenAndServe()
}
