package svc

import (
	"github.com/moyrne/sabe-reply/api-service/internal/config"
	"github.com/moyrne/sabe-reply/api-service/model"
	"github.com/moyrne/sabe-reply/sabeagent/sabeagent"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Model     model.ChatRecordModel
	SabeAgent sabeagent.SabeAgent
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Model:     model.NewChatRecordModel(sqlx.NewMysql(c.DataSource), c.Cache),
		SabeAgent: sabeagent.NewSabeAgent(zrpc.MustNewClient(c.SabeAgent)),
	}
}
