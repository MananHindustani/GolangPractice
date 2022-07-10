package Advance

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

	// Migrate the schema
	//db.AutoMigrate(&Entity.Employee{})
	//emp := Entity.Employee{FirstName: "Ranan", LastName: "Verma", Email: "ranan@gmail.com", Age: 27, DOJ: time.Now(), Sal: 1000000.0}
	//result := d	b.Create(&emp)

	// SELECT * FROM employees WHERE id = 1;
	//	result := db.First(&employee, 1)

	// SELECT * FROM employees WHERE id = 1;
	//	result := db.First(&employee, "2")

	// SELECT * FROM employees ORDER BY id LIMIT 1;
	//	result := db.First(&employee)

	// SELECT * FROM employees LIMIT 1;
	//result := db.Take(&employee)

	// SELECT * FROM employees ORDER BY id LIMIT 1;
	//	result := db.First(&employee)

	// SELECT * FROM employees ORDER BY id DESC LIMIT 1;
	//	result := db.Last(&employee)

	//	result.RowsAffected 	// returns found records count

	// SELECT * FROM employees WHERE id IN (1,2,3);
	//	result := db.Find(&employees, []int{1, 2})
	if result.Error != nil { // returns error
		fmt.Println(result.Error)
	}

	fmt.Println(employees)
}
