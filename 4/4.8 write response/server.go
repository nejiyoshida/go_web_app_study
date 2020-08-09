package main

import "net/http"

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>responseに書き込むテスト</title></head>
<body><h1> テスト </h1></body>
</html>`
	w.Write([]byte(str))
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/write/", writeExample)
	server.ListenAndServe()
}
