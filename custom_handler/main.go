package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {
}

func (handler MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
func main() {
	myhandler := MyHandler{}
	server := http.Server{
		Addr:    ":8000",
		Handler: &myhandler,
	}
	server.ListenAndServe()
}
