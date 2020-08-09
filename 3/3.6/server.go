package main

import (
	"fmt"
	"net/http"
)

/*
Goにおけるハンドラ
ServeHTTPというメソッドを持っているインタフェースのこと。これをもっていればなんでもハンドラ。
ListenAndServeの第二引数にnilを与えると、DefaultServeMuxが利用されるが、これもServeHTTPを持っていて、URLに応じてリクエストをハンドラに投げる。
*/
type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// どんなURLにアクセスが来ても、これを返してしまう
	// 普通はURLごとにアクションを変えたいので、ListenAndServeの第二引数はデフォルトのまま利用することが多い
	fmt.Fprintf(w, "Hello Hogehoge!!!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}

	server.ListenAndServe()
}
