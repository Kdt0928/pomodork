package controller

import (
	"context"
	"pomodork-backend/usecase"
)

// UserAccountControllerInterface Controllerのインターフェース
type UserAccountControllerInterface interface {
	CreateUser(ctx context.Context) (string, error)
}

// UserAccountRepositories UseCaseのインターフェース群
type UserAccountRepositories struct {
	UserAccountUseCase usecase.UserAccountUseCaseInterface
}

// NewUserAccountRepositories UserAccountRepositoriesの生成
func NewUserAccountRepositories() *UserAccountRepositories {
	useCase := new(UserAccountRepositories)
	return useCase
}

// CreateUser ユーザ作成
func (u *UserAccountRepositories) CreateUser(ctx context.Context) (string, error) {
	return u.UserAccountUseCase.CreateUser(ctx)
}
