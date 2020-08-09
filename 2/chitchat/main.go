package main

import (
	"net/http"
)

// TODO 他の章をやり終えたら完成させる

func main() {
	mux := http.NewServeMux() // リクエストに応じて各ハンドラに回すためのマルチプレクサ

	files := http.FileServer(http.Dir("/public"))

	mux.Handle("/static/", http.StripPrefix("/static/", files)) // "/static/"で始まるリクエストは、URLからこの部分をむしり取って、/public/AAA~というファイルを返す

	mux.HandleFunc("/", index) // "/"のURLに対するリクエストは、indexで示されるハンドラ関数に回す
	mux.HandleFunc("/err", err)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

/*
func index(w http.ResponseWriter, r *http.Request) {
	files := []string{"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads()
	if err == nil {
		templates.ExecuteTemplate(W, "layout", threads)
	}
}
*/
