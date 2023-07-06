package usecase

/* クエストに関連するビジネスロジックを実装する部分 */

import (
	"bulletin-board-rest-api/model"
	"bulletin-board-rest-api/repository"
	"bulletin-board-rest-api/validator"
)

type IQuestUsecase interface {
	GetAllQuests() ([]model.QuestResponse, error)
	GetUserQuests(userId uint) ([]model.QuestResponse, error)
	GetQuestById(userId uint, questId uint) (model.QuestResponse, error)
	CreateQuest(quest model.Quest) (model.QuestResponse, error)
	UpdateQuest(quest model.Quest, userId uint, questId uint) (model.QuestResponse, error)
	DeleteQuest(userId uint, questId uint) error
	JoinQuest(userId uint, questId uint) error
	CancelQuest(userId uint, questId uint) error
}

type questUsecase struct {
	// repositoryのinterfaceに依存
	qr repository.IQuestRepository //IQuestRepositoryを実装した構造体
	ur repository.IUserRepository
	qv validator.IQuestValidator
}

func NewQuestUsecase(qr repository.IQuestRepository, ur repository.IUserRepository, qv validator.IQuestValidator) IQuestUsecase {
	return &questUsecase{qr, ur, qv}
}

func (qu *questUsecase) GetAllQuests() ([]model.QuestResponse, error) {
	quests := []model.Quest{}
	if err := qu.qr.GetAllQuestsFromDB(&quests); err != nil {
		return nil, err
	}

	resQuests := []model.QuestResponse{} // QuestResponseの空の配列（スライス）を作成
	for _, quest := range quests {       // クエスト一覧の中身を1つずつ取り出す
		res := model.QuestResponse{
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
			UserName:        quest.User.UserName,                        // User構造体のUserNameを取得
			Participants:    make([]string, 0, len(quest.Participants)), // 参加者の名前の空のリストを作成
		}
		//* クエスト参加者情報から名前だけ取り出して配列に格納
		for _, p := range quest.Participants {
			res.Participants = append(res.Participants, p.User.UserName)
		}
		resQuests = append(resQuests, res)
	}
	return resQuests, nil
}

func (qu *questUsecase) GetUserQuests(userId uint) ([]model.QuestResponse, error) {
	quests := []model.Quest{}                                          //Questの配列（スライス）を作成
	if err := qu.qr.GetUserQuestsFromDB(&quests, userId); err != nil { //questRepositoryのGetAllQuestsFromDBを呼び出す -> questsに格納
		return nil, err
	}
	// 成功したときの処理
	resQuests := []model.QuestResponse{} //QuestResponseの配列を作成
	for _, quest := range quests {       //questsの中身を1つずつ取り出す
		res := model.QuestResponse{
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
		resQuests = append(resQuests, res) //resQuestsにmodel.QuestResponseを追加
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
	if err := qu.qr.CreateQuest(&quest); err != nil {
		return model.QuestResponse{}, err
	}

	// Get user info
	User := model.User{} //Userの空の構造体を作成
	//* questのUserIDからUser情報を取得する！
	if err := qu.ur.GetUserByID(&User, quest.UserId); err != nil {
		return model.QuestResponse{}, err
	}

	resQuest := model.QuestResponse{
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
		UserName:        User.UserName,
		Participants:    []string{},
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

func (qu *questUsecase) JoinQuest(userId uint, questId uint) error {
	return qu.qr.JoinQuest(userId, questId)
}

func (qu *questUsecase) CancelQuest(userId uint, questId uint) error {
	if err := qu.qr.CancelQuest(userId, questId); err != nil {
		return err
	}
	return nil
}
