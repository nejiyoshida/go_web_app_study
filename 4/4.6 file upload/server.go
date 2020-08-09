package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)                        // マルチパートのフォームから指定のバイト分だけ取得
	fileHeader := r.MultipartForm.File["uploaded"][0] // Fileフィールドの頭を取得
	file, err := fileHeader.Open()                    // ファイルを取得
	if err == nil {
		data, err := ioutil.ReadAll(file) // dataにバイト配列を読み取る
		if err == nil {
			fmt.Fprintln(w, string(data)) // sringに変換して出力
		}
	}
}

// アップロードされるファイルが一つだけならこちらの方が早い
func process2(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("uploaded") // 第二戻り値が上の関数のfileHeaderに相当
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
