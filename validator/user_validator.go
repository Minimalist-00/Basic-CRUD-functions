package validator

import (
	"bulletin-board-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	UserValidate(user model.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(&user, //構造体の検証
		validation.Field( //検証したいフィールドを指定
			&user.Email,
			validation.Required.Error("メールアドレスを入力してください"),
			validation.RuneLength(1, 30).Error("メールアドレスは30文字以内で入力してください"),
			is.Email.Error("入力されたメールアドレスの形式が適切ではありません"), //Emailのフォーマットに準拠しているか
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("パスワードを入力してください"),
			validation.RuneLength(6, 30).Error("パスワードは6～30文字以内で入力してください"),
		),
	)
}
