package atm

import "github.com/sharmamanan796/account"

func Deposit(account account.Account, amount float64) account.Account {
	account.AccountBalance = account.AccountBalance + amount
	return account
}

func Withdraw(account account.Account, amount float64) account.Account {
	account.AccountBalance = account.AccountBalance - amount
	return account
}
