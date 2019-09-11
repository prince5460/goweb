package main

//step1：导入包
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	uid        int
	username   string
	departname string
	created    string
}

/*
	查询：处理查询后的结果：
	    思路一：创建结构体
	    思路二：将数据，存入到map中
*/
func main() {

	/*
	   查询操作：
	*/
	//step2:打开数据库，建立连接
	db, _ := sql.Open("mysql", "root:zhou123@tcp(127.0.0.1:3306)/user")

	//stpt3：查询数据库
	rows, err := db.Query("SELECT uid,username,departname,created FROM userinfo")
	if err != nil {
		fmt.Println("查询有误。。")
		return
	}

	//fmt.Println(rows.Columns()) //[uid username departname created]

	//创建slice，存入struct，
	datas := make([]User, 0)

	//step4：操作结果集获取数据
	for rows.Next() {
		var uid int
		var username string
		var departname string
		var created string
		if err := rows.Scan(&uid, &username, &departname, &created); err != nil {
			fmt.Println("获取失败。。")
		}

		//每读取一行，创建一个user对象，存入datas2中
		user := User{uid, username, departname, created}
		datas = append(datas, user)
	}

	for _, v := range datas {
		fmt.Println(v)
	}

	//step5：关闭资源
	rows.Close()
	db.Close()

}
