package main

import (
	"database/sql"
	"fmt"
)

const (
	host     = "192.168.1.19"
	port     = 5432
	user     = "admin"
	password = "adminadmin"
	dbname   = "admin"
)

func dbConnect() {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
