package Entity

import (
	"time"	
	"gorm.io/gorm"
)

type Employee struct {
	ID        uint      `gorm:"primary_Key;auto_increment"`
	FirstName string    `gorm:"type:varchar(100)"`
	LastName  string    `gorm:"type:varchar(100)"`
	Email     string    `gorm:"type:varchar(200)"`
	Age       int       `gorm:"type:int"`
	DOJ       time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Sal       float64   `gorm:"type:numeric(10,2)"`
	Team string  `gorm:"type:varchar(100)"`
}

				//Belong
// 	gorm.Model
// 	Name      string  `gorm:"type:varchar(100)"`
// 	CompanyRefer int
// 	//CompanyID int `gorm:"foreignKey:CompanyRefer"`  
// 	Company   Company `gorm:"foreignKey:CompanyRefer"`
//   }
  
//   type Company struct {
// 	ID   int  `gorm:"primary_Key;auto_increment"`
// 	Name string `gorm:"type:varchar(100)"`
//   }
// User has one CreditCard, CreditCardID is the foreign key

						  //HAS-ONE
// 					//Override References
// type User struct {
// 	gorm.Model
// 	CreditCard CreditCard `gorm:"foreignKey:UserName"`
// 	// use UserName as foreign key
//   }
  
//   type CreditCard struct {
// 	gorm.Model
// 	Number   string
// 	UserName string
//   }


				  
type Coder struct {
	gorm.Model
	Languages []Language `gorm:"many2many:coders_languages;"`
  }
  
  type Language struct {
	gorm.Model
	Name string
	Coder []*Coder `gorm:"many2many:coders_languages;"`
  }				  
					  