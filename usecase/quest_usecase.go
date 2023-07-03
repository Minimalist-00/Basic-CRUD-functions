package usecase

/* クエストに関連するビジネスロジックを実装する部分 */

import (
	"bulletin-board-rest-api/model"
	"bulletin-board-rest-api/repository"
	"bulletin-board-rest-api/validator"
)

type IQuestUsecase interface {
	GetUserQuests(userId uint) ([]model.QuestResponse, error)
	GetQuestById(userId uint, questId uint) (model.QuestResponse, error)
	CreateQuest(quest model.Quest) (model.QuestResponse, error)
	UpdateQuest(quest model.Quest, userId uint, questId uint) (model.QuestResponse, error)
	DeleteQuest(userId uint, questId uint) error
}

type questUsecase struct {
	// repositoryのinterfaceに依存
	qr repository.IQuestRepository //IQuestRepositoryを実装した構造体
	qv validator.IQuestValidator
}

func NewQuestUsecase(qr repository.IQuestRepository, qv validator.IQuestValidator) IQuestUsecase {
	return &questUsecase{qr, qv}
}

func (qu *questUsecase) GetUserQuests(userId uint) ([]model.QuestResponse, error) {
	quests := []model.Quest{}                                          //Questの配列（スライス）を作成
	if err := qu.qr.GetUserQuestsFromDB(&quests, userId); err != nil { //questRepositoryのGetAllQuestsFromDBを呼び出す -> questsに格納
		return nil, err
	}
	// 成功したときの処理
	resQuests := []model.QuestResponse{} //QuestResponseの配列を作成
	for _, v := range quests {           //questsの中身を1つずつ取り出す
		q := model.QuestResponse{
			ID:              v.ID,
			Title:           v.Title,
			Description:     v.Description,
			Category:        v.Category,
			Max_paticipants: v.Max_paticipants,
			Deadline:        v.Deadline,
			StartTime:       v.StartTime,
			EndTime:         v.EndTime,
			Image:           v.Image,
			URL:             v.URL,
			CreatedAt:       v.CreatedAt,
			UpdatedAt:       v.UpdatedAt,
		}
		resQuests = append(resQuests, q) //resQuestsにmodel.QuestResponseを追加
	}
	return resQuests, nil
}

func (qu *questUsecase) GetQuestById(userId uint, questId uint) (model.QuestResponse, error) {
	quest := model.Quest{}                                              //Questの空の構造体を作成
	if err := qu.qr.GetQuestById(&quest, userId, questId); err != nil { //空の構造体とuser・questのIDを渡す
		return model.QuestResponse{}, err
	}
	resQuest := model.QuestResponse{ // QuestResponseのインスタンスを作成
		ID:              quest.ID,
		Title:           quest.Title,
		Description:     quest.Description,
		Category:        quest.Category,
		Max_paticipants: quest.Max_paticipants,
		Deadline:        quest.Deadline,
		StartTime:       quest.StartTime,
		EndTime:         quest.EndTime,
		Image:           quest.Image,
		URL:             quest.URL,
		CreatedAt:       quest.CreatedAt,
		UpdatedAt:       quest.UpdatedAt,
	}
	return resQuest, nil
}

func (qu *questUsecase) CreateQuest(quest model.Quest) (model.QuestResponse, error) {
	if err := qu.qv.QuestValidate(quest); err != nil {
		return model.QuestResponse{}, err
	}
	if err := qu.qr.CreateQuest(&quest); err != nil { //questRepositoryのCreateQuestを呼び出す
		return model.QuestResponse{}, err //QuestResponseの空の構造体とエラーを返す
	}
	resQuest := model.QuestResponse{ //QuestResponseの構造体を作成
		ID:              quest.ID,
		Title:           quest.Title,
		Description:     quest.Description,
		Category:        quest.Category,
		Max_paticipants: quest.Max_paticipants,
		Deadline:        quest.Deadline,
		StartTime:       quest.StartTime,
		EndTime:         quest.EndTime,
		Image:           quest.Image,
		URL:             quest.URL,
		CreatedAt:       quest.CreatedAt,
		UpdatedAt:       quest.UpdatedAt,
	}
	return resQuest, nil
}

func (qu *questUsecase) UpdateQuest(quest model.Quest, userId uint, questId uint) (model.QuestResponse, error) {
	if err := qu.qv.QuestValidate(quest); err != nil {
		return model.QuestResponse{}, err
	}
	if err := qu.qr.UpdateQuest(&quest, userId, questId); err != nil {
		return model.QuestResponse{}, err
	} //questのアドレスが指すメモリのクエストが更新後の値に書き換わっている
	resQuest := model.QuestResponse{ //QuestResponseの構造体を作成
		ID:              quest.ID,
		Title:           quest.Title,
		Description:     quest.Description,
		Category:        quest.Category,
		Max_paticipants: quest.Max_paticipants,
		Deadline:        quest.Deadline,
		StartTime:       quest.StartTime,
		EndTime:         quest.EndTime,
		Image:           quest.Image,
		URL:             quest.URL,
		CreatedAt:       quest.CreatedAt,
		UpdatedAt:       quest.UpdatedAt,
	}
	return resQuest, nil
}

func (qu *questUsecase) DeleteQuest(userId uint, questId uint) error {
	if err := qu.qr.DeleteQuest(userId, questId); err != nil {
		return err
	}
	return nil
}
