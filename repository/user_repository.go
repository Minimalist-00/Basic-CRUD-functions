package repository

/* データベース操作 */

import (
	"bulletin-board-rest-api/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error //引数のuserにemailをもつユーザーを格納｜返り値 エラーを返すときに使う
	CreateUser(user *model.User) error                   //引数のuserをDBに保存
	GetUserByID(user *model.User, userId uint) error
	UpdateUserName(userId uint, newUserName string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	// 1. データベースから指定された email に一致するユーザーを取得する。
	// 2. もしデータベース操作中にエラーが発生した場合、そのエラーを err 変数に代入する。
	// 3. err 変数が nil でない場合は、エラーを呼び出し元に返す。
	err := ur.db.Where("email = ?", email).First(user).Error //DBから指定されたemailに一致するユーザーを取得
	if err != nil {
		return err //そのままエラー文を返す
	} else {
		return nil
	}
}

func (ur *userRepository) CreateUser(user *model.User) error {
	err := ur.db.Create(user).Error //引数のuserをDBに保存
	if err != nil {
		return err
	} else {
		return nil
	}
}

// ユーザーIDを指定してユーザー情報を取得
func (ur *userRepository) GetUserByID(user *model.User, userId uint) error {
	err := ur.db.First(user, userId).Error // ユーザーIDを指定してユーザー情報を取得
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (ur *userRepository) UpdateUserName(userId uint, newUserName string) error {
	err := ur.db.Model(&model.User{}).Where("id = ?", userId).Update("name", newUserName).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}
