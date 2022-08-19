package repository

import (
	"database/sql"
	"errors"

	"github.com/AntonioCarlos850/bank-go/domain"
)

type TransactionRepositoryDb struct {
	Db *sql.DB
}

func NewTransactionRepositoryDb(db *sql.DB) TransactionRepositoryDb {
	return TransactionRepositoryDb{db}
}

func (tr TransactionRepositoryDb) GetCreditCard(c domain.CreditCard) (domain.CreditCard, error) {
	var creditCard domain.CreditCard
	stmt, err := tr.Db.Prepare("select id, balance, limit from credit_card where number=$1")
	if err != nil {
		return creditCard, nil
	}
	if err = stmt.QueryRow(c.Number).Scan(&creditCard.ID, &creditCard.Balance, &creditCard.Limit); err != nil {
		return creditCard, errors.New("credit card doesn't exists")
	}

	return creditCard, nil
}

func (tr TransactionRepositoryDb) CreateCreditCard(c domain.CreditCard) error {
	stmt, err := tr.Db.Prepare("insert into credit_card (id, name, number, expiration_month, expiration_year, cvv, balance, limit) values ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(c.ID, c.Name, c.Number, c.ExpirationMonth, c.ExpirationYear, c.CVV, c.Balance, c.Limit)
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

func (tr TransactionRepositoryDb) SaveTransaction(c domain.CreditCard, t domain.Transaction) error {
	stmt, err := tr.Db.Prepare("insert into transaction ('id', 'amount', 'status', 'description', 'store', 'credit_card_id') values ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(t.ID, t.Amount, t.Status, t.Description, t.Store, t.CreditCardId)
	if err != nil {
		return err
	}

	if t.Status == "Approved" {
		err = tr.UpdateBalance(c)
		if err != nil {
			return err
		}
	}

	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

func (tr TransactionRepositoryDb) UpdateBalance(c domain.CreditCard) error {
	stmt, err := tr.Db.Prepare("update into credit_card (balance) values ($1)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(c.Balance)
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil
}
