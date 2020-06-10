package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/ahamtat/go-rest-api/internal/app/store/sqlstore"
)

func Start(config *Config) error {
	db, _ := newDB(config.DatabaseURL)
	defer db.Close()

	store := sqlstore.New(db)
	srv := newServer(store)
	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
