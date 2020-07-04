package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)
// データ格納用

// 大会テーブル格納用
type Tournament struct {
	TID        int
	Tname      string
	regulation string
	league     string
	HPlink     string
}

//デッキテーブル格納用
type Deck struct {
	TID      int
	Rank     int
	DID      int
	PID      int
	DeckCode string
	blog     string
}

//デッキタイプ格納
type Deckname struct {
	DID      int
	Japanese string
	English  string
}

type Player struct {
	PID     int
	name    string
	Twitter string
}

type Rule struct {
	regulation string
	league     string
}

func createTable(db *sql.DB) error {
	// q := `
	// CREATE TABLE tournament (
	// 	TID INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	Tname VARCHAR(255),
	// 	regulation VARCHAR(255),
	// 	league VARCHAR(255),
	// 	HPlink VARCHAR(255),
	// 	FOREIGN KEY(regulation, league) REFERENCES Rule(regulation, league)
	// );
	// `
	// q = `
	// CREATE TABLE deck (
	// 	TID INTEGER,
	// 	Rank INTEGER,
	// 	DID INTEGER,
	// 	PID INTEGER,
	// 	Deckcode VARCHAR(255),
	// 	blog VARCHAR(255),
	// 	FOREIGN KEY(TID) REFERENCES Tournament(TID),
	//	FOREIGN KEY(PID) REFERENCES Player(PID),
	// );
	// `
	// if _, err := db.Exec(q); err != nil {
	// 	// log.Fatal(err)
	// 	return err
	// }

	q := `
	CREATE TABLE rule (
		regulation VARCHAR(255),
		league VARCHAR(255)
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
	// q := "INSERT into tournament (name, regulation) values (?, ?)"
	// insertTournamentData(q, db, "CLAichi", "standard")
	// insertTournamentData(q, db, "CLMiyagi", "standard")
	// insertTournamentData(q, db, "CLKyouto", "extra")
	// 大会ルールの初期データ
	q := "INSERT into rule (regulation, league) values (?, ?)"
	insertRuleData(q, db, "stan", "open")
	// insertRuleData(q, db, "エクストラ", "マスター")
	// insertRuleData(q, db, "殿堂", "シニア")
	// insertRuleData(q, db, "特殊ルール", "ジュニア")

}

// func getRuleData(db *sql.DB) {
// 	cmd := "SELECT * FROM Rule"
// 	rows, _ := DbConnection.Query(cmd)

// 	// データ保存領域を確保
// 	var rule []Rule
// 	for rows.Next() {
// 		var r Rule
// 		// Scan にて、struct のアドレスにデータを入れる
// 		err := rows.Scan(&r.regulation, &r.league)
// 		// エラーハンドリング(共通関数にした方がいいのかな)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		// データ取得
// 		rule = append(rule, r)
// 	}

// 	// 操作結果を確認
// 	for _, r := range rule {
// 		fmt.Println(r.regulation, r.league)
// 	}

// 	defer rows.Close()

// }

func insertRuleData(query string, db *sql.DB, regulation string, league string) {
	_, err := db.Exec(query, regulation, league)
	if err != nil {
		log.Fatal(err)
	}
}

func insertTournamentData(query string, db *sql.DB, name string, regulation string) {
	_, err := db.Exec(query, name, regulation)
	if err != nil {
		log.Fatal(err)
	}
}

func insertPlayerData(query string, db *sql.DB, name string, twitter string, blog string) {
	_, err := db.Exec(query, name, twitter, blog)
	if err != nil {
		log.Fatal(err)
	}
}
