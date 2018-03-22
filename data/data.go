package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "xxx:xxxx@tcp(www.xxx.xxx:3306)/xxx", )
	if err != nil {
		log.Fatal(err)
	}
	return
}
