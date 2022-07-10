package Account

import "github.com/sharmamanan796/onlineBank/Config"

func DeleteAccount(id int) (int64, error) {
	db, err := Config.GetConnection()
	defer db.Close()

	sqlStatement := "delete from accounts where id=$1"
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	return count, err
}
