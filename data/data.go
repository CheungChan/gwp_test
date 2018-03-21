package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:1234@tcp(www.cheungchan.cc:3306)/gwp", )
	if err != nil {
		log.Fatal(err)
	}
	return
}
