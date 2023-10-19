package db_handle

import (
	"context"
	"gorm.io/gorm"
)

type DBAction struct {
	ctx context.Context
	db  *gorm.DB
}

func CreateDBHandle(ctx context.Context, db *gorm.DB) *DBAction {
	return &DBAction{
		ctx: ctx,
		db:  db,
	}
}
