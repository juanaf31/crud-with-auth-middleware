package main

import (
	"liveCodeAPI/api-master/master"
	"liveCodeAPI/configs"
	"log"
)

func main() {
	db, err, dbHost, portServer := configs.InitDB()
	if err != nil {
		log.Println(err)
	}
	router := configs.CreatRouter()
	master.InitData(router, db)

	configs.RunServer(router, dbHost, portServer)
}
