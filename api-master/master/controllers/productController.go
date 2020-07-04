package controllers

import (
	"encoding/json"
	"liveCodeAPI/api-master/master/models"
	"liveCodeAPI/api-master/master/usecases/productusecase"
	"liveCodeAPI/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	productUseCase productusecase.ProductUsecase
}

func ProductController(r *mux.Router, service productusecase.ProductUsecase) {
	productHandler := ProductHandler{productUseCase: service}

	r.Use(middleware.ActivityLogMiddleware)

	product := r.PathPrefix("/products").Subrouter()
	product.HandleFunc("", productHandler.listProduct).Methods(http.MethodGet)
	product.Use(middleware.TokenValidationMiddleware)
	product.HandleFunc("/{id}", productHandler.product).Methods(http.MethodGet)
	product.HandleFunc("/add", productHandler.addProduct).Methods(http.MethodPost)
	product.HandleFunc("/delete/{id}", productHandler.deleteProduct).Methods(http.MethodDelete)
	product.HandleFunc("/update/{id}", productHandler.updateProduct).Methods(http.MethodPut)
}

func (s *ProductHandler) listProduct(w http.ResponseWriter, r *http.Request) {
	product, err := s.productUseCase.GetProducts()
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = product
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (s *ProductHandler) product(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	product, err := s.productUseCase.GetProductByID(id)
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = product
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (s *ProductHandler) addProduct(w http.ResponseWriter, r *http.Request) {
	var productRequest []*models.Product
	_ = json.NewDecoder(r.Body).Decode(&productRequest)
	_, err := s.productUseCase.AddProduct(productRequest)
	if err != nil {
		w.Write([]byte("Cannot Add Data"))
	}
	byteData, err := json.Marshal(productRequest)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (s *ProductHandler) deleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	err := s.productUseCase.DeleteProduct(id)
	if err != nil {
		w.Write([]byte("Delete Data Failed!"))
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success Deleted Data"
	byteData, err := json.Marshal(response)
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (s *ProductHandler) updateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var data models.Product
	_ = json.NewDecoder(r.Body).Decode(&data)

	product, err := s.productUseCase.UpdateProduct(id, &data)
	if err != nil {
		w.Write([]byte("Update Data Failed!"))
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = product
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}
