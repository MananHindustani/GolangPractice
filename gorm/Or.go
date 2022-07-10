package main

import (
	"fmt"

	"github.com/sharmamanan796/gorm/Connection"

	"github.com/sharmamanan796/gorm/Entity"
)

func main() {

	db, err := Connection.Conn()
	var employees []Entity.Employee
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Conneted")
	// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';
	//	result := db.Where("age = ?", "24").Or("sal =40000 ").Find(&employees)

	// SELECT * FROM users WHERE first_name = 'Manan' OR (name = 'Manan2' AND age >=20 );
	result := db.Where("first_name = 'Manan'").Or(map[string]interface{}{"first_name": "Raman"}).Find(&employees)

	if result.Error != nil { // returns error
		fmt.Println(result.Error)
	}
	fmt.Println(employees)

	// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';
	//	db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)

}
