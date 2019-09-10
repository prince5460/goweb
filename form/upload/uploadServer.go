package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

//处理文件上传
//客户端通过multipart.Write把文件的文本流写入一个缓存中，
// 然后调用http的Post方法把缓存传到服务器。
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("upload.html")
		t.Execute(w, token)
	} else {
		//里面的参数表示maxMemory,服务端调用r.ParseMultipartForm,
		// 把上传的文件存储在内存和临时文件中
		r.ParseMultipartForm(32 << 20)
		srcfile, header, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println("err:", err)
		}
		defer srcfile.Close()
		fmt.Fprintf(w, "%v", header.Header)
		dstfile, err := os.OpenFile("./file/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		io.Copy(dstfile, srcfile)
		defer dstfile.Close()

	}
}

func main() {
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
