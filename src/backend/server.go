package main

import (
	"context"
	pomodork_constant "pomodork-backend/constant"
	pomodork_gorm "pomodork-backend/infrastructure/gorm"
	pomodork_openapi "pomodork-backend/infrastructure/openapi"
	pomodork_controller "pomodork-backend/interface/controller"
	pomodork_database "pomodork-backend/interface/database"
	pomodork_error "pomodork-backend/message/error"
	pomodork_usecase "pomodork-backend/usecase"
	pomodork_util "pomodork-backend/util"
)

func main() {

	logger := pomodork_util.NewLogger()

	// DB Connectionの生成
	db, err := pomodork_gorm.NewGormHandler()
	if err != nil {
		serverFatalErr := pomodork_error.ServeFatalError{}
		logger.Error(context.Background(), serverFatalErr.Code(), serverFatalErr.Error())
		return
	}
	// SQLHandlerの生成
	handler := pomodork_database.NewSqlRepositories()
	handler.SqlHandler = db
	// useCaseの生成
	userUserCase := pomodork_usecase.NewUserAccountRepositories()
	userUserCase.UserAccountRepository = handler
	// Controllerの生成
	controller := pomodork_controller.NewUserAccountRepositories()
	controller.Domain = userUserCase
	// Routerの生成
	users := pomodork_openapi.NewUserRepositories()
	users.UserAccountController = controller

	// Server起動
	apis := pomodork_openapi.NewApiRepositories()
	apis.UserRouter = users
	gin := apis.NewRouter()

	// PORTの取得
	port, err := pomodork_util.GetRequiredEnv(pomodork_constant.SERVER_PORT, "string")
	if err != nil {
		serverFatalErr := pomodork_error.ServeFatalError{}
		logger.Error(context.Background(), serverFatalErr.Code(), serverFatalErr.Error())
		return
	}
	gin.Run(":" + port.(string))
}
