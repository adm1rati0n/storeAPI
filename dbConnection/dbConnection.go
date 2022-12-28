package dbConnection

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func DBConnect() {

	//cfg := mysql.Config{
	//	User:                 "root",
	//	Passwd:               "",
	//	Net:                  "tcp",
	//	Addr:                 "127.0.0.1:3306",
	//	DBName:               "store",
	//	AllowNativePasswords: true,
	//}

	cfg := mysql.Config{
		User:                 "u1833241_evgen",
		Passwd:               "2wJ-drp-Mt5-DTd",
		Net:                  "tcp",
		Addr:                 "31.31.196.229:3306",
		DBName:               "u1833241_storedb",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
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
