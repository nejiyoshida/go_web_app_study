package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter" // URLとのパターンマッチに変数とかが使えるし軽い
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello!!!m %s\n", p.ByName("name"))
}

func main() {
	mux := httprouter.New()
	// 標準ライブラリのmuxであっても、/{pass}/ の形にしておけば、完全一致しない場合に前方一致でリクエストをルーティングできる
	mux.GET("/hello/:name", hello)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
