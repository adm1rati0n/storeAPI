package dbConnection

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func DBConnect() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/store?parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	DB = db
	fmt.Println("Кайф")
}
