package codecs

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/casmelad/stori/cmd/dto"
	"github.com/casmelad/stori/pkg/accounts"
)

func MapTransactionsFromText(accountMovements []string) ([]accounts.AccountTransaction, error) {

	trans := []accounts.AccountTransaction{}
	for _, record := range accountMovements {
		tx := mapTransactionFromString(record)
		trans = append(trans, tx)
	}
	return trans, nil
}

func mapTransactionFromString(txString string) accounts.AccountTransaction {
	record := strings.Split(txString, ",")

	id, err := strconv.Atoi(record[0])
	date, err := time.Parse("2006/01/02", fmt.Sprintf("%d/%s", time.Now().Year(), fmt.Sprintf("%05s", record[1])))
	ammount, err := strconv.ParseFloat(strings.TrimSpace(record[2]), 64)

	if err != nil {
		return accounts.AccountTransaction{}
	}
	tx := accounts.AccountTransaction{
		Id:      id,
		Date:    date,
		Ammount: ammount,
	}

	return tx
}

func MapAccountToAccountBalanceNotificationDto(ab accounts.Account) dto.AccountBalanceNotificationDto {
	dto := &dto.AccountBalanceNotificationDto{}
	dto.Balance = ab.CalculateCurrentBalance()
	dto.AvgCreditAmmount = ab.GetAvgCreditAmmount()
	dto.AvgDebitAmmount = ab.GetAvgDebitAmmount()
	dto.To = ab.User.Email
	dto = filterTransactionsByMonth(dto, ab.AccountTransactions)

	return *dto
}

func filterTransactionsByMonth(accDto *dto.AccountBalanceNotificationDto, ts []accounts.AccountTransaction) *dto.AccountBalanceNotificationDto {

	for _, tx := range ts {
		switch tx.Date.Month() {
		case time.January:
			accDto.JanuaryTransactions++
		case time.February:
			accDto.FebruaryTransactions++
		case time.March:
			accDto.MarchTransactions++
		case time.April:
			accDto.AprilTransactions++
		case time.May:
			accDto.MayTransactions++
		case time.June:
			accDto.JuneTransactions++
		case time.July:
			accDto.JulyTransactions++
		case time.August:
			accDto.AugustTransactions++
		case time.September:
			accDto.SeptemberTransactions++
		case time.October:
			accDto.OctoberTransactions++
		case time.November:
			accDto.NovemberTransactions++
		case time.December:
			accDto.DecemberTransactions++
		}
	}

	return accDto
}
