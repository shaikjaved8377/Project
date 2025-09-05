package api

import (
	"database/sql"
	"net/http"

	"Project/dataservice"
)

func CreateBookLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	return dataservice.CreateBook(db, w, r)
}
