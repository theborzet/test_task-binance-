package db

import (
	"fmt"
	"log"

	"test_api/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Init(c *config.Config) *sqlx.DB {
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.DBPort, c.User, c.Password, c.DBname)

	db, err := sqlx.Open("postgres", url)

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func Close(db *sqlx.DB) error {
	if err := db.Close(); err != nil {
		return err
	}
	return nil
}

func MigrateDB(db *sqlx.DB) {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS tickers (
			id SERIAL PRIMARY KEY,
			ticker VARCHAR(10) NOT NULL UNIQUE
		);
    `)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS ticker_prices (
			id SERIAL PRIMARY KEY,
			ticker_id INT REFERENCES tickers(id),
			price NUMERIC(20, 10),
			timestamp TIMESTAMP
		);
    `)
	if err != nil {
		log.Fatalln(err)
	}
}
