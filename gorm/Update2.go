package main

import (
	"fmt"

	"github.com/sharmamanan796/gorm/Connection"

	"github.com/sharmamanan796/gorm/Entity"
)

type Pizzas struct {
	FirstName string
	Pizza string
	  Size string
  }
  
func main() {

	db, err := Connection.Conn()
//	var employee Entity.Employee	
	
	var pizza pizzas
//	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Conneted")


	
//	db.Model(&Entity.Employee).Update("piza_size", db.Model(&Pizzas).Select("name").Where("companies.id = users.company_id"))
// UPDATE "employee" SET "company_name" = (SELECT name FROM companies WHERE companies.id = users.company_id);

db.Table("employee as e").Where("first_name = ?", "").Update("company_name", db.Table("companies as c").Select("name").Where("c.id = u.company_id"))

}
