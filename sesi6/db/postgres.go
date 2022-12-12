package db

import (
	"database/sql"

	"github.com/antonlindstrom/pgstore"
	_ "github.com/lib/pq"
)

func NewPostgres() (*sql.DB, error) {
	dsn := "postgresql://noobee:iniPassword@localhost:6432/bank-neo?sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewPostgresStore(db *sql.DB) (*pgstore.PGStore, error) {

	authKey := []byte("authkey")
	encryptionKey := []byte("qMT53qxDCPmRhivTW7nLWVn3oLDdkiRT")

	store, err := pgstore.NewPGStoreFromPool(db, authKey, encryptionKey)
	if err != nil {
		return nil, err
	}

	return store, nil
}
