package main

import (
	// ビルド時のみ使用する
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// DB Path(相対パスでも大丈夫かと思うが、筆者の場合、絶対パスでないと実行できなかった)
const dbPath = "./test.DB"

// コネクションプールを作成
var DbConnection *sql.DB

func main() {

	// f, err := os.Open("PJNO.csv")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// r := csv.NewReader(f)

	// for i := 0; i < 5; i++ {
	// 	record, err := r.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	slice := strings.Split(f, ",")
	// 	for _, x := range slice {
	// 		fmt.Println(x)
	// 	}

	// }

	// Open(driver,  sql 名(任意の名前))
	DbConnection, _ := sql.Open("sqlite3", dbPath)

	// Connection をクローズする。(defer で閉じるのが Golang の作法)
	defer DbConnection.Close()

	err := createTable(DbConnection)
	if err != nil {
		log.Fatal(err)
	}

	//insertData(DbConnection)
	// //getRuleData(DbConnection)
	// cmd := "SELECT * FROM Rule"
	// rows, _ := DbConnection.Query(cmd)

	// // データ保存領域を確保
	// var rule []Rule
	// for rows.Next() {
	// 	var r Rule
	// 	// Scan にて、struct のアドレスにデータを入れる
	// 	err := rows.Scan(&r.regulation, &r.league)
	// 	// エラーハンドリング(共通関数にした方がいいのかな)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	// データ取得
	// 	rule = append(rule, r)
	// }

	// // 操作結果を確認
	// for _, r := range rule {
	// 	fmt.Println(r.regulation, r.league)
	// }

	// defer rows.Close()

	// cmd := "SELECT * FROM tournament"
	// rows, _ := DbConnection.Query(cmd)

	// // データ保存領域を確保
	// var tournaments []Tournament
	// for rows.Next() {
	// 	var t Tournament
	// 	// Scan にて、struct のアドレスにデータを入れる
	// 	err := rows.Scan(&t.ID, &t.name, &t.regulation)
	// 	// エラーハンドリング(共通関数にした方がいいのかな)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	// データ取得
	// 	tournaments = append(tournaments, t)
	// }

	// // 操作結果を確認
	// for _, t := range tournaments {
	// 	fmt.Println(t.ID, t.name, t.regulation)
	// }

	// defer rows.Close()

	// cmd = "SELECT * FROM player"
	// playerRows, _ := DbConnection.Query(cmd)

	// defer playerRows.Close()

	// // データ保存領域を確保
	// var players []Player
	// for playerRows.Next() {
	// 	var p Player
	// 	// Scan にて、struct のアドレスにデータを入れる
	// 	err := playerRows.Scan(&p.ID, &p.name, &p.twitter, &p.blog)
	// 	// エラーハンドリング(共通関数にした方がいいのかな)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	// データ取得
	// 	players = append(players, p)
	// }

	// // 操作結果を確認
	// for _, p := range players {
	// 	fmt.Println(p.ID, p.name, p.twitter, p.blog)
	// }

}
