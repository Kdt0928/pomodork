package repository

import (
	"context"
	"pomodork-backend/entities"
)

type UserAccountsRepository interface {
	CreateUserInterface(ctx context.Context) string
	GetUserInterface(ctx context.Context, userId string) (entities.UserAccount, string)
	DeleteUserInterface(ctx context.Context, userId string) string
	UpdateUserAccountInterface(ctx context.Context, user entities.UserAccount) string
	GetUserAccountInterface(ctx context.Context, userId string) (entities.UserAccount, string)
}
