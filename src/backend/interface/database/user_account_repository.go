package database

import (
	"context"
	"pomodork-backend/entities"
	pomodork_error "pomodork-backend/message/error"
	pomodork_util "pomodork-backend/util"
	"time"
)

// UserAccount ユーザアカウント
type UserAccount struct {
	UserId         string
	MailAddress    string
	Password       string
	AccountName    string
	LinkTwitterId  string
	LastUpdateDate time.Time
	LastUpdateUser string
}

// CreateUser ユーザ作成
func (u *SqlRepositories) CreateUser(ctx context.Context, userId string) error {
	logger := pomodork_util.NewLogger()
	userAccount := &UserAccount{
		UserId:         userId,
		LastUpdateDate: time.Now(),
		LastUpdateUser: userId,
	}

	if err := u.SqlHandler.Create(ctx, userAccount); err != nil {
		sqlErr := &pomodork_error.SQLExecError{TableName: "ユーザアカウント", Err: err}
		logger.Error(ctx, sqlErr.Code(), sqlErr.Error())

		err := u.SqlHandler.SqlErrorConverter(err)
		SqlErrorHandler(ctx, "ユーザアカウント", err, userId)
		return err
	}
	return nil
}

// IsUserExists ユーザ存在チェック
func (u *SqlRepositories) IsUserExists(ctx context.Context) (bool, error) {
	logger := pomodork_util.NewLogger()
	userAccount := UserAccount{}
	userExists := true

	if count, err := u.SqlHandler.Count(ctx, &userAccount); count == 0 || err != nil {
		userExists = false
		if err != nil {
			sqlErr := &pomodork_error.SQLExecError{TableName: "ユーザアカウント", Err: err}
			logger.Error(ctx, sqlErr.Code(), sqlErr.Error())

			err := u.SqlHandler.SqlErrorConverter(err)
			SqlErrorHandler(ctx, "ユーザアカウント", err, "")
			return userExists, sqlErr
		}
	}
	return userExists, nil
}

func (u *SqlRepositories) GetUser(ctx context.Context, userId string) (entities.UserAccount, string) {
	return entities.UserAccount{}, ""
}

// GetMaxUserId 最新のユーザID取得
func (u *SqlRepositories) GetMaxUserId(ctx context.Context) (string, error) {
	logger := pomodork_util.NewLogger()
	userAccount := &UserAccount{}

	if err := u.SqlHandler.Latest(ctx, userAccount); err != nil {
		sqlErr := &pomodork_error.SQLExecError{TableName: "ユーザアカウント", Err: err}
		logger.Error(ctx, sqlErr.Code(), sqlErr.Error())

		err := u.SqlHandler.SqlErrorConverter(err)
		SqlErrorHandler(ctx, "ユーザアカウント", err, "")
		return "", err
	}
	return userAccount.UserId, nil
}

func (u *SqlRepositories) DeleteUser(ctx context.Context, userId string) string {
	return ""
}

func (u *SqlRepositories) UpdateUserAccount(ctx context.Context, user entities.UserAccount) string {
	return ""
}

func (u *SqlRepositories) GetUserAccount(ctx context.Context, userId string) (entities.UserAccount, string) {
	return entities.UserAccount{}, ""
}
