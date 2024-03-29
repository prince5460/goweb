package main

// step1：导入包
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//CREATE TABLE `userinfo` (
	//	`uid` INT(10) NOT NULL AUTO_INCREMENT,
	//	`username` VARCHAR(64) NULL DEFAULT NULL,
	//	`departname` VARCHAR(64) NULL DEFAULT NULL,
	//	`created` DATE NULL DEFAULT NULL,
	//	PRIMARY KEY (`uid`)
	//)
	//CREATE TABLE `userdetail` (
	//	`uid` INT(10) NOT NULL DEFAULT '0',
	//	`intro` TEXT NULL,
	//	`profile` TEXT NULL,
	//	PRIMARY KEY (`uid`)
	//)

	// step2：打开数据库，相当于和数据库建立连接：db对象
	/*
	   func Open(driverName, dataSourceName string) (*DB, error)
	   drvierName,"mysql"
	   dataSourceName,用户名:密码@协议(地址:端口)/数据库?参数=参数值
	*/
	db, err := sql.Open("mysql", "root:zhou123@tcp(127.0.0.1:3306)/user")
	if err != nil {
		fmt.Println("连接失败。。")
		return
	}

	//step3：插入一条数据

	stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) values(?,?,?)")
	if err != nil {
		fmt.Println("操作失败。。")
	}
	//补充完整sql语句，并执行
	result, err := stmt.Exec("zhou", "技术部", "2018-11-21")
	if err != nil {
		fmt.Println("插入数据失败。。")
	}
	//step4：处理sql操作后的结果
	lastInsertId, err := result.LastInsertId()
	rowsAffected, err := result.RowsAffected()
	fmt.Println("lastInsertId", lastInsertId)
	fmt.Println("影响的行数：", rowsAffected)

	//再次插入数据：
	result, _ = stmt.Exec("ruby", "人事部", "2018-11-11")
	count, _ := result.RowsAffected()
	fmt.Println("影响的行数：", count)

	//step5：关闭资源
	stmt.Close()
	db.Close()

}
