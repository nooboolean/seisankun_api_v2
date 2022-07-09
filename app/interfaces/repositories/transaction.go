package repositories

import (
	"context"

	"github.com/nooboolean/seisankun_api_v2/transaction"
	"gorm.io/gorm"
)

type contextKey struct{}

var txKey = contextKey{}

type tx struct {
	db *gorm.DB
}

func NewTransaction(db *gorm.DB) transaction.Transaction {
	return &tx{db: db}
}

func (t *tx) DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	tx := t.db.Begin()
	ctx = context.WithValue(ctx, &txKey, tx)
	v, err := f(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return v, nil
}

func GetTx(ctx context.Context) (*gorm.DB, bool) {
	tx, ok := ctx.Value(&txKey).(*gorm.DB)
	return tx, ok
}
