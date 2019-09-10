package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

//简易路由
func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHello(w, r)
		return
	}
	http.NotFound(w, r)
	return

}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello,route!")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":8080", mux)
}
