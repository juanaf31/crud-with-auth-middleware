package transactionusecase

import "liveCodeAPI/api-master/master/models"

type TransactionUsecase interface {
	GetTransactions() ([]*models.Transaction, error)
	GetThisDayTransaction() ([]*models.Transaction, error)
	GetMonthTransaction(string) ([]*models.Transaction, error)
	AddTransaction([]*models.Transaction) ([]*models.Transaction, error)
}
