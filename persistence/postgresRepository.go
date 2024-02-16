package persistence

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/casmelad/stori/pkg/accounts"
	"github.com/lib/pq"
)

const (
	SELECTACCOUNTINFO = "SELECT  a.id, a.number, u.name, u.email FROM accounts a INNER JOIN users u ON a.user_id=u.id WHERE a.number = $1"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgreRepository(dbConn *sql.DB) PostgresRepository {
	return PostgresRepository{
		db: dbConn,
	}
}

func (p PostgresRepository) GetAccountInfo(accountNumber string) (accounts.Account, error) {
	accountInfo := accounts.Account{}
	if p.db == nil {
		return accountInfo, errors.New("SQL connection object not found")
	}

	err := p.db.QueryRow(SELECTACCOUNTINFO, accountNumber).Scan(&accountInfo.Id, &accountInfo.AccountNumber, &accountInfo.User.Name, &accountInfo.User.Email)
	if err == sql.ErrNoRows {
		return accountInfo, nil
	}
	return accountInfo, err
}

func (p PostgresRepository) Save(transactions []accounts.AccountTransaction) error {
	if p.db == nil {
		return errors.New("SQL connection object not found")
	}
	txn, err := p.db.Begin()
	if err != nil {
		fmt.Println(err)
		txn.Rollback()
		return err
	}

	stm, err := txn.Prepare(pq.CopyIn("transactions", "account_id", "id", "date", "ammount"))
	if err != nil {
		fmt.Println(err)
		txn.Rollback()
		return err
	}

	for _, accTx := range transactions {
		stm.Exec(accTx.AccountId, accTx.Id, accTx.Date, accTx.Ammount)
		if err != nil {
			fmt.Println(err)
			txn.Rollback()
			return err
		}
	}

	_, err = stm.Exec()
	if err != nil {
		fmt.Println(err)
		txn.Rollback()
		return err
	}

	err = stm.Close()
	if err != nil {
		fmt.Println(err)
		txn.Rollback()
		return err
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println(err)
		txn.Rollback()
		return err
	}

	return nil
}
