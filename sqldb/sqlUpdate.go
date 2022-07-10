package main

import (
	"fmt"

	"github.com/sharmamanan796/config"

	_ "github.com/lib/pq"
)

func main() {
	db := config.GetConnection()
	defer db.Close()
	sqlStatement := " UPDATE accounts set balance=$2 where id=$1"
	res, err := db.Exec(sqlStatement, 1, 20000)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d row(s) updated \n", count)

}
