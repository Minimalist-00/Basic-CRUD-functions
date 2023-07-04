package model

/* 作成するテーブルを定義するところ */

import "time"

type Quest struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Category        string    `json:"category"`
	Max_paticipants uint      `json:"max_paticipants" `
	Deadline        time.Time `json:"deadline" `
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time" gorm:"default: NULL"`
	Image           []byte    `json:"image"` // 画像をバイナリデータで保存
	URL             string    `json:"url"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User   User `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"` // UserIDのリレーション｜ユーザー削除時にクエストも消える
	UserId uint `json:"user_id" gorm:"not null"`
}

// クライアントに返す情報
type QuestResponse struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	Title           string    `json:"title" `
	Description     string    `json:"description"`
	Category        string    `json:"category" `
	Max_paticipants uint      `json:"max_paticipants" `
	Deadline        time.Time `json:"deadline" `
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
	Image           []byte    `json:"image"` // 画像をバイナリデータで保存
	URL             string    `json:"url"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
