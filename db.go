package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func connect() {
	db, err := sql.Open("mysql", "user:pass@(api_sample_db_1:3306)/test")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("データベース接続失敗")
	} else {
		fmt.Println("データベース接続成功")
	}

}
