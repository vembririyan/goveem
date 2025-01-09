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

func ConnectPostgre() error {
	var err error
	DB, err = PostgreEngine()
	DB_VERSION = "postgre"

	if err != nil {
		fmt.Println("Failure connected to Postgre!", time.Now())
		return err
	} else {
		fmt.Println("Successfully connected to the Postgre!", time.Now())
	}

	return nil
}

func CloseDBPostgre() {
	if err := DB.Close(); err != nil {
		log.Fatal("Error closing the database:", err)
	}
}

func PostgreEngine() (*sql.DB, error) {
	ENV := PostgreENV()
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", ENV["USERNAME"], ENV["PASSWORD"], ENV["HOST"], ENV["PORT"], ENV["DB_NAME"])
	var err error
	var CON *sql.DB
	CON, err = sql.Open("postgres", dsn)
	return CON, err
}

func PostgreENV() map[string]string {
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Error loading .env file")
		return nil
	}

	return map[string]string{
		"HOST":     os.Getenv("DB_HOST_PG"),
		"USERNAME": os.Getenv("DB_USERNAME_PG"),
		"PASSWORD": url.QueryEscape(os.Getenv("DB_PASSWORD_PG")),
		"DB_NAME":  os.Getenv("DB_NAME_PG"),
		"PORT":     os.Getenv("DB_PORT_PG")}
}
