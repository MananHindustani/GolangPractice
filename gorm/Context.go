package main

import (
	"fmt"		
	"context"
	"github.com/sharmamanan796/gorm/Connection"
//	"gorm.io/gorm"
	"github.com/sharmamanan796/gorm/Entity"
)

func main() {
	var emp []Entity.Employee
	db, err := Connection.Conn()
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Conneted")
	var ctx  context.Context
	tx := db.WithContext(ctx)
	tx.First(&emp, 1)
	fmt.Println(emp) 


}	