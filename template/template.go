package main

import (
	"html/template"
	"os"
)

type Person struct {
	Username string
	email    string //未导出的字段，首字母是小写
}

func main() {
	t := template.New("example")
	parse, _ := t.Parse("hello,{{ .Username }},{{ .email }}")
	p := Person{
		Username: "Tom",
		email:    "abc@123.com",
	}
	parse.Execute(os.Stdout, p)
}
