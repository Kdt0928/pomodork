package repository

import "context"

type TransactionRepository interface {
	Begin(ctx context.Context)
	RollBack(ctx context.Context)
	Commit(ctx context.Context)
}
