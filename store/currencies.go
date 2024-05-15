package store

import (
	"database/sql"
	"log"
)

func currenciesDB() *sql.DB {
	db, err := database("currencies.db")
	if err != nil {
		log.Fatal(err)
	}

	db.Exec(`CREATE TABLE IF NOT EXISTS currencies (
		iso TEXT PRIMARY KEY,
		usd_value DECIMAL(10, 6) NOT NULL,
		last_updated INTEGER NOT NULL DEFAULT 0
	)`)

	return db
}

func GetCurrencyUsdValue(iso string) (float64, int, error) {
	db := currenciesDB()
	defer db.Close()

	var value float64
	var lastUpdated int
	err := db.QueryRow("SELECT usd_value, last_updated FROM currencies WHERE iso = ?", iso).Scan(&value, &lastUpdated)
	if err != nil {
		return 0, 0, err
	}

	return value, lastUpdated, nil
}

func SaveCurrencyUsdValue(iso string, value float64, lastUpdated int) {
	db := currenciesDB()
	defer db.Close()

	_, err := db.Exec("INSERT OR REPLACE INTO currencies (iso, usd_value, last_updated) VALUES (?, ?, ?)", iso, value, lastUpdated)
	if err != nil {
		log.Fatal(err)
	}
}
