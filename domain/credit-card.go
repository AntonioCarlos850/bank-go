package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type CreditCard struct {
	ID              string    `json:"id"`
	Name            string    `json:"-"`
	Number          string    `json:"account_number"`
	ExpirationMonth int32     `json:"-"`
	ExpirationYear  int32     `json:"-"`
	CVV             int32     `json:"-"`
	Balance         float64   `json:"-"`
	Limit           float64   `json:"-"`
	CreatedAt       time.Time `json:"-"`
}

func NewCreditCard() *CreditCard {
	c := new(CreditCard)
	c.CreatedAt = time.Now()
	c.ID = uuid.NewV4().String()
	return c
}
