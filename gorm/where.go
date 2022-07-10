package main


import (
	"fmt"

	"github.com/sharmamanan796/gorm/Connection"

	"github.com/sharmamanan796/gorm/Entity"
)
type Result struct {
	age  int
	FirstName string
}
func main() {

	db, err := Connection.Conn()
	//	var employee Entity.Employee
	//var employee []Entity.Employee
	var result Result
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Conneted")

	// SELECT * FROM employees WHERE first_name = 'Manan' ORDER BY id LIMIT 1;
	//	result := db.Where("First_Name = ?", "Manan").First(&employee)

	// SELECT * FROM employees WHERE first_name = 'Manan'
	//	result := db.Where("First_Name = ?", "Manan").Find(&employee)

	// SELECT * FROM employees WHERE first_name <> 'jinzhu';
	//result := db.Where("First_Name <> ?", "Manan").Find(&employee)

	// SELECT * FROM employees WHERE last_name LIKE '%jin%';
	//result := db.Where("Last_Name LIKE ?", "%Ve%").Find(&employee)

	//SELECT * FROM employees WHERE doj BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';
	//result := db.Where("DOJ BETWEEN ? AND ?", "2020-10-14 06:16:10.073024+00", "2020-10-17 06:16:10.073024+00").Find(&employee)
	
	db.Table("employee").Select("first_name", "age").Where("first_name = ?", "Manan").Scan(&result)
	
	// if result.Error != nil { // returns error
	// 	fmt.Println(result.Error)
	// }

//	fmt.Println(employee)
	fmt.Println(result)

}
