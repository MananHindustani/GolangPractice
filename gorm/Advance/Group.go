package main
import (
	"fmt"

	"github.com/sharmamanan796/gorm/Connection"

)

type Pizzas struct {
	Pizza string
	Size string
}

func main() {

db, err := Connection.Conn()
if err != nil {
      panic("failed to connect database")
}

var pizzas Pizzas
		
// SELECT * FROM `pizzas` WHERE (pizza = "pepperoni" AND (size = "small" OR size = "medium")) OR (pizza = "hawaiian" AND size = "xlarge")
/*result:=db.Where(
  db.Where("pizza = ?", "pepperoni").Where(db.Where("size = ?", "small").Or("size = ?", "medium")),
).Or(
  db.Where("pizza = ?", "hawaiian").Where("size = ?", "xlarge"),
).Find(&Pizzas{}).Statement
*/
db.Where(("pizza = ?", "hawaiian").Where("size = ?", "xlarge")).Find(&Pizzas{})

fmt.Println(result)

fmt.Println(pizzas)

}
