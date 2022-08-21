package database

import (
	"context"
	pomodork_error "pomodork-backend/message/error"
	pomodork_util "pomodork-backend/util"
)

func SqlErrorHandler(ctx context.Context, tableName string, err error, detail interface{}) {
	logger := pomodork_util.NewLogger()

	switch e := err.(type) {
	case *pomodork_error.NotFoundError:
		e.TableName = tableName
		e.Conditions = detail.(string)
		logger.Warn(ctx, e.Code(), e.Error())
	case *pomodork_error.DuplicateError:
		e.TableName = tableName
		e.Key = detail.(string)
		logger.Warn(ctx, e.Code(), e.Error())
	}
}
