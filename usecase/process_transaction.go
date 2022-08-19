package usecase

import (
	"github.com/AntonioCarlos850/bank-go/domain"
	"github.com/AntonioCarlos850/bank-go/dto"
)

type TransactionUseCase struct {
	TransactionRepository domain.TransactionRepository
}

func NewTransactionUseCase(tr domain.TransactionRepository) TransactionUseCase {
	return TransactionUseCase{tr}
}

func (tuc TransactionUseCase) ProcessTransaction(transactionDto dto.TransactionDto) (domain.Transaction, error) {
	cc := domain.NewCreditCard()
	cc.Number = transactionDto.Number

	ccBalanceAndLimit, err := tuc.TransactionRepository.GetCreditCard(*cc)
	if err != nil {
		return domain.Transaction{}, err
	}

	t := tuc.newTransaction(transactionDto, ccBalanceAndLimit)
	t.ProcessAndValidate(&ccBalanceAndLimit)

	err = tuc.TransactionRepository.SaveTransaction(ccBalanceAndLimit, *t)
	if err != nil {
		return domain.Transaction{}, err
	}

	return *t, nil
}

func (tuc TransactionUseCase) newTransaction(transaction dto.TransactionDto, cc domain.CreditCard) *domain.Transaction {
	t := domain.NewTransaction()
	t.Amount = transaction.Amount
	t.CreditCardId = cc.ID
	t.Description = transaction.Description
	t.Store = transaction.Store
	return t
}
