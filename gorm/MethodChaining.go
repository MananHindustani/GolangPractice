package main

import (
	"fmt"		

	"github.com/sharmamanan796/gorm/Connection"
//	"gorm.io/gorm"
	"github.com/sharmamanan796/gorm/Entity"
)

func main() {

	db, err := Connection.Conn()
	var	emp []Entity.Employee
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected");

				//Chain Method 
	db.Where("first_name = ?", "Manan").Where("id = ?",2).First(&emp)
	fmt.Println(emp);
/*	
	// db is new initialized *gorm.DB, which under `New Session Mode`
	db.Where("first_name = ?", "Manan").Where("age >= ?", 24).Find(&emp)
	// `Where("name = ?", "jinzhu")` is the first method call, it will creates a new `Statement`
	// `Where("age = ?", 18)` reuse the `Statement`, and add conditions to the `Statement`
	// `Find(&users)` is a finisher, it executes registered Query Callbacks, generate and run following SQL
	// SELECT * FROM users WHERE name = 'jinzhu' AND age = 18;
	fmt.Println(emp);
	db.Where("name = ?", "Raman").Where("age >= ?", 20).Find(&emp)
	// `Where("name = ?", "jinzhu2")` is also the first method call, it creates new `Statement` too
	// `Where("age = ?", 20)` reuse the `Statement`, and add conditions to the `Statement`
	// `Find(&users)` is a finisher, it executes registered Query Callbacks, generate and run following SQL
	// SELECT * FROM users WHERE name = 'jinzhu2' AND age = 20;
	fmt.Println(emp);
	db.Find(&emp)
	// `Find(&users)` is a finisher method and also the first method call for a `New Session Mode` `*gorm.DB`
	// It creates a new `Statement` and executes registered Query Callbacks, generates and run following SQL
	// SELECT * FROM users;
	fmt.Println(emp);
*/

	// db is new initialized *gorm.DB, which under `New Session Mode`
	tx := db.Where("first_name = ?", "Manan")
	fmt.Println(emp);

	// `Where("name = ?", "jinzhu")` is the first method call, it creates a new `Statement` and add conditions
	
	tx.Where("age >= ?",20 ).Find(&emp)
	fmt.Println(emp);

	// `tx.Where("age = ?", 18)` REUSE above `Statement`, and add conditions to the `Statement`
	// `Find(&users)` is a finisher, it executes registered Query Callbacks, generate and run following SQL
	// SELECT * FROM users WHERE name = 'jinzhu' AND age = 18
	
	tx.Where("age = ?", 28).Find(&emp)
	fmt.Println(emp);

	// `tx.Where("age = ?", 18)` REUSE above `Statement` also, and add conditions to the `Statement`
	// `Find(&users)` is a finisher, it executes registered Query Callbacks, generate and run following SQL
	// SELECT * FROM users WHERE name = 'jinzhu' AND age = 18 AND age = 20;
}		