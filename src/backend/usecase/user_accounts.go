package usecase

import (
	"context"
	"pomodork-backend/entities"
)

// CreateUser ユーザ登録
func CreateUser(ctx context.Context) string {
	// ユーザID取得(MAX)
	// ユーザ登録
	return ""
}

// GetUser ユーザ取得
func GetUser(ctx context.Context, userId string) (entities.UserAccount, entities.Twitter) {
	return entities.UserAccount{}, entities.Twitter{}
}

// DeleteUser ユーザ削除
func DeleteUser(ctx context.Context, userId string) string {
	return ""
}

// UpdateUserAccount ユーザアカウント更新
func UpdateUserAccount(ctx context.Context, user entities.UserAccount) string {
	return ""
}

// GetUserAccount ユーザアカウント取得
func GetUserAccount(ctx context.Context, userId string) (entities.UserAccount, entities.Twitter) {
	return entities.UserAccount{}, entities.Twitter{}
}
