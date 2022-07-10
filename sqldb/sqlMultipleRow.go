package main

import (
	"fmt"
	"time"

	"github.com/sharmamanan796/config"

	_ "github.com/lib/pq"
)

func main() {
	db := config.GetConnection()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM accounts")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	fmt.Println("ID\tOwner\t\tbalance\tCurrency\tTime_of_Creation")
	for rows.Next() {
		var id int32
		var owner string
		var balance int64
		var currency string
		var timeofCreation time.Time
		err := rows.Scan(&id, &owner, &balance, &currency, &timeofCreation)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d\t%s\t%d\t%s\t%s\t\n", id, owner, balance, currency, timeofCreation)
	}

}
