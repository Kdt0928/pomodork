package gorm

import (
	"errors"
	"net"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"

	pomodork_error "pomodork-backend/message/error"
)

const (
	DUPLICATE_ERROR = 1062
)

// SqlErrorConverter エラー変換
func (g *GormHandler) SqlErrorConverter(err error) error {

	if errors.Is(gorm.ErrRecordNotFound, err) {
		// 対象レコードなし
		return &pomodork_error.NotFoundError{}
	}

	if mysqlErr, isErr := err.(*mysql.MySQLError); isErr {
		switch mysqlErr.Number {
		case DUPLICATE_ERROR:
			// キー重複
			return &pomodork_error.DuplicateError{}
		}
	}

	if _, isErr := err.(*net.OpError); isErr {
		// SQL実行エラー
		return &pomodork_error.SQLExecError{}
	}

	// SQL実行エラー
	return &pomodork_error.SQLExecError{}

}
