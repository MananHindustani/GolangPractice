//inline works similar to work

package main

import (
	"fmt"

	"github.com/sharmamanan796/gorm/Connection"

	"github.com/sharmamanan796/gorm/Entity"
)

func main() {

	db, err := Connection.Conn()
	//	var employee Entity.Employee
	var employees []Entity.Employee

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Conneted")

	//	db.First(&employee, "id = ?", "3")
	//	fmt.Println(employee)

	// SELECT * FROM employees WHERE first_name = "Manan";
	// result := db.Find(&employee, "First_Name = ?", "Manan")

	// SELECT * FROM employees WHERE first_name <> "Manan" AND age > 24;
	//	result := db.Find(&employees, "First_Name <> ? AND age <> ?", "Manan", 26)

	// SELECT * FROM users WHERE age = 20;
	result := db.Find(&users, User{Age: 20})

	if result.Error != nil { // returns error
		fmt.Println(result.Error)
	}

	fmt.Println(employees)
}
