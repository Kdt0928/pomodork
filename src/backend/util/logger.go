package util

import (
	"context"
	"os"

	"go.uber.org/zap"

	pomodork_constant "pomodork-backend/constant"
)

type Logger struct {
	zap *zap.Logger
}

// NewLogger Loggerの生成
func NewLogger() *Logger {
	// デバッグモードの取得
	debugMode := os.Getenv(pomodork_constant.DEBUG_MODE)
	if debugMode == "" || debugMode == "1" {
		zap, _ := zap.NewDevelopment()
		return &Logger{zap: zap}
	} else {
		zap, _ := zap.NewProduction()
		return &Logger{zap: zap}
	}
}

func (logger *Logger) Info(ctx context.Context, msg string) {
	logger.zap.Info(msg, zap.String("SERVER_NAME", pomodork_constant.SERVER_NAME), zap.String("TRANSACTION_ID", ctx.Value(pomodork_constant.TRANSACTION_ID).(string)))
}

func (logger *Logger) Warn(ctx context.Context, msg string) {
	logger.zap.Warn(msg, zap.String("SERVER_NAME", pomodork_constant.SERVER_NAME), zap.String("TRANSACTION_ID", ctx.Value(pomodork_constant.TRANSACTION_ID).(string)))
}

func (logger *Logger) Error(ctx context.Context, msg string) {
	logger.zap.Error(msg, zap.String("SERVER_NAME", pomodork_constant.SERVER_NAME), zap.String("TRANSACTION_ID", ctx.Value(pomodork_constant.TRANSACTION_ID).(string)))
}
