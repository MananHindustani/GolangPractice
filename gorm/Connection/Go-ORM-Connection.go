package Connection

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	
)

func Conn() (*gorm.DB, error) {
	dsn := "host=172.17.0.2 user=postgres password=postgres dbname=company port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
