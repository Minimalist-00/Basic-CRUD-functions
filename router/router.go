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
	// ログイン関係のエンドポイントにの設定
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{ // CORSのミドルウェアの設定
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")}, // フロントエンドのURLを許可
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{ // CSRFのミドルウェアの設定
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		// CookieSameSite: http.SameSiteNoneMode, // フロントエンドとの通信にはSameSiteNoneModeを設定
		CookieSameSite: http.SameSiteDefaultMode, // Postmanでのテスト用
		// CookieMaxAge:   60, // csrfトークンの有効期限（秒）
	}))

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)
	t := e.Group("/quests") // クエスト関係のエンドポイントのグループ化
	// ミドルウェアの設定
	t.Use(echojwt.WithConfig(echojwt.Config{ //エンドポイントにミドルウェアの追加
		SigningKey:  []byte(os.Getenv("SECRET")), // 環境変数からシークレットキーを取得
		TokenLookup: "cookie:token",              // cookieからトークンを取得
	}))
	// クエスト関係のエンドポイントの設定
	t.GET("", qc.GetUserQuests)
	t.GET("/:questId", qc.GetQuestById)
	t.POST("", qc.CreateQuest)
	t.PUT("/:questId", qc.UpdateQuest)
	t.DELETE("/:questId", qc.DeleteQuest)
	return e
}
