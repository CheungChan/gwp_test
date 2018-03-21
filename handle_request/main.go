package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Post struct {
	User    string
	Threads []string
}

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
	a := r.Header["Accept-Encoding"]
	fmt.Fprintln(w, a)
	b := r.Header.Get("Accept-Encoding")
	fmt.Fprintln(w, b)
}
func body(w http.ResponseWriter, r *http.Request) {
	length := r.ContentLength
	body := make([]byte, length)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}
func process(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//fmt.Fprintln(w, r.Form)
	//fmt.Fprintln(w,r.Form.Get("first_name"))
	file, _, err := r.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Println(string(data))
			fmt.Fprintln(w, string(data))
		}
	}
}
func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>GWP</title><head>
<body><h1>Hello World</h1></body>
</html>`
	w.Write([]byte(str))
}
func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintln(w, "未实现")
	fmt.Fprintln(w, "用WriteHeader设置完状态码,不可以再修改Header了,但是可以修改Body")
}
func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://www.google.com.hk")
	w.WriteHeader(http.StatusFound)
}
func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "cheung chan",
		Threads: []string{"first", "second", "third"},
	}
	j, _ := json.Marshal(post)
	w.Write(j)
}
func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "GWP",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "MPG",
		HttpOnly: true,
	}
	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())
}
func setCookie2(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "GWP",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "MPG",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}
func getCookie(w http.ResponseWriter, r *http.Request) {
	c1 := r.Header["Cookie"]
	c2 := r.Header.Get("Cookie")
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, c2)
	c3, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "找不到第一个Cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c3)
	fmt.Fprintln(w, c3.HttpOnly)
	fmt.Fprintln(w, c3.Domain)
	fmt.Fprintln(w, cs)
}
func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello world")
	c := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "no message found")
		}
	} else {
		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}
func main() {
	server := http.Server{Addr: ":8000"}
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/body", body)
	http.HandleFunc("/process/", process)
	http.HandleFunc("/writeExample", writeExample)
	http.HandleFunc("/writeHeaderExample", writeHeaderExample)
	http.HandleFunc("/headerExample", headerExample)
	http.HandleFunc("/jsonExample", jsonExample)
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/setCookie2", setCookie2)
	http.HandleFunc("/getCookie", getCookie)
	http.HandleFunc("/setMessage", setMessage)
	http.HandleFunc("/showMessage", showMessage)
	server.ListenAndServe()

}
