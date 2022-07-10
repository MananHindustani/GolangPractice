package main
import (
	"fmt"

	"github.com/sharmamanan796/gorm/Connection"

	"github.com/sharmamanan796/gorm/Entity"
)

/*type APIEmployee struct {
	ID   int
	FirstName string
}*/ 
type Pizzas struct {
  FirstName string
  Pizza string
	Size string
}

  func main() {
   db, err := Connection.Conn()
   var employees []Entity.Employee
   var pizzas Pizzas 
    if err != nil {
      panic("failed to connect database")
    }

    fmt.Println("Conneted")

    /*
    // Select `id`, `first_name` automatically when querying
    // result:=db.Model(&Entity.Employee{}).Limit(2).Find(&APIEmployee{})                  
    // SELECT `id`, `firs_name` FROM `users` LIMIT 2
  
  if result.Error != nil { // returns error
	fmt.Println(result.Error)
 
 */
 
 // SELECT * FROM "employees" WHERE sal > (SELECT AVG(sal) FROM "employees");
//db.Where("sal <  (?)", db.Table("employees").Select("AVG(sal)")).Find(&employees)


//subQuery := db.Select("AVG(age)").Where("first_name LIKE ?", "%an").Table("employees")
//db.Select("first_name,AVG(age) as avgage").Group("first_name").Having("AVG(age) > (?)", subQuery).Find(&employees)

subQuery1 := db.Model(&Entity.Employee{}).Select("first_name")
subQuery2 := db.Model(&Pizzas{}).Select("first_name")

result:=db.Table("(?) as u, (?) as p", subQuery1, subQuery2).Find(&Entity.Employee{})
// SELECT * FROM (SELECT `first_name` FROM `users`) as u, (SELECT `first_name` FROM `pets`) as p
fmt.Println(employees)
fmt.Println(pizzas)
fmt.Println(result)

}
