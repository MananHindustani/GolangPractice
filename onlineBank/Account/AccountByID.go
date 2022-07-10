package Account

import (
	"strconv"
	"time"

	"github.com/sharmamanan796/onlineBank/Config"
	"github.com/sharmamanan796/onlineBank/Entity"

	_ "github.com/lib/pq"
)

func AccountById(Id int) (Entity.Account, error, string) {
	db, err := Config.GetConnection()
	defer db.Close()
	msg := "row number" + strconv.Itoa(Id)
	sqlStatement := ` select * from accounts where id=$1`
	var id int64
	var owner string
	var balance int64
	var currencyType string
	var timeofCreation time.Time
	row := db.QueryRow(sqlStatement, Id)
	err = row.Scan(&id, &owner, &balance, &currencyType)
	if err != nil {
		msg = "This row doesn't exist!"
	}
	account := Entity.Account{
		ID:             id,
		Owner:          owner,
		Balance:        balance,
		CurrencyType:   currencyType,
		TimeofCreation: timeofCreation,
	}

	return account, err, msg
}
