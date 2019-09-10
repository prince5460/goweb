package main

import (
	"io"
	"net/http"
)

//根据不同的URL，可以执行不同的handler

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello,handler!")
	})

	mux.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Bye!")
	})

	//重定向
	mux.HandleFunc("/baidu", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "http://www.baidu.com", http.StatusTemporaryRedirect)
	})

	mux.HandleFunc("/", sayhello)

	http.ListenAndServe(":8080", mux)
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello,World!")
}
