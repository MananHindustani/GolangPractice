package Config

import (
	"os"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)


func GetConnection() (*sql.DB, error) {
	host     := "db"
	port     := 5432
	user     := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname   := os.Getenv("POSTGRES_DB")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db, err
}
