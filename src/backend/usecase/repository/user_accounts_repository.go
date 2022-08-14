package repository

import (
	"context"
	"pomodork-backend/entities"
)

type UserAccountRepository interface {
	CreateUser(ctx context.Context, firstUserId string) error
	GetUser(ctx context.Context, userId string) (entities.UserAccount, string)
	GetMaxUserId(ctx context.Context) (string, error)
	IsUserExists(ctx context.Context) (bool, error)
	DeleteUser(ctx context.Context, userId string) string
	UpdateUserAccount(ctx context.Context, user entities.UserAccount) string
	GetUserAccount(ctx context.Context, userId string) (entities.UserAccount, string)
}
