package transactionusecase

import (
	"fmt"
	"liveCodeAPI/api-master/master/models"
	"liveCodeAPI/api-master/master/repositories/transactionrepository"
)

type TransactionUsecaseImpl struct {
	transactionRepo transactionrepository.TransactionRepository
}

func InitTransactionUsecase(transactionRepo transactionrepository.TransactionRepository) TransactionUsecase {
	return &TransactionUsecaseImpl{transactionRepo: transactionRepo}
}

func (t *TransactionUsecaseImpl) GetTransactions() ([]*models.Transaction, error) {
	transactions, err := t.transactionRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
func (t *TransactionUsecaseImpl) GetThisDayTransaction() ([]*models.Transaction, error) {
	transactions, err := t.transactionRepo.GetThisDay()
	if err != nil {
		return nil, err
	}
	fmt.Println(transactions)
	return transactions, nil
}
func (t *TransactionUsecaseImpl) GetMonthTransaction(data string) ([]*models.Transaction, error) {
	transactions, err := t.transactionRepo.GetByMonth(data)

	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (t *TransactionUsecaseImpl) AddTransaction(data []*models.Transaction) ([]*models.Transaction, error) {
	product, err := t.transactionRepo.Add(data)
	if err != nil {
		return nil, err
	}
	return product, nil
}
