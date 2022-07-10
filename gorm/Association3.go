package main

import (
	"fmt"

	"github.com/sharmamanan796/gorm/Connection"
//	"gorm.io/gorm"
	//"github.com/sharmamanan796/gorm/Entity"
)			
type Emails struct{
	Id int  `gorm:"primary_Key;auto_increment"`
	Email string
	}
type Languages struct{
	Id int   `gorm:"primary_Key;auto_increment"`
	Name string
	} 
type Address struct{
 	Id int  `gorm:"primary_Key;auto_increment"`
 	Address1 string
 	}
type Candidate struct{ 
	Id int  `gorm:"primary_Key;auto_increment"`
	Name string
	BillingAddress Address
	ShippingAddress Address 
	Email  []Emails 
	Language []Languages
}
func main() {
	db, err := Connection.Conn()
	if err != nil {
		panic("failed to connect database")
	} 
	db.Create(Emails{})  
	db.Create(Languages{})  
	db.Create(Address{})



	fmt.Println("Conneted")
	candidate := Candidate{
		Name:            "jinzhu",
	BillingAddress:  Address{Address1: "Billing Address - Address 1"},
	ShippingAddress: Address{Address1: "Shipping Address - Address 1"},
	Email:Emails{
		  {Email: "manuu@example.com"},
		  {Email: "manu-2@example.com"},
		},
		Language:[]Languages{	
		  {Name: "HI"},
		  {Name: "EN"},
		},
	  }

	  db.Create(&candidate)
	  db.Save(&candidate)
	// BEGIN TRANSACTION;
// INSERT INTO "addresses" (address1) VALUES ("Billing Address - Address 1"), ("Shipping Address - Address 1") ON DUPLICATE KEY DO NOTHING;
// INSERT INTO "users" (name,billing_address_id,shipping_address_id) VALUES ("jinzhu", 1, 2);
// INSERT INTO "emails" (user_id,email) VALUES (111, "jinzhu@example.com"), (111, "jinzhu-2@example.com") ON DUPLICATE KEY DO NOTHING;
// INSERT INTO "languages" ("name") VALUES ('ZH'), ('EN') ON DUPLICATE KEY DO NOTHING;
// INSERT INTO "user_languages" ("user_id","language_id") VALUES (111, 1), (111, 2) ON DUPLICATE KEY DO NOTHING;
// COMMIT;
	
}