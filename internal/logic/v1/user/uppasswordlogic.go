package user

import (
	"context"

	"cheddar/internal/svc"
	"cheddar/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpPasswordLogic {
	return &UpPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpPasswordLogic) UpPassword(req *types.UpPasswordReq) error {
	// todo: add your logic here and delete this line

	return nil
}
