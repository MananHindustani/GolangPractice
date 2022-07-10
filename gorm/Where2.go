package main

import (
	"fmt"

	"github.com/sharmamanan796/gorm/Connection"

	"github.com/sharmamanan796/gorm/Entity"
)

func main() {

	db, err := Connection.Conn()
	//	var employee Entity.Employee
	var employee []Entity.Employee

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Conneted")

	// SELECT * FROM  WHERE first_name = "Manan" AND age = 24 ORDER BY id LIMIT 1
	//	result := db.Where(&Entity.Employee{FirstName: "Manan", Age: 24}).First(&employee)

	// SELECT * FROM  WHERE first_name = "Manan" AND age = 24
	//	result := db.Where(&Entity.Employee{FirstName: "Manan", Age: 24}).Find(&employee)

	// SELECT * FROM  WHERE first_name = "Manan" AND age = 24
	//	result := db.Where(map[string]interface{}{"first_name": "Manan", "age": 24}).Find(&employee)

	// SELECT * FROM  WHERE id IN (20, 21, 22);
	db.Where([]int64{0, 2, 3}).Find(&employee)

/*	if result.Error != nil { // returns error
		fmt.Println(result.Error)
	}
*/	fmt.Println(employee)
}
