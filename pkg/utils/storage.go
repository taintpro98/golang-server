package utils

import (
	"context"

	"gorm.io/gorm"
)

const (
	CtxTxKey = "Tx"
	MaxLimit = 100
)

func getTX(ctx context.Context, db *gorm.DB) *gorm.DB {
	tx, ok := ctx.Value(CtxTxKey).(*gorm.DB)
	if !ok {
		tx = db.WithContext(ctx)
	}
	return tx
}

func Paginate(offset, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if limit > MaxLimit {
			limit = MaxLimit
		}

		if limit < 1 {
			limit = 20
		}
		return db.Offset(offset).Limit(limit)
	}
}
