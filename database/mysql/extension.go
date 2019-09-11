package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

type User1 struct {
	Uid        int
	Username   string
	Departname string
	Created    string
}

func main() {
	//Get和Select是一个非常省时的扩展，可直接将结果赋值给结构体，其内部封装了StructScan进行转化。
	// Get用于获取单个结果然后Scan，
	// Select用来获取结果切片。
	db, err := sqlx.Open("mysql", "root:zhou123@tcp(127.0.0.1:3306)/user")

	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	Db = db
	var users []User1
	err = Db.Select(&users, "SELECT uid,username,departname,created FROM userinfo")
	if err != nil {
		fmt.Println("Select error", err)
	}
	fmt.Printf("this is Select res:%v\n", users)
	var user User1
	err1 := Db.Get(&user, "SELECT uid,username,departname,created FROM userinfo where uid = ?", 1)
	if err1 != nil {
		fmt.Println("GET error :", err1)
	} else {
		fmt.Printf("this is GET res:%v", user)
	}
}
