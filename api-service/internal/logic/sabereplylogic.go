package logic

import (
	"context"
	"time"

	"github.com/moyrne/sabe-reply/api-service/internal/svc"
	"github.com/moyrne/sabe-reply/api-service/internal/types"
	"github.com/moyrne/sabe-reply/api-service/model"
	"github.com/moyrne/sabe-reply/sabeagent/sabeagent"

	"github.com/zeromicro/go-zero/core/logx"
)

type SabeReplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSabeReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SabeReplyLogic {
	return &SabeReplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SabeReplyLogic) SabeReply(req *types.SabeReplyRequest) (resp *types.SabeReplyResponse, err error) {
	if _, err := l.svcCtx.Model.Insert(l.ctx, &model.ChatRecord{
		CreatedAt:  time.Now(),
		Kind:       req.Kind,
		Sender:     req.Sender,
		Receiver:   req.Receiver,
		Content:    req.Content,
		RawContent: req.RawContent,
	}); err != nil {
		l.Logger.Error("insert chat record failed", "err", err)
	}

	// TODO 记录最近的聊天记录

	reply, err := l.svcCtx.SabeAgent.SabeReply(l.ctx, &sabeagent.SabeReplyRequest{
		Sender:  req.Sender,
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}

	return &types.SabeReplyResponse{Reply: reply.GetReply()}, nil
}
