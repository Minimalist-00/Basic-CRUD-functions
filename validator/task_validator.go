package validator

import (
	"bulletin-board-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITaskValidator interface {
	TaskValidate(task model.Task) error //バリデーションで評価したいタスクの構造体を引数に取る
}

type taskValidator struct{}

func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

func (tv *taskValidator) TaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task, //構造体の検証
		validation.Field( //検証したいフィールドを指定
			&task.Title,
			validation.Required.Error("タイトルを入力してください"),
			validation.RuneLength(1, 10).Error("タイトルは10文字以内で入力してください"),
		),
	)
}
