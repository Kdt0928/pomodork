/*
 * POMODORK Backend API
 *
 * POMODORKのバックエンドAPI
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
	pomodork_constant "pomodork-backend/constant"
	"pomodork-backend/interface/controller"

	"github.com/gin-gonic/gin"
)

// UserControllerRepositories Controllerのインターフェース群
type UserControllerRepositories struct {
	UserAccountController controller.UserAccountControllerInterface
}

// NewUserControllerRepositories UserControllerRepositoriesの生成
func NewUserControllerRepositories() *UserControllerRepositories {
	controller := new(UserControllerRepositories)
	return controller
}

// UserDelete - ユーザ削除
func (u *UserControllerRepositories) UserDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UserCreate - ユーザ登録
func (u *UserControllerRepositories) UserCreate(c *gin.Context) {
	// Controllerの呼び出し
	ctx := context.Background()
	ctx = context.WithValue(ctx, pomodork_constant.TRANSACTION_ID, c.Request.Header.Get(pomodork_constant.TRANSACTION_ID))
	user_id, err := u.UserAccountController.CreateUser(ctx)
	if err != nil {
		// TODO エラーハンドリング
		c.JSON(http.StatusOK, gin.H{})
	}

	// 正常レスポンス
	response := InlineResponse200{
		UserId: user_id,
	}
	c.JSON(http.StatusOK, response)
}

// UserSearch - ユーザ取得
func (u *UserControllerRepositories) UserSearch(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
