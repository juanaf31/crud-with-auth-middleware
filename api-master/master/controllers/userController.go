package controllers

import (
	"encoding/json"
	"fmt"
	"liveCodeAPI/api-master/master/models"
	"liveCodeAPI/api-master/master/usecases/userusecase"
	"liveCodeAPI/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	userUseCase userusecase.UserUsecase
}

func UserController(r *mux.Router, service userusecase.UserUsecase) {
	userHandler := UserHandler{userUseCase: service}
	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("", userHandler.getUser).Methods(http.MethodPost)
}

func (u *UserHandler) getUser(w http.ResponseWriter, r *http.Request) {
	var data models.User
	_ = json.NewDecoder(r.Body).Decode(&data)
	isValid, _ := u.userUseCase.GetUser(&data)

	if isValid {
		w.WriteHeader(http.StatusOK)
		token, err := utils.JwtEncoder(data.UserName, "rahasiadong")
		if err != nil {
			http.Error(w, "Failed token generation", http.StatusUnauthorized)
		} else {
			w.Write([]byte(token))
		}
	} else {
		http.Error(w, "Invalid login", http.StatusUnauthorized)
	}

	fmt.Println(isValid)
}
