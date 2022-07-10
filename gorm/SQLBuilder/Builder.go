package main

import (
	"fmt"

	"github.com/sharmamanan796/gorm/Connection"
	"github.com/sharmamanan796/gorm/Entity"
)

func main() { 

	db, err := Connection.Conn()
	var employee Entity.Employee
//	var employees []Entity.Employee
	var avg int
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Conneted")
  
  
 //  db.Raw("SELECT id, first_name,last_name, age FROM employees WHERE id = ?", 3).Scan(&employee)
  
// 	 db.Raw("SELECT id, first_name,last_name, age FROM employees WHERE first_name = ?", "Manan").Scan(&employee)
  
//  var age intfmt.Println(employee)

// db.Raw("select sum(age) as sum from employees where sal>=?", 100000).Scan(&sum) 
//fmt.Println(sum)

//	db.Exec("DROP TABLE pizzas")

//db.Exec("update users set money=? where name = ?", gorm.Expr("money * ? + ?", 10000, 1), "jinzhu")

//db.Where("first_name = @name OR first_name = @name", sql.Named("name", "Manan")).Find(&employee)

//db.Where("first_name = @name OR first_name = @name2", map[string]interface{}{"name": "Manan"}).First(&employee)

// SELECT * FROM users WHERE name1 = "jinzhu1" OR name2 = "jinzhu2" OR name3 = "jinzhu1"
//db.Raw("SELECT * FROM users WHERE  first_name= @name OR last_name = @name",sql.Named("name", "jinzhu1"), sql.Named("name2", "jinzhu2")).Find(&user)

//db.Select("AVG(age) as avg").Group("first_name").Having("AVG	(age) > (?)", subQuery).Find(&employee)

//fmt.Println(employee)

fmt.Println(avg)

}