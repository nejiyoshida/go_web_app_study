package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	// Request構造体にheaderやらbodyやらのフィールドがあるから、それをとってくる
	h := r.Header // mapでとれるので、keyを指定すると必要分だけとれる
	fmt.Fprintln(w, h)
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/headers/", headers)
	server.ListenAndServe()
}
