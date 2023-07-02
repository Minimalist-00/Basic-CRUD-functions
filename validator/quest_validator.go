package validator

import (
	"bulletin-board-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IQuestValidator interface {
	QuestValidate(quest model.Quest) error //バリデーションで評価したいクエストの構造体を引数に取る
}

type questValidator struct{}

func NewQuestValidator() IQuestValidator {
	return &questValidator{}
}

func (qv *questValidator) QuestValidate(quest model.Quest) error {
	return validation.ValidateStruct(&quest, //構造体の検証
		validation.Field( //検証したいフィールドを指定
			&quest.Title,
			validation.Required.Error("タイトルを入力してください"),
			validation.RuneLength(1, 10).Error("タイトルは10文字以内で入力してください"),
		),
	)
}
