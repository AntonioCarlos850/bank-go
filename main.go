package main

import (
	"database/sql"

	"github.com/AntonioCarlos850/bank-go/domain"
	"github.com/AntonioCarlos850/bank-go/dto"
	"github.com/AntonioCarlos850/bank-go/infrastructure/repository"
	"github.com/AntonioCarlos850/bank-go/usecase"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

var tr repository.TransactionRepositoryDb

func main() {
	e := echo.New()
	db, err := DbConnect()
	if err != nil {
		e.Server.ErrorLog.Fatalf("db conection refused")
		e.Close()
	}
	tr = repository.NewTransactionRepositoryDb(db)
	e.POST("/bank-accounts", createBankAccount)
	//e.POST("/bank-accounts/transfer", transfer)
	e.Logger.Fatal(e.Start(":8000"))
}

func DbConnect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "db/database.db")
	if err != nil {
		return &sql.DB{}, err
	}
	return db, nil
}

func createBankAccount(c echo.Context) error {
	defer tr.Db.Close()

	dto := new(dto.CreditCardDto)
	if err := c.Bind(dto); err != nil {
		return err
	}

	cc := hidrateCreditCard(dto)
	useCase := usecase.NewCreditCardUseCase(tr)

	err := useCase.NewCreditCard(cc)
	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(201, cc)
}

func hidrateCreditCard(ccDto *dto.CreditCardDto) domain.CreditCard {
	cc := domain.NewCreditCard()
	cc.Name = "Antonio"
	cc.Number = ccDto.Number
	cc.ExpirationMonth = 12
	cc.ExpirationYear = 25
	cc.CVV = 155
	cc.Balance = 0
	cc.Limit = 1000

	return *cc
}

func transfer(c echo.Context) error {
	return nil
}
