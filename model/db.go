package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

type connection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBname   string
}

func Init() {
	err := godotenv.Load("app/config/.env")

	if err != nil {
		log.Fatalf("Error loading env file: %s", err.Error())
	}

	connInfo := connection{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		DBname:   os.Getenv("MYSQL_DBNAME"),
	}

	db, err = sql.Open("mysql", connToString(connInfo))
	if err != nil {
		log.Fatalf("Error connecting to the DataBase: %s", err.Error())

	} else {
		fmt.Print("DataBase connection succeed\n")
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error could not ping DataBase: %s", err.Error())
	} else {
		fmt.Print("DataBases ping succeed")
	}
}

//Runs locally
func connToString(info connection) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", info.User, info.Password, info.Host, info.Port, info.DBname)
}

//Runs on docker
//func connToString(info connection) string {
//	return fmt.Sprintf("%s:%s@tcp(mysql)/%s?charset=utf8mb4&parseTime=True&loc=Local", info.User, info.Password, info.DBname)
//}
