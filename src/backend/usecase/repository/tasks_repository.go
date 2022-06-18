package repository

import (
	"context"
	"pomodork-backend/entities"
)

// TasksRepository タスクのリポジトリ
type TasksRepository interface {
	GetTaskInterface(ctx context.Context, userId string, task entities.Task) (string, entities.Task)
	GetTasksInterface(ctx context.Context, userId string, task entities.Task) (string, entities.Task)
	CreateTaskInterface(ctx context.Context, userId string, task entities.Task) (string, entities.Task)
	UpdateTaskInterface(ctx context.Context, userId string, task entities.Task) (string, entities.Task)
}
