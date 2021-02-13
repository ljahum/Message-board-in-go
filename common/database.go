package common

import (
	"GoCode/model"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Main_domain string = "localhost"

var DB *sql.DB

func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "root:123321@tcp(127.0.0.1:3306)/database?charset=utf8")
	if err != nil { // 连接失败
		fmt.Printf("connect mysql fail ! [%s]", err)
	} else { // 连接成功
		fmt.Println("connect to mysql success")
	}
	sqlStr := "select name, passwd ,mail from user_tab where id=?"
	rows, err := db.Query(sqlStr, 1)
	if err != nil {
		panic("fail to connect databse,err:")
	}
	defer rows.Close()
	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.Name, &u.Passwd, &u.Mail)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
		}
		fmt.Printf("name:%s passwd:%s mail:%s\n", u.Name, u.Passwd, u.Mail)
	}

	DB = db
	return db
}

func GetDB() *sql.DB {
	return DB
}
