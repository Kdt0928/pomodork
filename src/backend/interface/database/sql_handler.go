package database

import "context"

type SqlHandlerInterface interface {
	Create(ctx context.Context, model interface{}) error
	Find(ctx context.Context, model interface{}) error
	Latest(ctx context.Context, model interface{}) error
	Count(ctx context.Context, model interface{}) (int64, error)
	SqlErrorConverter(error) error
}

type TransactionHandler interface {
	Create(model interface{}) interface{}
	// First(model interface{}) (interface{}, error)
	// Find(model interface{}) (interface{}, error)
	// Save(model interface{}) (interface{}, error)
	// Delete(model interface{}) (interface{}, error)
	// Where(query interface{}, args ...interface{}) (interface{}, error)
	// Order(value interface{}) (interface{}, error)
	// Limit(limit int) (interface{}, error)
	// Commit() (interface{}, error)
	// RollBack() (interface{}, error)
}

// SqlRepositories Interface層で利用するinfra層のInterface群
type SqlRepositories struct {
	SqlHandler SqlHandlerInterface
}

// NewSqlRepositories sqlRepositoriesの生成
func NewSqlRepositories() *SqlRepositories {
	sqlHandler := new(SqlRepositories)
	return sqlHandler
}
