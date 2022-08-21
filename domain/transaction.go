package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type TransactionRepository interface {
	GetCreditCard(c CreditCard) (CreditCard, error)
	CreateCreditCard(c CreditCard) error
	SaveTransaction(c CreditCard, t Transaction) error
}

type Transaction struct {
	ID           string
	Amount       float64
	Status       string
	Description  string
	Store        string
	CreditCardId string
	CreatedAt    time.Time
}

func NewTransaction() *Transaction {
	t := new(Transaction)
	t.CreatedAt = time.Now()
	t.ID = uuid.NewV4().String()
	return t
}

func (t *Transaction) ProcessAndValidate(c *CreditCard) {
	if t.Amount+c.Balance > c.Limit {
		t.Status = "Rejected"
	} else {
		t.Status = "Approved"
		c.Balance = t.Amount + c.Balance
	}
}
