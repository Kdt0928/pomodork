package database

import "context"

type SqlHandlerInterface interface {
	GetMaxUserId(ctx context.Context, userAccount *UserAccount) error
	UserCount(ctx context.Context, userAccount *UserAccount) (int64, error)
	CreateUser(ctx context.Context, userAccount *UserAccount) error
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
