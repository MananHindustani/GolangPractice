package Account

import (
	"time"

	_ "github.com/lib/pq"
	"github.com/sharmamanan796/onlineBank/Config"
	"github.com/sharmamanan796/onlineBank/Entity"
)

func ListAccount(account []Entity.Account) ([]Entity.Account, error) {
	db, err := Config.GetConnection()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM accounts")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	i := 0

	for rows.Next() {
		var id int64
		var owner string
		var balance int64
		var currencyType string
		var timeofCreation time.Time
		err := rows.Scan(&id, &owner, &balance, &currencyType, &timeofCreation)
		if err != nil {
			panic(err)
		}

		accountCount := Entity.Account{
			ID:             id,
			Owner:          owner,
			Balance:        balance,
			CurrencyType:   currencyType,
			TimeofCreation: timeofCreation,
		}
		account = append(account, accountCount)
		i++
	}
	return account, err
}
