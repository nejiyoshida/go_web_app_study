package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}

	c2 := http.Cookie{

		Name:     "name_sono_2", // ちなみにNameのstringに空白が入ってるのはダメみたい(´;ω;`)。どこかのタイミングでmapのkey扱いになってるとか？ -> URLエンコードすればOKか？
		Value:    "value sono 2",
		HttpOnly: true,
	}

	http.SetCookie(w, &c1) // w.Header().Set("Set-Cookie", c1.String()) でもOK
	http.SetCookie(w, &c2) // w.Header().Add("Set-Cookie", c2.String()) でもOK
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("name_sono_2")
	if err != nil {
		fmt.Fprintln(w, "cookie 取れなかったわ")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
}

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("ふらっしゅ！")
	c := http.Cookie{
		Name:  "flash!!!",
		Value: base64.URLEncoding.EncodeToString(msg), // メッセージに空白やらその他記号を入れる場合があるのでURLエンコードする
	}
	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash!!!")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "メッセージがないぞ")
		}

	} else {
		// 既存のcookieを置き換えるが、期限を過去に指定してるのですぐに賞味期限が切れてこのcookieも削除される。
		// リロードするとcookieがないからerrのカッコ内の処理になる。
		// ログインに成功したとか失敗したとかそういう一回こっきりの情報を出したりするのに使う。（フラッシュメッセージ）
		rc := http.Cookie{
			Name:    "flash!!!",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)
	server.ListenAndServe()
}
