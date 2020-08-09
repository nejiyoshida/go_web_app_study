package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんちは！")
}
func log(h http.HandlerFunc) http.HandlerFunc {
	// 引数の関数を実行する無名関数を返す
	return func(w http.ResponseWriter, r *http.Request) {
		/*
			FuncForPCは引数で受けたプログラムカウンタの情報を戻す。今回はreflectで複製したやつのインタフェースを戻して、
			ポインタ取ってきてそこのNameを取り出してる
		*/
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println(name + " がよばれたぞい")
		h(w, r) // 引数の関数を呼んだ
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", log(hello))
	server.ListenAndServe()
}
