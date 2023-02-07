package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DATABASE_PATH string = "./internal/database/sqlite_database.db"
)

var db *sql.DB

func NewSQLiteConnection() (err error) {

	db, err = sql.Open("sqlite3", DATABASE_PATH)
	if err != nil {
		return
	}

	err = db.Ping()

	return
}

func Query(sqlInstruction string, args ...any) (row *sql.Rows, err error) {

	if len(args) == 0 {
		row, err = db.Query(sqlInstruction)
		return
	}

	row, err = db.Query(sqlInstruction, args...)

	return
}

func QueryRow(sqlInstruction string, args ...any) (row *sql.Row) {

	if len(args) == 0 {
		row = db.QueryRow(sqlInstruction)
		return
	}

	row = db.QueryRow(sqlInstruction, args...)

	return
}

func ExecStatement(sqlInstruction string, args ...any) (err error) {

	var statement *sql.Stmt

	statement, err = db.Prepare(sqlInstruction)
	if err != nil {
		return
	}

	_, err = statement.Exec(args...)

	if err != nil {
		return
	}

	return
}
