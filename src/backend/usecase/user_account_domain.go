package usecase

import (
	"context"
	pomodork_constant "pomodork-backend/constant"
	"pomodork-backend/entities"
	common "pomodork-backend/usecase/common"
	"pomodork-backend/usecase/repository"
)

/*
	user_account_domain
	ユーザ_アカウントに関する処理を実行するDomain層
	・CreateUser ユーザ登録
	・GetUser ユーザ取得
	・DeleteUser ユーザ削除
	・UpdateUserAccount ユーザアカウント更新
	・GetUserAccount ユーザアカウント取得
*/

// UserAccountUseCaseInterface UseCaseのインターフェイス
type UserAccountUseCaseInterface interface {
	CreateUser(ctx context.Context) (string, error)
	GetUser(ctx context.Context, userId string) (entities.UserAccount, entities.Twitter)
	DeleteUser(ctx context.Context, userId string) string
	UpdateUserAccount(ctx context.Context, user entities.UserAccount) string
	GetUserAccount(ctx context.Context, userId string) (entities.UserAccount, entities.Twitter)
}

// UserAccountRepositories domain層で利用するInterface層のInterface群
type UserAccountRepositories struct {
	UserAccountRepository repository.UserAccountRepository
	TransactionRepository repository.TransactionRepository
}

// NewUserAccountRepositories UserAccountRepositoriesの生成
func NewUserAccountRepositories() *UserAccountRepositories {
	repository := new(UserAccountRepositories)
	return repository
}

// CreateUser ユーザ登録
func (u *UserAccountRepositories) CreateUser(ctx context.Context) (string, error) {
	// ユーザアカウント存在チェック
	userExists, err := u.UserAccountRepository.IsUserExists(ctx)
	if err != nil {
		return "", err
	}

	userId := pomodork_constant.FIRST_USER_ID
	if userExists {
		// ユーザID取得
		previousUserId, err := u.UserAccountRepository.GetMaxUserId(ctx)
		if err != nil {
			return "", err
		}
		//　ユーザID発番
		userId = common.CreateId(previousUserId)
	}

	// ユーザ登録
	err = u.UserAccountRepository.CreateUser(ctx, userId)
	if err != nil {
		return "", err
	}

	return userId, nil
}

// GetUser ユーザ取得
func (u *UserAccountRepositories) GetUser(ctx context.Context, userId string) (entities.UserAccount, entities.Twitter) {
	return entities.UserAccount{}, entities.Twitter{}
}

// DeleteUser ユーザ削除
func (u *UserAccountRepositories) DeleteUser(ctx context.Context, userId string) string {
	return ""
}

// UpdateUserAccount ユーザアカウント更新
func (u *UserAccountRepositories) UpdateUserAccount(ctx context.Context, user entities.UserAccount) string {
	// ユーザ存在チェック
	return ""
}

// GetUserAccount ユーザアカウント取得
func (u *UserAccountRepositories) GetUserAccount(ctx context.Context, userId string) (entities.UserAccount, entities.Twitter) {
	return entities.UserAccount{}, entities.Twitter{}
}
