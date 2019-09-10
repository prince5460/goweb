package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func login1(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "GET" {
		f, _ := template.ParseFiles("login.html")
		f.Execute(w, nil)
	} else {
		username := r.Form.Get("username")
		//func HTMLEscape(w io.Writer, b []byte) //把b进行转义之后写到w
		//func HTMLEscapeString(s string) string //转义s之后返回结果字符串
		//func HTMLEscaper(args ...interface{}) string //支持多个参数一起转义，返回结果字符串

		fmt.Println("username:", template.HTMLEscapeString(username)) //输出到服务器
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))

		template.HTMLEscape(w, []byte(username)) //输出到客户端
	}

}

func login2(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "GET" {
		f, _ := template.ParseFiles("login.html")
		f.Execute(w, nil)
	} else {
		username := r.Form.Get("username")
		fmt.Println(username)

		//进行模板解析
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		err = t.ExecuteTemplate(w, "T", username)
		//err = t.ExecuteTemplate(w, "T", template.HTML(username))//不转义

		//如果转义失败 抛出对应错误 终止程序
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	http.HandleFunc("/login", login2)        //设置访问的路由
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
