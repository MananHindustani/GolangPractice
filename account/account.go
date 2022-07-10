package account

type Account struct {
	accountNo int64

	AccountBalance float64
}

func AccountDetails(accountNo int64, accountBalance float64) Account {
	accountDetails := Account{accountNo, accountBalance}
	return accountDetails
}
