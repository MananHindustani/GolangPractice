package main

import (
	"fmt"

	"github.com/sharmamanan796/gorm/Connection"
//	"gorm.io/gorm"
//	"github.com/sharmamanan796/gorm/Entity"
)

type Toy struct {
	ID        int
	Name      string
	OwnerID   int
	OwnerType string
}
					//HAS ONE
type Cat struct {
ID    int
Name  string
Toy   Toy `gorm:"polymorphic:Owner;"`
}	
											  
type Dog struct {
ID   int
Name string
Toy  Toy `gorm:"polymorphic:Owner;polymorphicValue:master"`
}
											  
							
func main() {

	db, err := Connection.Conn()
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Conneted")
// `User` belongs to `Company`, `CompanyID` is the foreign key


  	// // Migrate the schema
	// db.AutoMigrate(&Entity.User{}) 
	// db.AutoMigrate(&Entity.CreditCard{}) 

	//	db.AutoMigrate(&Entity.Company{})

	db.AutoMigrate(Toy{})
	db.AutoMigrate(Dog{})
	db.AutoMigrate(Cat{})

	db.Create(&Dog{Name: "dog1",Toy:Toy{Name: "toy1"}})
	db.Create(&Cat{Name: "cat1",Toy:Toy{Name: "toy1"}})

	// INSERT INTO `dogs` (`name`) VALUES ("dog1")
	// INSERT INTO `toys` (`name`,`owner_id`,`owner_type`) VALUES ("toy1","1","dogs")
	fmt.Println("Table created")	

}	