package validator

import (
	"errors"
	"strings"

	"bulletin-board-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	ValidateUserSignUp(user model.User) error
	ValidateUserLogIn(user model.User) error
}

type userValidator struct{}
type allowedEmailRule struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (a *allowedEmailRule) Validate(value interface{}) error {
	email, _ := value.(string)
	if strings.HasSuffix(email, "@st.pu-toyama.ac.jp") || strings.HasSuffix(email, "@puc.pu-toyama.ac.jp") {
		return nil
	}
	return errors.New("学内メールアドレスを入力してください")
}

func (uv *userValidator) ValidateUserSignUp(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("メールアドレスを入力してください"),
			validation.RuneLength(1, 30).Error("メールアドレスは30文字以内で入力してください"),
			is.Email.Error("入力されたメールアドレスの形式が適切ではありません"),
			&allowedEmailRule{}, // ここで新しいバリデーションルールを追加
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("パスワードを入力してください"),
			validation.RuneLength(6, 20).Error("パスワードは6～20文字以内で入力してください"),
		),
		validation.Field(
			&user.UserName,
			validation.Required.Error("ユーザー名を入力してください"),
			validation.RuneLength(1, 10).Error("ユーザー名は10文字以内で入力してください"),
		),
	)
}

func (uv *userValidator) ValidateUserLogIn(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("メールアドレスを入力してください"),
			validation.RuneLength(1, 30).Error("メールアドレスは30文字以内で入力してください"),
			is.Email.Error("入力されたメールアドレスの形式が適切ではありません"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("パスワードを入力してください"),
			validation.RuneLength(6, 20).Error("パスワードは6～20文字以内で入力してください"),
		),
	)
}
