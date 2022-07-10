package customer

import (
	"github.com/sharmamanan796/account"
)

type Customer struct {
	customerID     int32
	customerName   string
	Accountdetails account.Account
}

func CustomerDetails(customerID int32, customerName string, Accountdetails account.Account) Customer {
	customer := Customer{customerID, customerName, Accountdetails}
	return customer
}
