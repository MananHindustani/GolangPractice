package Account

import (
	"github.com/sharmamanan796/onlineBank/Config"
	"github.com/sharmamanan796/onlineBank/Entity"

	_ "github.com/lib/pq"
)

func UpdateAccount(account *Entity.Account, id int) (int64, error) {
	db, err := Config.GetConnection()

	defer db.Close()
	var (
		firstName    string = account.FirstName
		lastName     string = account.LastName
		balance      int64  = account.Balance
		currencyType string = account.CurrencyType
	)
	account.Owner = firstName + " " + lastName
	owner := account.Owner

	sqlStatement := " UPDATE accounts set owner=$2,balance=$3,currency=$4 where id=$1"
	res, err := db.Exec(sqlStatement, id, owner, balance, currencyType)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)

	}
	return count, err
}
