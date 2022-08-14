package main

import (
	pomodork_constant "pomodork-backend/constant"
	"pomodork-backend/infrastructure/gorm"
	"pomodork-backend/infrastructure/openapi"
	"pomodork-backend/interface/controller"
	"pomodork-backend/interface/database"
	"pomodork-backend/usecase"
	pomodork_util "pomodork-backend/util"
)

func main() {

	// DB Connectionの生成
	db, err := gorm.NewGormHandler()
	if err != nil {
		// TODO
		return
	}
	// SQLHandlerの生成
	handler := database.NewUserAccountSqlHandlerRepositories()
	handler.SqlHandlerRepository = db
	// useCaseの生成
	userUserCase := usecase.NewUserAccountRepositories()
	userUserCase.UserAccountRepository = handler
	// Controllerの生成
	controller := controller.NewUserAccountRepositories()
	controller.UserAccountUseCase = userUserCase
	// Routerの生成
	users := openapi.NewUserControllerRepositories()
	users.UserAccountController = controller

	// Server起動
	apis := openapi.NewApiRepositories()
	apis.UserRouter = users
	gin := apis.NewRouter()

	// PORTの取得
	port, err := pomodork_util.GetEnv(pomodork_constant.SERVER_PORT, "string")
	if err != nil {
		return
	}
	gin.Run(":" + port.(string))
}
