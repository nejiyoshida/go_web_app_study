package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process_if(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl_if.html")
	rand.Seed(time.Now().Unix())
	n := rand.Intn(10)
	fmt.Println(n)
	t.Execute(w, n > 5) // boolをテンプレートに渡した。
}

func process_iterate(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl_iterate.html")
	chars := []string{"rangeで", "受け取った要素を", "表示", "する"}
	//empty := []string{} // こっちだとelseの方を表示
	t.Execute(w, chars) // boolをテンプレートに渡した。
}

func process_set(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl_set.html")
	t.Execute(w, "hogehoge")
}

func process_include(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl_base.html", "tmpl_part.html")
	t.Execute(w, "hogehoge")
}

// simple pipeline
func process_spipe(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl_spipe.html")
	t.Execute(w, "")
}

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func process_func(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}   // FuncMap型（キーが関数名で、valueが実際の関数）
	t := template.New("tmpl_func.html").Funcs(funcMap) // .Funcs(funcMap)で、テンプレートに付与
	t, _ = t.ParseFiles("tmpl_func.html")              // 付与した後にParse（当たり前だが）
	t.Execute(w, time.Now())                           // formatDataに渡すものを引数にしてる
}

func process_context(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl_context.html")
	content := `I asked: <i>"What's up?"</i>`
	t.Execute(w, content)
}

func process_xss(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "0")
	t, _ := template.ParseFiles("tmpl_xss.html")
	//t.Execute(w, r.FormValue("comment")) // form.htmlのcommentからの入力をうけて、エスケープしてtmpl_xss.htmlに渡す
	t.Execute(w, template.HTML(r.FormValue("comment"))) // form.htmlのcommentからの入力をうけて、エスケープしないでtmpl_xss.htmlに渡す
	// なので、<script> alert('ほげほげ');</script>とかを埋め込み出来てしまう
}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("form.html")
	t.Execute(w, nil)
}

func process_nest(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("layout.html", "hogefuga.html")
	} else {
		t, _ = template.ParseFiles("layout.html", "foobar.html")
	}
	t.ExecuteTemplate(w, "layout", "") // layout.htmlの中で定義されたlayoutというtemplateを読み込んでいる。

}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/process_if", process_if)
	http.HandleFunc("/process_iterate", process_iterate)
	http.HandleFunc("/process_set", process_set)
	http.HandleFunc("/process_include", process_include)
	http.HandleFunc("/process_spipe", process_spipe)
	http.HandleFunc("/process_func", process_func)
	http.HandleFunc("/process_context", process_context)
	http.HandleFunc("/process_xss", process_xss)
	http.HandleFunc("/form", form)
	http.HandleFunc("/process_nest", process_nest)
	server.ListenAndServe()
}
