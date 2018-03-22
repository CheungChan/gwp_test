package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process_bool(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("/Users/chenzhang/GolandProjects/src/gwp_test/show_content/tmpl.html")
	if err != nil {
		fmt.Println(err)
	}
	rand.Seed(time.Now().Unix())
	b := rand.Intn(10) > 5
	t.Execute(w, b)
}
func process_range(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("/Users/chenzhang/GolandProjects/src/gwp_test/show_content/range.html")
	if err != nil {
		fmt.Println(err)
	}
	//daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t.Execute(w, nil)
}
func process_with(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("/Users/chenzhang/GolandProjects/src/gwp_test/show_content/with.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, "hello")
}
func process_include(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("/Users/chenzhang/GolandProjects/src/gwp_test/show_content/t1.html",
		"/Users/chenzhang/GolandProjects/src/gwp_test/show_content/t2.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, "Hello World!")
}
func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func process_func(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("/Users/chenzhang/GolandProjects/src/gwp_test/show_content/use_func.html").Funcs(funcMap)
	t, err := t.ParseFiles("/Users/chenzhang/GolandProjects/src/gwp_test/show_content/use_func.html")
	//t, err := t.Parse("{{fdate .}}")
	// 访问之后html一直是空的,而是用Parse发现可以出来,原因是template.New()里面的和t.ParseFiles()里面的和html文件里的define里的三者
	// 必须一模一样,因为html里面没有写define名字,而没有的话默认是按文件名作为模板名字的,所有三者不一致.导致出不来.
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, time.Now())
}
func handleContext(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("/Users/chenzhang/GolandProjects/src/gwp_test/show_content/context.html")
	if err != nil {
		fmt.Println(err)
	}
	content := `I asked:<i>"What's up?"</i>"'"`
	t.Execute(w, content)
}
func process_nomakexss_html(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("/Users/chenzhang/GolandProjects/src/gwp_test/show_content/nomake_xss.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}
func process_nomakexss(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("/Users/chenzhang/GolandProjects/src/gwp_test/show_content/noxss.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, r.FormValue("comment"))
}
func process_makexss_html(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("/Users/chenzhang/GolandProjects/src/gwp_test/show_content/make_xss.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}

func process_makexss(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("/Users/chenzhang/GolandProjects/src/gwp_test/show_content/xss.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, template.HTML(r.FormValue("comment")))
}
func process_forcemakexss_html(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("/Users/chenzhang/GolandProjects/src/gwp_test/show_content/forcemake_xss.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}

func process_forcemakexss(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "0")
	t, err := template.ParseFiles("/Users/chenzhang/GolandProjects/src/gwp_test/show_content/forcexss.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, template.HTML(r.FormValue("comment")))
}
func main() {
	server := http.Server{Addr: ":8000"}
	http.HandleFunc("/process_bool", process_bool)
	http.HandleFunc("/process_range", process_range)
	http.HandleFunc("/process_with", process_with)
	http.HandleFunc("/process_include", process_include)
	http.HandleFunc("/process_func", process_func)
	http.HandleFunc("/process_context", handleContext)
	http.HandleFunc("/process_nomakexss_html", process_nomakexss_html)
	http.HandleFunc("/process_nomakexss", process_nomakexss)
	http.HandleFunc("/process_makexss_html", process_makexss_html)
	http.HandleFunc("/process_makexss", process_makexss)
	http.HandleFunc("/process_forcemakexss_html", process_forcemakexss_html)
	http.HandleFunc("/process_forcemakexss", process_forcemakexss)
	server.ListenAndServe()
}
