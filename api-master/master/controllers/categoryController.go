package controllers

import (
	"encoding/json"
	"liveCodeAPI/api-master/master/models"
	"liveCodeAPI/api-master/master/usecases/categoryusecase"
	"liveCodeAPI/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

type CategoryHandler struct {
	categoryUseCase categoryusecase.CategoryUsecase
}

func CategoryController(r *mux.Router, service categoryusecase.CategoryUsecase) {

	categoryHandler := CategoryHandler{categoryUseCase: service}

	r.Use(middleware.ActivityLogMiddleware)

	category := r.PathPrefix("/category").Subrouter()
	category.HandleFunc("", categoryHandler.listCategory).Methods(http.MethodGet)
	category.Use(middleware.TokenValidationMiddleware)
	category.HandleFunc("/{id}", categoryHandler.category).Methods(http.MethodGet)
	category.HandleFunc("/add", categoryHandler.addCategory).Methods(http.MethodPost)
	category.HandleFunc("/delete/{id}", categoryHandler.deleteCategory).Methods(http.MethodDelete)
	category.HandleFunc("/update/{id}", categoryHandler.updateCategory).Methods(http.MethodPut)
}

func (s *CategoryHandler) listCategory(w http.ResponseWriter, r *http.Request) {
	category, err := s.categoryUseCase.GetCategories()
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = category
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (s *CategoryHandler) category(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	category, err := s.categoryUseCase.GetCategoryByID(id)
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = category
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (s *CategoryHandler) addCategory(w http.ResponseWriter, r *http.Request) {
	var categoryRequest []*models.Category
	_ = json.NewDecoder(r.Body).Decode(&categoryRequest)
	_, err := s.categoryUseCase.AddCategory(categoryRequest)
	if err != nil {
		w.Write([]byte("Cannot Add Data"))
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = categoryRequest
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (s *CategoryHandler) deleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	err := s.categoryUseCase.DeleteCategory(id)
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

func (s *CategoryHandler) updateCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var data models.Category
	_ = json.NewDecoder(r.Body).Decode(&data)

	category, err := s.categoryUseCase.UpdateCategory(id, &data)
	if err != nil {
		w.Write([]byte("Update Data Failed!"))
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = category
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}
