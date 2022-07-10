package Account

import (
	_ "github.com/lib/pq"
	"github.com/sharmamanan796/onlineBank/Config"
	"github.com/sharmamanan796/onlineBank/Entity"
)

func CreateAccount(account *Entity.Account) error {
	db, err := Config.GetConnection()
	if err != nil {
		panic(err)
		return err
	}
	defer db.Close()
	var (
		firstName    string = account.FirstName
		lastName     string = account.LastName
		balance      int64  = account.Balance
		currencyType string = account.CurrencyType
		id           int64
	)

	account.Owner = firstName + " " + lastName
	owner := account.Owner
	sqlStatement := `
	INSERT INTO accounts (owner, balance, currencytype)
	VALUES ($1, $2, $3)
	RETURNING id`
	err = db.QueryRow(sqlStatement, owner, balance, currencyType).Scan(&id)
	if err != nil {
		panic(err)
	}
	account.Owner = owner
	account.ID = id
	return err
}
