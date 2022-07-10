package main

import (
	"fmt"

	"github.com/sharmamanan796/config"

	_ "github.com/lib/pq"
)

const (
	host     = "172.17.0.2"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "bank"
)

func main() {
	db := config.GetConnection()
	defer db.Close()

	/*	sqlStatement := `
		INSERT INTO accounts (owner, balance, currency)
		VALUES ('Manan Sharma', 100011, '700')`
		_, err = db.Exec(sqlStatement)
	*/
	sqlStatement := `
	INSERT INTO accounts (owner, balance, currency)
	VALUES ($1, $2, $3)
	RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, "Rajesh SharmA", 20000, "dollor").Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println("New record ID is:", id)
}
