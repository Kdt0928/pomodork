package controller

import (
	"context"
	"pomodork-backend/usecase"
)

// User ユーザ情報
type User struct {
	UserId        string // ユーザID
	MailAddress   string // メールアドレス
	Password      string // パスワード
	AccountName   string // アカウント名
	TwitterLinkId string // Twitter連携ID
}

// UserAccountController Controllerのインターフェース
type UserAccountController interface {
	CreateUser(ctx context.Context) (string, error)
	GetUser(ctx context.Context, userId string) (User, error)
}

// UserAccountDomains UseCaseのインターフェース群
type UserAccountDomains struct {
	Domain usecase.UserAccountDomain
}

// NewUserAccountDomains UserAccountDomainsの生成
func NewUserAccountRepositories() *UserAccountDomains {
	useCase := new(UserAccountDomains)
	return useCase
}

// CreateUser ユーザ作成
func (u *UserAccountDomains) CreateUser(ctx context.Context) (string, error) {
	userId, err := u.Domain.CreateUser(ctx)
	if err != nil {
		return "", err
	}
	return userId, nil
}

// GetUser ユーザ取得
func (u *UserAccountDomains) GetUser(ctx context.Context, userId string) (User, error) {
	return User{}, nil
}
