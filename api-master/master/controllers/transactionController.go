package controllers

import (
	"encoding/json"
	"liveCodeAPI/api-master/master/models"
	"liveCodeAPI/api-master/master/usecases/transactionusecase"
	"liveCodeAPI/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

type TransactionsHandler struct {
	transactionUseCase transactionusecase.TransactionUsecase
}

func TransactionController(r *mux.Router, service transactionusecase.TransactionUsecase) {

	transactionHandler := TransactionsHandler{transactionUseCase: service}
	transactions := r.PathPrefix("/transactions").Subrouter()
	transactions.Use(middleware.TokenValidationMiddleware)
	transactions.HandleFunc("", transactionHandler.listTransactions).Methods(http.MethodGet)

	transactions.HandleFunc("/add", transactionHandler.addTransactions).Methods(http.MethodPost)
	transactions.HandleFunc("/month/{month}", transactionHandler.listTransactionsMonth).Methods(http.MethodGet)
	transactions.HandleFunc("/thisday", transactionHandler.listTransactionThisDay).Methods(http.MethodGet)
}

func (t *TransactionsHandler) listTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := t.transactionUseCase.GetTransactions()
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = transactions
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}
func (t *TransactionsHandler) listTransactionsMonth(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	data := params["month"]
	transactions, err := t.transactionUseCase.GetMonthTransaction(data)
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = transactions
	byteData, err := json.Marshal(response)

	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}
func (t *TransactionsHandler) listTransactionThisDay(w http.ResponseWriter, r *http.Request) {
	transactions, err := t.transactionUseCase.GetThisDayTransaction()
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = transactions
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (t *TransactionsHandler) addTransactions(w http.ResponseWriter, r *http.Request) {
	var transactionRequest []*models.Transaction
	_ = json.NewDecoder(r.Body).Decode(&transactionRequest)
	data, err := t.transactionUseCase.AddTransaction(transactionRequest)
	if err != nil {
		w.Write([]byte("Cannot Add Data"))
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = data
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}
