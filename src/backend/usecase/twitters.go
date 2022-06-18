package usecase

import (
	"context"
	"pomodork-backend/entities"
)

// CreateLinkTwitter Twitter連携登録
func CreateLinkTwitter(ctx context.Context, twitter entities.Twitter) (string, string) {
	// ユーザ存在チェック
	// Twitter連携登録
	return "", ""
}

// DeleteLinkTwitter Twitter連携削除
func DeleteLinkTwitter(ctx context.Context, twitter entities.Twitter) string {
	return ""
}
