package controller

/* リクエストの受け付けとレスポンスの生成 */

import (
	"bulletin-board-rest-api/model"
	"bulletin-board-rest-api/usecase"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
	GetUserName(c echo.Context) error
	GetUserInfo(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

// usecaseを「Dependency Injection」するための関数（コンストラクタ）
// usecaseのインスタンスを受け取る
func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil { //リクエストボディをuserにバインド（User型に変換して格納）
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.uu.SignUp(user) //usecaseのSignUpメソッドを呼び出し
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes) //Created(201)のステータスと、作成したユーザー情報を返す
}

func (uc *userController) LogIn(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.uu.Login(user) //usecaseのLoginメソッドを呼び出し（JWTtokenが入る）
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// Cookieの設定
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(time.Hour * 24 * 1) //TODO: 有効期限
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true //TODO: 本番環境:ture ｜ postman test:false
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie) //*作成したCookieをセット
	return c.NoContent(http.StatusOK)
}

func (uc *userController) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true //TODO: 本番環境:ture ｜ postman test:false
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}

func (uc *userController) GetUserName(c echo.Context) error {
	// JWTのclaimsからユーザーIDを取得
	user := c.Get("user").(*jwt.Token) // jwtをデコードした内容を取得
	claims := user.Claims.(jwt.MapClaims)
	userId := uint(claims["user_id"].(float64)) // float64をuintにキャスト

	// ユーザーIDを元にユーザー名を取得
	username, err := uc.uu.GetUserName(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, username) //* ここでUserNameを取得してJSON形式で返す！
}

func (uc *userController) GetUserInfo(c echo.Context) error {
	// JWTのclaimsからユーザーIDを取得
	user := c.Get("user").(*jwt.Token) // jwtをデコードした内容を取得
	claims := user.Claims.(jwt.MapClaims)
	userId := uint(claims["user_id"].(float64)) // float64をuintにキャスト

	// ユーザーIDを元にユーザー名を取得
	userRes, err := uc.uu.GetUserInfo(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	response := map[string]interface{}{
		"email":     userRes.Email,
		"user_name": userRes.UserName,
	}

	return c.JSON(http.StatusOK, response) //* ここでUserNameを取得してJSON形式で返す！
}
