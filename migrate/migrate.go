package main

/* データベースのマイグレーションを実行 */

import (
	"bulletin-board-rest-api/db"
	"bulletin-board-rest-api/model"
	"fmt"
)

func main() {
	dbConn := db.NewDB() //DB型のオブジェクトのアドレスを取得
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Quest{}, &model.QuestParticipant{}) //DBに反映させたいモデル構造のアドレスを取得して渡す
}
