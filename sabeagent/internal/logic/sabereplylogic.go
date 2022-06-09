package logic

import (
	"context"

	"github.com/moyrne/sabe-reply/sabeagent/internal/svc"
	"github.com/moyrne/sabe-reply/sabeagent/sabeagent"

	"github.com/zeromicro/go-zero/core/logx"
)

type SabeReplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSabeReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SabeReplyLogic {
	return &SabeReplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SabeReplyLogic) SabeReply(in *sabeagent.SabeReplyRequest) (*sabeagent.SabeReplyResponse, error) {
	// todo: add your logic here and delete this line

	return &sabeagent.SabeReplyResponse{}, nil
}
