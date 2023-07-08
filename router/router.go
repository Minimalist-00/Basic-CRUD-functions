package router

/*
ルーティングとコントローラの結びつけ
 1. エンドポイントの設定
 2. ミドルウェアの設定
 3. コントローラとの結びつけ
*/

import (
	"bulletin-board-rest-api/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, qc controller.IQuestController) *echo.Echo {
	e := echo.New()

	//* CORSのミドルウェアの設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")}, // フロントエンドのURLを許可
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))

	//* CSRFのミドルウェアの設定
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode, //TODO: フロントエンドとの通信にはSameSiteNoneModeを使う
		// CookieSameSite: http.SameSiteDefaultMode, //TODO: Postmanでのテスト用
		// CookieMaxAge:   60, // csrfトークンの有効期限（秒）
	}))

	//* ログイン関係のエンドポイントの設定
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)

	//* ユーザー関係のエンドポイントの設定
	u := e.Group("/users")
	u.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	u.GET("/userName", uc.GetUserName)
	u.GET("/userInfo", uc.GetUserInfo)

	//* ミドルウェアの設定
	q := e.Group("/quests")                  // クエスト関係のエンドポイントのグループ化
	q.Use(echojwt.WithConfig(echojwt.Config{ //エンドポイントにミドルウェアの追加
		SigningKey:  []byte(os.Getenv("SECRET")), // 環境変数からシークレットキーを取得
		TokenLookup: "cookie:token",              // cookieからトークンを取得
	}))

	//* クエスト関係のエンドポイントの設定
	q.GET("", qc.GetAllQuests)
	q.GET("/:questId", qc.GetQuestById)
	q.POST("", qc.CreateQuest)
	q.PUT("/:questId", qc.UpdateQuest)
	q.DELETE("/:questId", qc.DeleteQuest)

	q.POST("/join/:questId", qc.JoinQuest) // クエストの参加
	q.DELETE("/cancel/:questId", qc.CancelQuest)
	q.GET("/created", qc.GetUserQuests)  // ユーザーが作成したクエスト一覧
	q.GET("/joined", qc.GetJoinedQuests) // ユーザーが参加したクエスト一覧
	return e
}
