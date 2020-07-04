package configs

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreatRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}

func RunServer(router *mux.Router, host string, port string) {
	fmt.Println("starting on " + host + ":" + port)
	err := http.ListenAndServe(host+":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
