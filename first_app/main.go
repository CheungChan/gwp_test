package main

import (
	"gwp_test/data"
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	server := &http.Server{Addr: ":8000", Handler: mux}
	server.ListenAndServe()

}
func index(w http.ResponseWriter, r *http.Request) {
	var templates *template.Template
	var tmpl_files []string
	_, err := session(w, r)
	if err != nil {
		tmpl_files = []string{"templates/layout.html", "templates/public.navbar.html", "templates/index.html"}
	} else {
		tmpl_files = []string{"templates/layout.html", "templates/private.navbar.html", "templates/index.html"}
	}
	templates = template.Must(template.ParseFiles(tmpl_files...))
	threads, err := data.Threads();
	if err != nil {
		fmt.Println(err)
		return
	}
	templates.ExecuteTemplate(w, "layout", threads)
}
