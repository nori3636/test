package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// データ格納用
type Dist struct {
	num  int
	name string
}

// 大会でーた格納用
type Tournament struct {
	ID         int
	name       string
	regulation string
}

func createTable(db *sql.DB) error {
	q := `
	CREATE TABLE tournament (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(255),
		regulation VARCHAR(255)
	);
	`
	if _, err := db.Exec(q); err != nil {
		// log.Fatal(err)
		return err
	}
	return nil
}

func insertData(db *sql.DB) {
	// 大会の初期データ
	q := "INSERT into tournament (name, regulation) values (?, ?)"
	insertTournamentData(q, db, "CLAichi", "standard")
	insertTournamentData(q, db, "CLMiyagi", "standard")
	insertTournamentData(q, db, "CLKyouto", "extra")

	// デッキの初期データ
}

func insertTournamentData(query string, db *sql.DB, name string, regulation string) {
	_, err := db.Exec(query, name, regulation)
	if err != nil {
		log.Fatal(err)
	}
}
