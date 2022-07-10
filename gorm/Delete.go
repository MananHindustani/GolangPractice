package main

import (

	"github.com/sharmamanan796/gorm/Connection"

 //   "github.com/sharmamanan796/gorm/Entity"
)
type Pizzas struct {
	FirstName string
	Pizza string 
	  Size string
  }
  
func main() {		

	db, err := Connection.Conn()
//	var employee Entity.Employee	
	var pizzas Pizzas
	if err != nil {
		panic("failed to connect database")
	}
//	fmt.Println("Conneted")
// db.Where("size LIKE ?", "%ll").Delete(pizzas)
	
//db.Delete(employee,1)	

//db.Delete(pizzas, "pizza=?","pepperoni")
//db.Exec("DELETE FROM pizzas")

db.Where("pizza = ?", "pepperoni").Delete(pizzas)
}	
