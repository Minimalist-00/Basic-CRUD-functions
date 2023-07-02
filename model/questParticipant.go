package model

/* クエスト参加者を管理するテーブルの定義 */

import "time"

type QuestParticipant struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	QuestId  uint      `json:"quest_id" gorm:"not null"`
	UserId   uint      `json:"user_id" gorm:"not null"`
	JoinedAt time.Time `json:"joined_at"`
	User     User      `gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE"`  // ユーザーが削除されたら参加記録も削除
	Quest    Quest     `gorm:"foreignKey:QuestID; constraint:OnDelete:CASCADE"` // クエストが削除されたら参加記録も削除
}
