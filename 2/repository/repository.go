package repository

import (
	"database/sql"
)

func InsertLog(pagination string, searchWord string) error {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/local")
	if err != nil {
		return err
	}
	defer db.Close()

	// asumsi id dan created_at auto generate
	query := "insert into logs (pagination, search_word) value (?, ?)"
	_, err = db.Exec(query, pagination, searchWord)
	if err != nil {
		return err
	}

	return nil
}
