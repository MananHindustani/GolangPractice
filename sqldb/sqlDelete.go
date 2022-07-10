package main

import (
	"fmt"

	"github.com/sharmamanan796/config"

	_ "github.com/lib/pq"
)

func main() {
	db := config.GetConnection()
	defer db.Close()

	sqlStatement := "delete from accounts where id=$1"
	res, err := db.Exec(sqlStatement, 3)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d row(s) deleted \n", count)
}
