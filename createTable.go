package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// データ格納用

// 大会テーブル格納用
type Tournament struct {
	TID          int
	Tname        string
	Abbreviation string
	regulation   string
	league       string
	HPlink       string
}

//デッキテーブル格納用
type Deck struct {
	TID      int
	Rank     int
	DID      int
	DeckCode string
	blog     string
}

//デッキタイプ格納
type Deckname struct {
	DID      int
	Dtype    string
	Japanese string
	English  string
}

type Rule struct {
	regulation string
	league     string
}

func createTable(db *sql.DB) error {
	//大会テーブル作成
	// q := `
	// CREATE TABLE tournament (
	// 	TID INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	Tname VARCHAR(255),
	// 	Abbreviation VARCHAR(255),
	// 	regulation VARCHAR(255),
	// 	league VARCHAR(255),
	// 	HPlink VARCHAR(255),
	// 	FOREIGN KEY(regulation, league) REFERENCES Rule(regulation, league)
	// );
	// `
	//デッキテーブル作成
	// q := `
	// CREATE TABLE deck (
	// 	TID INTEGER,
	// 	Rank INTEGER,
	// 	DID INTEGER,
	// 	Deckcode VARCHAR(255),
	// 	blog VARCHAR(255),
	// 	FOREIGN KEY(TID) REFERENCES Tournament(TID)
	// );
	// `
	//デッキ名テーブル作成
	q := `
	CREATE TABLE deckname (
		TID INTEGER,
		japanese VARCHAR(255),
		english VARCHAR(255),
		decktype VARCHAR(255),
		FOREIGN KEY(TID) REFERENCES Tournament(TID)
	);
	 `
	// if _, err := db.Exec(q); err != nil {
	// 	// log.Fatal(err)
	// 	return err
	// }

	// q := `
	// CREATE TABLE rule (
	// 	regulation VARCHAR(255),
	// 	league VARCHAR(255)
	// );
	// `
	if _, err := db.Exec(q); err != nil {
		// log.Fatal(err)
		return err
	}
	return nil

}

func insertData(db *sql.DB) {
	// 大会のデータ
	// q := "INSERT into tournament (Tname, Abbreviation, regulation, league, HPlink) values (?, ?, ?, ?, ?)"
	// insertTournamentData(q, db, "ポケモンジャパンナショナルオンライン2020","PJNO","standard", "open", "https://pjnonline.net/")
	// insertTournamentData(q, db, "リザードンHR争奪戦", "シールド戦", "Special", "open", "https://www.pokemon-card.com/info/2020/20200703_002471.html")
	// 大会ルールの初期データ
	q := "INSERT into rule (regulation, league) values (?, ?)"
	insertRuleData(q, db, "standard", "open")
	insertRuleData(q, db, "Expanded", "master")
	insertRuleData(q, db, "HoF", "Senior")
	insertRuleData(q, db, "Special", "Junior")

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
