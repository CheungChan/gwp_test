package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func handleHelloFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HelloFunc!")
}
func handleWorldFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "WorldFunc!")
}

// 串联处理器函数
func log(h http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Fprintln(writer, "Hanlder function called - "+name)
		h(writer, request)
	}
}

type HelloHandler struct{}

func (handler HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}

type WorldHandler struct{}

func (handler WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "World!")
}

//串联处理器
func protect(h http.Handler) http.Handler {
	// http.HandlerFunc()将一个处理器函数转换为了处理器
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handler called - %T\n", h)
		h.ServeHTTP(w, r)
	})
}
func main() {
	helloHandler := HelloHandler{}
	worldHandler := WorldHandler{}
	server := http.Server{Addr: ":8000"}
	http.HandleFunc("/helloFunc", handleHelloFunc)
	http.HandleFunc("/worldFunc", log(handleWorldFunc))
	http.Handle("/hello/", helloHandler)
	http.Handle("/world", protect(worldHandler))
	server.ListenAndServe()
}
