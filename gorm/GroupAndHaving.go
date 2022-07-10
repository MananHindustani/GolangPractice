package main

import (
	"fmt"

	"github.com/sharmamanan796/gorm/Connection"

	"github.com/sharmamanan796/gorm/Entity"
)

 type Result struct {
	Team string  
	Total float64

}
  

func main() {

	var result []Result
	db, err := Connection.Conn()
	var employees []Entity.Employee
	if err != nil {
		panic("failed to connect database")
	}
	// SELECT sum(age) as total FROM `employees` ;
	//	db.Model(&Entity.Employee{}).Select("sum(age) as total").Find(&total)
	
	//SELECT team, sum(age) as total FROM employees GROUP BY "team"	
	//	db.Model(&Entity.Employee{}).Select("team, avg(sal) as total").Group("team").Find(&result) 
								
	//SELECT team, sum(age) as total FROM employees GROUP BY "team having team='Frontend'"	
	//db.Model(&Entity.Employee{}).Select("team, avg(sal) as total").Group("team").Having("team = ?", "Frontend").Find(&result) 

	

	fmt.Println(result)

	fmt.Println(employees)
	
}
	