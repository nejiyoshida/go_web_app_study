package main

import (
	"fmt"
	"net/http"
)

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>responseに書き込むテスト</title></head>
<body><h1> テスト </h1></body>
</html>`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501) // ステータスコードを指定
	fmt.Fprintln(w, "そこはまだ実装してないんじゃぁ")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302) // ステータスコードを書き込む前にHeader()を実行すること
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/hoge", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	server.ListenAndServe()
}
