package main

import (
	"fmt"
	"gwp_test/data"
	"html/template"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	files := http.FileServer(http.Dir("/Users/chenzhang/GolandProjects/src/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	server := &http.Server{Addr: ":8000", Handler: mux}
	server.ListenAndServe()

}
func index(w http.ResponseWriter, r *http.Request) {
	var templates *template.Template
	var tmpl_files []string
	_, err := session(w, r)
	if err != nil {
		tmpl_files = []string{"/Users/chenzhang/GolandProjects/src/templates/layout.html",
			"/Users/chenzhang/GolandProjects/src/templates/public.navbar.html",
			"/Users/chenzhang/GolandProjects/src/templates/index.html"}
	} else {
		tmpl_files = []string{"/Users/chenzhang/GolandProjects/src/templates/layout.html",
			"/Users/chenzhang/GolandProjects/src/templates/private.navbar.html",
			"/Users/chenzhang/GolandProjects/src/templates/index.html"}
	}
	templates = template.Must(template.ParseFiles(tmpl_files...))
	threads, err := data.Threads();
	if err != nil {
		fmt.Println(err)
		return
	}
	templates.ExecuteTemplate(w, "layout", threads)
}
