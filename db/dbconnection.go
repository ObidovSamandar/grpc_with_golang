package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DBPsql *sqlx.DB

func ConnectionDB() {

	godotenv.Load()

	psqlConnectionString := fmt.Sprintf("host=%v  port=%v user=%v password=%v dbname=%v sslmode=%v", os.Getenv("PG_HOST"), os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DB_NAME"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_SSL_MODE"))

	fmt.Println(psqlConnectionString)

	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=mylocaldatabase sslmode=disable")

	if err != nil {
		log.Fatalf("Failed to connect database!")
	}

	err = db.Ping()

	if err != nil {
		log.Fatalf("Failed to connect database!")
	}

	DBPsql = db
}
