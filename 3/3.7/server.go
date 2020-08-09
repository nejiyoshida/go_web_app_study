package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello~~~~ ")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "世界")
}

func main() {
	hello := HelloHandler{}
	world := WorldHandler{}

	server := http.Server{
		Addr: ":8080", // 省略できる
		// Handlerを指定しないと、DefaultServeMuxが利用される
	}

	// 明示されてないけど、DefaultServeMuxに各ハンドラを登録している
	http.Handle("/hello", &hello)
	http.Handle("/world", &world)

	server.ListenAndServe()
}
