package Account

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sharmamanan796/OnlineBank/Config"
)

func CreateAccount() {
	db := Config.GetConnection()
	defer db.Close()
	var (
		owner              string
		firstName          string
		lastName           string
		balance            int
		currencyType       string
		zeroBalanceAccount bool
	)
	fmt.Println("Enter Details for opening an Account")
	fmt.Println("Enter First Name")
	fmt.Scanf("%s", &firstName)
	fmt.Println("Enter Last Name")
	fmt.Scanf("%s", &lastName)
	owner = firstName + " " + lastName
	fmt.Println("Currency Type")
	fmt.Scanf("%s\n", &currencyType)
	fmt.Println("Do you want to open with zero balance ?[true|false]")
	fmt.Scanf("%t", &zeroBalanceAccount)
	if zeroBalanceAccount == false {
		fmt.Println("Insert Balance")
		fmt.Scanf("%d", &balance)
		fmt.Print(balance)
	}
	sqlStatement := `
	INSERT INTO accounts (owner, balance, currency)
	VALUES ($1, $2, $3)
	RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, owner, balance, currencyType).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}
