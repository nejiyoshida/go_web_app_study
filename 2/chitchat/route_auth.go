package main

import (
	//"github.com/nejiyoshida/go_app/2/chitchat/data"
	"net/http"
)

// POST /authenticate
func authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()                                    // POSTで送られてくるフォームからの入力を request.Form()で取れるようになる
	user, err := data.UserByEmail(request.PostFormValue("email")) // request（フォームで送られてきたやつ）のemailがキーの値をUserByEmailに渡すことでUser構造体を得る(DBからユーザ情報を引いてくる)
	if err != nil {
		danger(err, "Cannot find User")
	}
	if user.Password == data.Encrypt(request.PostFormValue("password")) { // 引いてきたユーザのパスワード（暗号化済み）が、ユーザが入力したものを暗号化したのと同じなら認証成功
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302) // 認証に失敗したのでログイン画面に戻す
	}
}

// GET /login
func login(writer http.ResponseWriter, request *http.Request) {

}
