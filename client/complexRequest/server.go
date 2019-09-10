package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func HandlerRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("body:", string(body))
	fmt.Printf("r:%v\n", r)
	fmt.Printf("Request：Header：%v\n", r.Header)
	fmt.Printf("cookies:%v\n", r.Cookies())

	cookie, err := r.Cookie("userId")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("cookie.Name:%v,cookie.Value:%v\n", cookie.Name, cookie.Value)

	if len(r.Form["username"]) > 0 {
		username := r.Form["username"][0]
		fmt.Println("username:", username)
	}
	if len(r.Form["password"]) > 0 {
		password := r.Form["password"][0]
		fmt.Println("password:", password)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	io.WriteString(w, "login success ....")
}

func main() {
	http.HandleFunc("/login", HandlerRequest)
	err := http.ListenAndServe(":2001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
