package controller

/*
  ・リクエストの受付 echo.Context エンドポイント
	・認証認可
	・ビジネスロジックの呼び出し（useCase）
	・レスポンスの作成
*/

import (
	"bulletin-board-rest-api/model"
	"bulletin-board-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IQuestController interface {
	GetAllQuests(c echo.Context) error
	GetUserQuests(c echo.Context) error
	GetQuestById(c echo.Context) error
	CreateQuest(c echo.Context) error
	UpdateQuest(c echo.Context) error
	DeleteQuest(c echo.Context) error
	JoinQuest(c echo.Context) error
	CancelQuest(c echo.Context) error
}

type questController struct {
	qu usecase.IQuestUsecase
}

// usecaseを「Dependency Injection」するための関数（コンストラクタ）
// usecaseのインスタンスを受け取る
func NewQuestController(qu usecase.IQuestUsecase) IQuestController {
	return &questController{qu}
}

func (qc *questController) GetAllQuests(c echo.Context) error {
	questsRes, err := qc.qu.GetAllQuests() // 引数はないよ
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, questsRes)
}

func (qc *questController) GetUserQuests(c echo.Context) error {
	// JWTのclaimsからユーザーIDを取得
	user := c.Get("user").(*jwt.Token) // jwtをデコードした内容を取得
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// ユーザーIDを元にクエストを取得
	questsRes, err := qc.qu.GetUserQuests(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, questsRes)
}

func (qc *questController) GetQuestById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("questId")       // クエストIDを取得！
	questId, _ := strconv.Atoi(id) // string型 -> int型に変換
	questRes, err := qc.qu.GetQuestById(uint(userId.(float64)), uint(questId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, questRes)
}

func (qc *questController) CreateQuest(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	quest := model.Quest{}
	if err := c.Bind(&quest); err != nil { // リクエストボディをquestにバインド
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	quest.UserId = uint(userId.(float64))
	err := qc.qu.CreateQuest(quest) // 返り値を無視する
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated) // 201 Created ステータスのみを返す
}

func (qc *questController) UpdateQuest(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("questId")
	questId, _ := strconv.Atoi(id)

	quest := model.Quest{}
	if err := c.Bind(&quest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err := qc.qu.UpdateQuest(quest, uint(userId.(float64)), uint(questId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

/*
* jwtトークンからuserIDの取得 + パラメータからQuestIDを取得
* userIDとQuestIDを元にクエストを削除するusecaseのメソッドを呼び出す
 */
func (qc *questController) DeleteQuest(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("questId")
	questId, _ := strconv.Atoi(id)

	err := qc.qu.DeleteQuest(uint(userId.(float64)), uint(questId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (qc *questController) JoinQuest(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("questId")
	questId, _ := strconv.Atoi(id)

	err := qc.qu.JoinQuest(uint(userId.(float64)), uint(questId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (qc *questController) CancelQuest(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("questId")
	questId, _ := strconv.Atoi(id)

	err := qc.qu.CancelQuest(uint(userId.(float64)), uint(questId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
