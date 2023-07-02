package usecase

/* クエストに関連するビジネスロジックを実装する部分 */

import (
	"bulletin-board-rest-api/model"
	"bulletin-board-rest-api/repository"
	"bulletin-board-rest-api/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error) //JWTを返すためにstring型
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}
	//パスワードのハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10) //GenerateFromPassword関数により、パスワードをハッシュ化
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{Email: user.Email, Password: string(hash)} //ハッシュ化したパスワードをnewUserに格納
	if err := uu.ur.CreateUser(&newUser); err != nil {               //引数のnewUserをDBに保存
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{ //レスポンス用のUserResponse型の変数を作成
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err //Loginメソッドの戻り値に適するように空の文字列とエラーを返す
	}
	//クライアントからのEmailがDB内に存在するかを確認
	storedUser := model.User{} //DBから取得したユーザー情報を格納するための変数
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	// ハッシュ化されたパスと元のパスの一致を比較
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ //JWTの生成
		"user_id": storedUser.ID,                         //ユーザーIDの設定
		"exp":     time.Now().Add(time.Hour * 12).Unix(), //有効期限の設定
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET"))) //著名済みの文字列を生成 <- トークンとして使うことで、信頼性↑
	if err != nil {
		return "", err
	}
	return tokenString, nil //JWTを返す
}
