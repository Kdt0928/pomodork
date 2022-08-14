package util

import (
	"context"
	"fmt"
	"os"
	pomodork_constant "pomodork-backend/constant"
	pomodork_error "pomodork-backend/message/error"
	pomodork_info "pomodork-backend/message/info"
	"strconv"
)

// GetEnv osの環境変数を取得し、変換する
func GetEnv(key string, envType string) (env interface{}, err error) {

	logger := NewLogger()
	ctx := context.Background()
	ctx = context.WithValue(ctx, pomodork_constant.TRANSACTION_ID, "")

	switch envType {
	case "string":
		// 存在チェック
		env = os.Getenv(key)
		if env == "" {
			err = &pomodork_error.EnvNotFoundError{Key: key}
			logger.Error(ctx, err.Error())
		}
	case "int":
		// 存在チェック
		env = os.Getenv(key)
		if env == "" {
			err = &pomodork_error.EnvNotFoundError{Key: key}
			logger.Error(ctx, err.Error())
		}
		// 型変換チェック
		if env, err = strconv.Atoi(env.(string)); err != nil {
			err = &pomodork_error.EnvConvertError{Key: key, BeforeEnv: env}
			logger.Error(ctx, err.Error())
		}
	default:
		// 存在チェック
		env = os.Getenv(key)
		if env == "" {
			err = &pomodork_error.EnvNotFoundError{Key: key}
			logger.Error(ctx, err.Error())
		}
	}

	if err == nil {
		logger.Info(ctx, fmt.Sprintf(pomodork_info.MSG_I00001, key, env))
	}
	return env, err
}
