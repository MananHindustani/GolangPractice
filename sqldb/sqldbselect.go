package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/sharmamanan796/config"

	_ "github.com/lib/pq"
)

func main() {
	db := config.GetConnection()
	defer db.Close()

	sqlStatement := ` select * from accounts where id=$1`
	var id int32
	var owner string
	var balance int64
	var currency string
	//	var createdAt
	var timeofCreation time.Time
	row := db.QueryRow(sqlStatement, 3)
	err := row.Scan(&id, &owner, &balance, &currency, &timeofCreation)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("This row doesn't exist!")
	case nil:
		fmt.Println(id, owner, balance, currency, timeofCreation)
	default:
		panic(err)
	}
}
