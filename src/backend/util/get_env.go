package util

import (
	"context"
	"fmt"
	"os"
	pomodork_error "pomodork-backend/message/error"
	pomodork_info "pomodork-backend/message/info"
	"strconv"
)

// GetRequiredEnv 必須環境変数を取得、変換
func GetRequiredEnv(key string, envType string) (interface{}, error) {
	logger := NewLogger()

	// 存在チェック
	var env interface{}
	env = os.Getenv(key)
	if env == "" {
		err := &pomodork_error.EnvNotFoundError{Key: key}
		logger.Error(context.Background(), err.Code(), err.Error())
		return "", err
	}

	switch envType {
	case "int":
		// 型変換チェック
		intEnv, err := strconv.Atoi(env.(string))
		if err != nil {
			err := &pomodork_error.EnvConvertError{Key: key, BeforeEnv: env.(string)}
			logger.Error(context.Background(), err.Code(), err.Error())
			return "", err
		}
		env = intEnv
	}
	logger.Info(context.Background(), pomodork_info.CODE_I00001, fmt.Sprintf(pomodork_info.MSG_I00001, key, env))
	return env, nil
}

// GetNonRequiredEnv 任意環境変数を取得、変換
func GetNonRequiredEnv(key string, envType string, defaultValue interface{}) (interface{}, error) {
	logger := NewLogger()

	// 存在チェック
	env := interface{}(os.Getenv(key))
	if env == "" {
		err := &pomodork_error.EnvNotFoundWarn{Key: key, DefaultValue: defaultValue}
		logger.Warn(context.Background(), err.Code(), err.Error())
		// デフォルト値を設定
		env = defaultValue
	}

	switch envType {
	case "int":
		// 型変換チェック
		if intEnv, err := strconv.Atoi(env.(string)); err != nil {
			err := &pomodork_error.EnvConvertWarn{Key: key, BeforeEnv: env.(string), DefaultValue: defaultValue}
			logger.Warn(context.Background(), err.Code(), err.Error())
			// デフォルト値を設定
			env = intEnv
		}
	}
	logger.Info(context.Background(), pomodork_info.CODE_I00001, fmt.Sprintf(pomodork_info.MSG_I00001, key, env))
	return env, nil
}
