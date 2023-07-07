package repository

/* データベース操作 */

import (
	"bulletin-board-rest-api/model"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IQuestRepository interface {
	GetAllQuestsFromDB(quests *[]model.Quest) error
	GetUserQuestsFromDB(quests *[]model.Quest, userId uint) error //全クエストを配列に格納する
	GetQuestById(quest *model.Quest, UserId uint, QuestId uint) error
	CreateQuest(quest *model.Quest) error
	UpdateQuest(quest *model.Quest, UserId uint, QuestId uint) error
	DeleteQuest(UserId uint, QuestId uint) error
	JoinQuest(UserId uint, QuestId uint) error
	CancelQuest(UserId uint, QuestId uint) error
}

type questRepository struct {
	db *gorm.DB
}

// コンストラクタの定義  引数はdb  戻り値はIQuestRepository
func NewQuestRepository(db *gorm.DB) IQuestRepository {
	return &questRepository{db}
}

func (qr *questRepository) GetAllQuestsFromDB(quests *[]model.Quest) error {
	if err := qr.db.Preload("User").Preload("Participants.User").Order("created_at DESC").Find(quests).Error; err != nil {
		// 募集主の情報 + 参加者の情報（全て）を取得。フィールド名を指定
		return err
	}
	return nil
}

func (qr *questRepository) GetUserQuestsFromDB(quests *[]model.Quest, userId uint) error {
	// クエスト一覧の中から、引数で渡されたuserIdと一致するクエスト一覧を取得する
	// UserテーブルのUserIdを参照 / created_atでソート / クエスト一覧をquestsに格納
	if err := qr.db.Joins("User").Where("user_id=?", userId).Order("created_at DESC").Find(quests).Error; err != nil {
		return err
	}
	return nil
}

func (qr *questRepository) GetQuestById(quest *model.Quest, userId uint, questId uint) error {
	// 指定されたUserIdのクエスト一覧で、QuestIdが一致するクエストを取得して quest に格納
	if err := qr.db.Joins("User").Where("user_id=?", userId).First(quest, questId).Error; err != nil {
		return err
	}
	return nil
}

func (qr *questRepository) CreateQuest(quest *model.Quest) error {
	if err := qr.db.Create(quest).Error; err != nil {
		return err
	}
	return nil
}

func (qr *questRepository) UpdateQuest(quest *model.Quest, userId uint, questId uint) error {
	result := qr.db.Model(quest).Clauses(clause.Returning{}).Where("id=? AND user_id=?", questId, userId).Updates(map[string]interface{}{
		"title":            quest.Title,
		"description":      quest.Description,
		"category":         quest.Category,
		"max_participants": quest.MaxParticipants,
		"deadline":         quest.Deadline,
		"start_time":       quest.StartTime,
		"end_time":         quest.EndTime,
		"image":            quest.Image,
		"url":              quest.URL,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist") // エラーメッセージ
	}
	return nil
}

func (qr *questRepository) DeleteQuest(userId uint, questId uint) error {
	result := qr.db.Where("id=? AND user_id=?", questId, userId).Delete(&model.Quest{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (qr *questRepository) JoinQuest(userId uint, questId uint) error {
	now := time.Now()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	participant := &model.QuestParticipant{
		JoinedAt: now.In(jst), // 現在時刻を取得
		UserId:   userId,
		QuestId:  questId,
	}
	if err := qr.db.Create(participant).Error; err != nil {
		return err
	}
	return nil
}

func (qr *questRepository) CancelQuest(userId uint, questId uint) error {
	result := qr.db.Where("quest_id=? AND user_id=?", questId, userId).Delete(&model.QuestParticipant{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
