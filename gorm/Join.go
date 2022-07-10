
package main

import (
	"fmt"

	"github.com/sharmamanan796/gorm/Connection"
	"github.com/sharmamanan796/gorm/Entity"
)

type Result struct {
	Id  int
	FirstName string
}

func main() {
	db, err := Connection.Conn()
	//	var employee Entity.Employee
	var employees []Entity.Employee
//	var result []Result
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Conneted")

	// SELECT departmrnt.id, employees.first_name FROM employees left join emails on  employees.team = departmrnt.department_name
	//	db.Model(&employees).Select("departmrnt.id, employees.first_name").Joins("left join departmrnt on employees.team=departmrnt.department_name ").Scan(&result)
	
	//rows, err := db.Table("employees").Select("employees.first_name, departmrnt.id").Joins("left join departmrnt on departmrnt.department_name = employees.team").Rows()
	
	// defer rows.Close()
	// for rows.Next() {
	// 	err := rows.Scan(&result.Id, &result.FirstName)
	// 	if err != nil {
	// 		//log.Fatal(err)
	// 	}	
	// 	fmt.Println(result.Id," ",result.FirstName)
	// }
					//work with ids	
//	db.Joins("departmrnt").Find(&employees)	
	fmt.Println(employees)		

}
