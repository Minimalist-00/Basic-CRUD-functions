package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB { //*gorm.DB型のポインタを返す
	/* ローカルで実行する用のコード */
	if os.Getenv("GO_ENV") == "dev" { //環境変数がdevの場合
		err := godotenv.Load() //ローカルの.envファイルを読み込む
		if err != nil {
			log.Fatalln(err) //エラーがあればログに出力し強制終了
		}
	}
	// DBに接続するURL
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{}) //gorm.Open()でDBに接続
	if err != nil {                                          //エラー時の処理
		log.Fatalln(err)
	}
	fmt.Println("DB connected")
	return db
}

// DBをCLOSEする関数
func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB() //*gorm.DB オブジェクトから実際の *sql.DB オブジェクトへのアクセスを取得
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
