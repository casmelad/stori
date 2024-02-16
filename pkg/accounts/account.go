package accounts

import (
	"time"

	"github.com/casmelad/stori/pkg/users"
)

type Account struct {
	Id                  int
	User                users.User
	AccountNumber       string
	AccountTransactions []AccountTransaction //validate nil

	accountBalance float64
	avgCredit      float64
	avgDebit       float64
}

func (b Account) CalculateCurrentBalance() float64 {
	for _, tx := range b.AccountTransactions {
		b.accountBalance += tx.Ammount
	}
	return b.accountBalance
}

func (b Account) GetAvgCreditAmmount() float64 {
	avgCredit := 0.0
	crdMovmCounter := 0.0
	for _, tx := range b.AccountTransactions {
		if tx.Ammount > 0 {
			avgCredit += tx.Ammount
			crdMovmCounter++
		}
	}
	if crdMovmCounter == 0 {
		return 0
	}
	avgCredit = avgCredit / crdMovmCounter
	return avgCredit
}

func (b Account) GetAvgDebitAmmount() float64 {
	avgDebit := 0.0
	dbtMovmCounter := 0.0
	for _, tx := range b.AccountTransactions {
		if tx.Ammount < 0 {
			avgDebit += tx.Ammount
			dbtMovmCounter++
		}
	}
	if dbtMovmCounter == 0 {
		return 0
	}
	avgDebit = avgDebit / dbtMovmCounter
	return avgDebit
}

type AccountTransaction struct {
	Id        int
	AccountId int
	Date      time.Time
	Ammount   float64
}
