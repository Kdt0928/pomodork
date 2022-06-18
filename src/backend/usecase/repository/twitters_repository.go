package repository

import (
	"context"
	"pomodork-backend/entities"
)

// TwittersRepository Twitter連携リポジトリ
type TwittersRepository interface {
	CreateLinkTwitterInterface(ctx context.Context, twitter entities.Twitter) (string, string)
	DeleteLinkTwitterInterface(ctx context.Context, twitter entities.Twitter) string
}
