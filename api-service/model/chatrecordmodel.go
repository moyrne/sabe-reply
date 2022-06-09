package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ChatRecordModel = (*customChatRecordModel)(nil)

type (
	// ChatRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChatRecordModel.
	ChatRecordModel interface {
		chatRecordModel
	}

	customChatRecordModel struct {
		*defaultChatRecordModel
	}
)

// NewChatRecordModel returns a model for the database table.
func NewChatRecordModel(conn sqlx.SqlConn, c cache.CacheConf) ChatRecordModel {
	return &customChatRecordModel{
		defaultChatRecordModel: newChatRecordModel(conn, c),
	}
}
