package goveem

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var DB *sql.DB
var DB_VERSION string

func ConnectMySQL() error {
	var err error
	DB, err = MySQLEngine()
	DB_VERSION = "mysql"

	if err != nil {
		fmt.Println("Failure connected to Postgre!", time.Now())
		return err
	} else {
		fmt.Println("Successfully connected to the Postgre!", time.Now())
	}

	return nil
}

func CloseDBMySQL() {
	if err := DB.Close(); err != nil {
		log.Fatal("Error closing the database:", err)
	}
}

func MySQLEngine() (*sql.DB, error) {
	ENV := PostgreENV()
	dsn := ENV["USERNAME"] + ":" + ENV["PASSWORD"] + "@tcp(" + ENV["HOST"] + ":" + ENV["PORT"] + ")/" + ENV["DB_NAME"]
	var err error
	var CON *sql.DB
	CON, err = sql.Open("mysql", dsn)
	return CON, err
}

func MysqlENV() map[string]string {
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Error loading .env file")
		return nil
	}

	return map[string]string{
		"HOST":     os.Getenv("DB_HOST"),
		"USERNAME": os.Getenv("DB_USERNAME"),
		"PASSWORD": url.QueryEscape(os.Getenv("DB_PASSWORD")),
		"DB_NAME":  os.Getenv("DB_NAME"),
		"PORT":     os.Getenv("DB_PORT")}
}
