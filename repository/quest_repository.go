package repository

/* データベース操作 */

import (
	"bulletin-board-rest-api/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IQuestRepository interface {
	GetUserQuestsFromDB(quests *[]model.Quest, userId uint) error //全クエストを配列に格納する
	GetQuestById(quest *model.Quest, UserId uint, QuestId uint) error
	CreateQuest(quest *model.Quest) error
	UpdateQuest(quest *model.Quest, UserId uint, QuestId uint) error
	DeleteQuest(UserId uint, QuestId uint) error
}

type questRepository struct {
	db *gorm.DB
}

// コンストラクタの定義  引数はdb  戻り値はIQuestRepository
func NewQuestRepository(db *gorm.DB) IQuestRepository {
	return &questRepository{db}
}

// DBからクエストの一覧を取得
func (qr *questRepository) GetUserQuestsFromDB(quests *[]model.Quest, userId uint) error {
	// クエスト一覧の中から、引数で渡されたuserIdと一致するクエスト一覧を取得する
	// UserテーブルのUserIdを参照 / created_atでソート / クエスト一覧をquestsに格納
	if err := qr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(quests).Error; err != nil {
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
	result := qr.db.Model(quest).Clauses(clause.Returning{}).Where("id=? AND user_id=?", questId, userId).Update("title", quest.Title)
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