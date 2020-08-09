package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // mapについて、formに入力された情報と、URLに同一のキーがある場合はそれも含まれるが、formに入力したものが前にくる
	// ParseFormはurlencodeしかサポートしていないので、multipart/form-dataが指定されている場合はParseMultipartFormを使う
	// Request.FormValue(kye)もあるが、複数の要素がある場合に先頭のものしか返らない（フォームの情報が前に来るのでまぁよいが）
	fmt.Fprintln(w, r.Form)
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
