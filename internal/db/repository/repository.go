package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	AddTicker(tickername string) error
	AddTickerPrice(tickeID int, price float64, timestamp string) error
	GetPriceAndDifference(ticker string, dateFrom, dateTo time.Time) (float64, float64, error)
	GetAllTickers() ([]string, error)
	GetTickerID(ticker string) (int, error)
}

type SQLRepository struct {
	db *sqlx.DB
}

func NewSQLRepository(db *sqlx.DB) *SQLRepository {
	return &SQLRepository{
		db: db,
	}
}

func (r *SQLRepository) AddTicker(tickername string) error {
	_, err := r.db.Exec("INSERT INTO tickers (ticker) VALUES ($1) ON CONFLICT (ticker) DO NOTHING", tickername)
	return err
}

func (r *SQLRepository) AddTickerPrice(tickeID int, price float64, timestamp string) error {
	_, err := r.db.Exec("INSERT INTO ticker_prices (ticker_id, price, timestamp) VALUES ($1, $2, $3) ", tickeID, price, timestamp)
	return err
}

func (r *SQLRepository) GetPriceAndDifference(ticker string, dateFrom, dateTo time.Time) (float64, float64, error) {
	var first_price, last_price float64

	query := `
        SELECT price
        FROM ticker_prices tp
        JOIN tickers t ON tp.ticker_id = t.id
        WHERE t.ticker = $1
        AND timestamp >= $2
        AND timestamp <= $3
        ORDER BY timestamp ASC
        LIMIT 1
    `
	if err := r.db.QueryRow(query, ticker, dateFrom, dateTo).Scan(&first_price); err != nil {
		return 0, 0, err
	}

	query = `
		SELECT price
		FROM ticker_prices tp
		JOIN tickers t ON tp.ticker_id = t.id
		WHERE t.ticker = $1
		AND timestamp >= $2
		AND timestamp <= $3
		ORDER BY timestamp DESC
		LIMIT 1
	`
	if err := r.db.QueryRow(query, ticker, dateFrom, dateTo).Scan(&last_price); err != nil {
		return 0, 0, err
	}

	return last_price, first_price, nil

}

func (r *SQLRepository) GetAllTickers() ([]string, error) {
	rows, err := r.db.Query(`SELECT ticker FROM tickers`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickers []string
	for rows.Next() {
		var ticker string
		if err := rows.Scan(&ticker); err != nil {
			return nil, err
		}
		tickers = append(tickers, ticker)
	}

	return tickers, nil
}

func (r *SQLRepository) GetTickerID(ticker string) (int, error) {
	var id int
	err := r.db.QueryRow("SELECT id FROM tickers WHERE ticker = $1", ticker).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
