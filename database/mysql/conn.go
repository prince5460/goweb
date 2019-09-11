package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:zhou123@tcp(127.0.0.1:3306)/user")
	if err != nil {
		fmt.Println("连接失败。。")
		return
	}

	fmt.Println("connect success ...")
	defer db.Close()
}
