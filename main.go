package main

import (
    // ビルド時のみ使用する
    "database/sql"
	"log"
	"fmt"

    _ "github.com/mattn/go-sqlite3"
)

// create table dist (num int primary key, name text);
// insert into dist values(1, 'redhat');
// insert into dist values(2, 'ubuntu');
// insert into dist values(3, 'debian');

// DB Path(相対パスでも大丈夫かと思うが、筆者の場合、絶対パスでないと実行できなかった)
const dbPath = "./test.DB"

// コネクションプールを作成
var DbConnection *sql.DB

// データ格納用
type Dist struct {
    num int
    name string
}

func main() {
    // Open(driver,  sql 名(任意の名前))
	DbConnection, _ := sql.Open("sqlite3", dbPath)

    // Connection をクローズする。(defer で閉じるのが Golang の作法)
    defer DbConnection.Close()

	cmd := "SELECT * FROM dist"
    rows, _ := DbConnection.Query(cmd)

    defer rows.Close()

    // データ保存領域を確保
    var dists []Dist
    for rows.Next() {
        var d Dist
        // Scan にて、struct のアドレスにデータを入れる
        err := rows.Scan(&d.num, &d.name)
        // エラーハンドリング(共通関数にした方がいいのかな)
        if err != nil {
            log.Println(err)
        }
        // データ取得
        dists = append(dists, d)
    }

    // 操作結果を確認
    for _, d := range dists {
        fmt.Println(d.num, d.name)
    }
}

