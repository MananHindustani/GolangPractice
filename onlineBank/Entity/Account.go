package Entity

import "time"

type Account struct {
	ID             int64     `json:"id"`
	FirstName      string    `json:"firstName" validate:"required,alpha,omitempty"`
	LastName       string    `json:"lastName" validate:"required,omitempty"`
	Owner          string    `json:"owner"`
	Balance        int64     `json:"balance" validate:"required"`
	CurrencyType   string    `json:"currencyType" validate:"required"`
	TimeofCreation time.Time `json:"timeOfCreation"`
}
