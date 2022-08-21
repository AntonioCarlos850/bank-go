package usecase

import "github.com/AntonioCarlos850/bank-go/domain"

type CreditCardUseCase struct {
	tr domain.TransactionRepository
}

func NewCreditCardUseCase(tr domain.TransactionRepository) CreditCardUseCase {
	return CreditCardUseCase{tr}
}

func (ccUseCase CreditCardUseCase) NewCreditCard(cc domain.CreditCard) error {
	err := ccUseCase.tr.CreateCreditCard(cc)
	if err != nil {
		return err
	}

	return nil
}
