package main

import (
	"fmt"

	"github.com/sharmamanan796/gorm/Connection"

//	"github.com/sharmamanan796/gorm/Entity"
)

func main() {

	db, err := Connection.Conn()
//	var employee Entity.Employee	
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Conneted")
/*	db.First(&employee)

	// UPDATE employees SET name='Raju', age=26 updated_at = '2013-11-17 21:34:10' WHERE id=1;
	employee.FirstName = "Raju"	
	employee.Age = 26
	employee.Email = "raju@gmail.com"
	
	db.Save(&employee)
	*/
	// Update with conditions
	
	// UPDATE employees SET sal='130000' WHERE id=3;
	//	db.Model(employee).Where("id= ?", 3).Update("sal", "1300000")

												//SELECT	
	//UPDATE employees SET last_name='ojha'  WHERE id=3;
	//	db.Model(employee).Select("last_name").Where("id= ?", 3).Updates(map[string]interface{}{"last_name": "Sharma", "age": 28})

												//OMIT
	// UPDATE employees SET age=18  WHERE id=3;											
	//	db.Model(&employee).Omit("last_name").Where("id= ?", 3).Updates(map[string]interface{}{"last_name":"ojha", "age": 28})
	
	//UPDATE "employees" SET "first_name"='Rajesh',"age"=25 where id=1
	//db.Model(&employee).Select("FirstName", "Age").Where("id= ?", 1).Updates(Entity.Employee{FirstName: "Rajesh", Age:25 })

	// UPDATE employees SET sal=100000 WHERE id IN (1, 4);	
	//db.Table("employees").Where("id IN ?", []int{1,4}).Updates(map[string]interface{}{"sal":200000})

	// UPDATE users SET name = "jinzhu"
	db.Exec("UPDATE employees SET last_name = ? where id= ?", "Ojha",3)

			
	//DB.Model(&employee).Update("sal", gorm.Expr("sal * sal + ?*sal",0.1))	
	// UPDATE "employees" SET "sal" = sal  + 0.1 * sal  WHERE "id" = 1;
	
	
	}	