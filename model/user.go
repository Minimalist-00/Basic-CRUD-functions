package model

/* 作成するテーブルを定義するところ */

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	UserName  string    `json:"user_name" gorm:"unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// クライアントに返す情報
type UserResponse struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique"`
	UserName string `json:"user_name" gorm:"unique"`
}

// リクエストを格納する構造体
type UpdateUserNameRequest struct {
	UserName string `json:"user_name"`
}
