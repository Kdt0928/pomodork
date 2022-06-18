package usecase

import (
	"context"
	"pomodork-backend/entities"
)

// CreateTask タスク登録
func CreateTask(ctx context.Context, userId string, task entities.Task) (string, entities.Task) {
	// ユーザ存在チェック
	// タスク登録
	return "", entities.Task{}
}

// UpdateTask タスク更新
func UpdateTask(ctx context.Context, userId string, task entities.Task) (string, entities.Task) {
	// ユーザ存在チェック
	// 更新対象のタスク存在チェック
	// タスク更新
	return "", entities.Task{}
}

// GetTasks タスク取得
func GetTasks(ctx context.Context, userId string, task entities.Task) (string, []entities.Task) {
	// ユーザ存在チェック
	// タスク取得
	return "", []entities.Task{}
}
