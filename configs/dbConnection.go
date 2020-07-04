package configs

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func readGoEnvVar(key string) string {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func InitDB() (*sql.DB, error, string, string) {
	dbUser := readGoEnvVar("dbUser")
	dbPass := readGoEnvVar("dbPass")
	dbHost := readGoEnvVar("dbHost")
	dbPort := readGoEnvVar("dbPort")
	dbName := readGoEnvVar("dbName")
	portServer := readGoEnvVar("portServer")

	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, _ := sql.Open("mysql", source)
	err := db.Ping()
	if err != nil {
		log.Print(err)
	}
	fmt.Println("Database Connected")

	return db, nil, dbHost, portServer

}

//ConnectDB adalah fungsi untuk koneksi ke database
func connectDB(source string) (*sql.DB, error) {

	db, _ := sql.Open("mysql", source)
	err := db.Ping()
	if err != nil {
		log.Print(err)
	}
	fmt.Println("Database Connected")

	return db, nil
}
