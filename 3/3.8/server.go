package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "世界")
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	// いちいち構造体に紐づく形にしなくても、関数を用意しておいてHandleFuncに投げると、ハンドラに変換して登録してくれる
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}
