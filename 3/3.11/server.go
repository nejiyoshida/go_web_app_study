package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんちは！")
}

func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("ハンドラが呼び出されました - %T\n", h)
		h.ServeHTTP(w, r)
	})
}

func protect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 	ここで何か処理やバリデーションを行って、問題がなければ実行、みたいな
		// ~~
		h.ServeHTTP(w, r)
	})
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	hello := HelloHandler{}
	http.Handle("/hello", protect(log(hello)))
	server.ListenAndServe()
}
