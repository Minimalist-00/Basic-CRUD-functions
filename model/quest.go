package model

/* 作成するテーブルを定義するところ */

import "time"

type Quest struct {
	ID              uint               `json:"id" gorm:"primaryKey"`
	Title           string             `json:"title"`
	Description     string             `json:"description"`
	Category        string             `json:"category"`
	MaxParticipants uint               `json:"max_participants" `
	Deadline        time.Time          `json:"deadline" `
	StartTime       time.Time          `json:"start_time"`
	EndTime         time.Time          `json:"end_time"`
	Image           []byte             `json:"image"` // 画像をバイナリデータで保存
	URL             string             `json:"url"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
	User            User               `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"` // UserIDを元にUserテーブルと紐付ける
	UserId          uint               `json:"user_id" gorm:"not null"`
	Participants    []QuestParticipant `json:"participants" gorm:"foreignKey:QuestId"` // QuestParticipantテーブルと紐付ける
}

// クライアントに返す情報
type QuestResponse struct {
	ID              uint       `json:"id" gorm:"primaryKey"`
	Title           string     `json:"title" `
	Description     string     `json:"description"`
	Category        string     `json:"category" `
	MaxParticipants uint       `json:"max_participants" `
	Deadline        *time.Time `json:"deadline" `
	StartTime       *time.Time `json:"start_time"`
	EndTime         *time.Time `json:"end_time"`
	Image           []byte     `json:"image"` // 画像をバイナリデータで保存
	URL             string     `json:"url"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	UserName        string     `json:"user_name"`    // 作成者の名前
	Participants    []string   `json:"participants"` // 参加者の名前のリスト
}

type EditQuestResponse struct {
	Title           string     `json:"title" `
	Description     string     `json:"description"`
	Category        string     `json:"category" `
	MaxParticipants uint       `json:"max_participants" `
	Deadline        *time.Time `json:"deadline" `
	StartTime       *time.Time `json:"start_time"`
	EndTime         *time.Time `json:"end_time"`
	Image           []byte     `json:"image"` // 画像をバイナリデータで保存
	URL             string     `json:"url"`
}
