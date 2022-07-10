package main

import (
	"fmt"

	"github.com/sharmamanan796/account"
	"github.com/sharmamanan796/atm"
	"github.com/sharmamanan796/customer"
)

func createAccount() customer.Customer {
	var accountDetails account.Account = account.AccountDetails(32154686, 7000.50)
	var customer customer.Customer = customer.CustomerDetails(1, "Manan", accountDetails)
	return customer
}

func main() {
	customer := createAccount()
	fmt.Println(customer)
	customer.Accountdetails = atm.Deposit(customer.Accountdetails, 100.0)
	fmt.Println(customer)
	customer.Accountdetails = atm.Withdraw(customer.Accountdetails, 200.0)
	fmt.Println(customer)

}
