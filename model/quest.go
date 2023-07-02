package model

/* 作成するテーブルを定義するところ */

import "time"

type Quest struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"` //空の値を許可しない
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// どのユーザーかを分かるように追加
	User   User `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"` // UserIDのリレーション｜ユーザー削除時にクエストも消える
	UserId uint `json:"user_id" gorm:"not null"`
}

// クライアントに返す情報
type QuestResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
