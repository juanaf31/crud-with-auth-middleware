package transactionrepository

import "liveCodeAPI/api-master/master/models"

type TransactionRepository interface {
	GetAll() ([]*models.Transaction, error)
	GetThisDay() ([]*models.Transaction, error)
	GetByMonth(string) ([]*models.Transaction, error)
	Add([]*models.Transaction) ([]*models.Transaction, error)
}
