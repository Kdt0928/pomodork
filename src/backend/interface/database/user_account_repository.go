package database

import (
	"context"
	pomodork_constant "pomodork-backend/constant"
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

// UserAccountSqlHandlerRepositories Interface層で利用するinfra層のInterface群
type UserAccountSqlHandlerRepositories struct {
	SqlHandlerRepository SqlHandlerInterface
}

// NewUserAccountSqlHandlerRepositories UserAccountSqlHandlerRepositoriesの生成
func NewUserAccountSqlHandlerRepositories() *UserAccountSqlHandlerRepositories {
	sqlHandler := new(UserAccountSqlHandlerRepositories)
	return sqlHandler
}

// CreateUser ユーザ作成
func (handler *UserAccountSqlHandlerRepositories) CreateUser(ctx context.Context, userId string) error {
	logger := pomodork_util.NewLogger()
	lastUpdateDate := ctx.Value(pomodork_constant.USER_ID)
	userAccount := &UserAccount{UserId: userId, LastUpdateDate: time.Now(), LastUpdateUser: lastUpdateDate.(string)}

	if err := handler.SqlHandlerRepository.CreateUser(ctx, userAccount); err != nil {
		sqlErr := &pomodork_error.SQLExecError{TableName: "ユーザアカウント", Err: err}
		logger.Error(ctx, sqlErr.Error())
		return sqlErr
	}
	return nil
}

func (handler *UserAccountSqlHandlerRepositories) GetUser(ctx context.Context, userId string) (entities.UserAccount, string) {
	return entities.UserAccount{}, ""
}

// GetMaxUserId 最新のユーザID取得
func (handler *UserAccountSqlHandlerRepositories) GetMaxUserId(ctx context.Context) (string, error) {
	logger := pomodork_util.NewLogger()
	userAccount := &UserAccount{}

	if err := handler.SqlHandlerRepository.GetMaxUserId(ctx, userAccount); err != nil {
		sqlErr := &pomodork_error.SQLExecError{TableName: "ユーザアカウント", Err: err}
		logger.Error(ctx, sqlErr.Error())
		return "", err
	}
	return userAccount.UserId, nil
}

// IsUserExists ユーザ存在チェック
func (handler *UserAccountSqlHandlerRepositories) IsUserExists(ctx context.Context) (bool, error) {
	logger := pomodork_util.NewLogger()
	userAccount := &UserAccount{}
	userExists := false

	if count, err := handler.SqlHandlerRepository.UserCount(ctx, userAccount); count == 0 || err != nil {
		userExists = true
		if err != nil {
			sqlErr := &pomodork_error.SQLExecError{TableName: "ユーザアカウント", Err: err}
			logger.Error(ctx, sqlErr.Error())
			return userExists, sqlErr
		}
		return userExists, nil
	}
	return userExists, nil
}

func (handler *UserAccountSqlHandlerRepositories) DeleteUser(ctx context.Context, userId string) string {
	return ""
}

func (handler *UserAccountSqlHandlerRepositories) UpdateUserAccount(ctx context.Context, user entities.UserAccount) string {
	return ""
}

func (handler *UserAccountSqlHandlerRepositories) GetUserAccount(ctx context.Context, userId string) (entities.UserAccount, string) {
	return entities.UserAccount{}, ""
}
