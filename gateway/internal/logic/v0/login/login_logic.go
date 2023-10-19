package login

import (
	"context"
	"github.com/JopenChen/zero-damai/service/client/service"

	"github.com/JopenChen/zero-damai/gateway/internal/svc"
	"github.com/JopenChen/zero-damai/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	resp = new(types.LoginResp)
	
	rpcResp, err := l.svcCtx.ServiceRpc.Login(l.ctx, &service.LoginReq{
		Mobile:    req.Mobile,
		Password:  req.Password,
		Code:      req.Code,
		LoginType: req.LoginType,
	})
	if err != nil {
		l.Logger.Errorf("Gateway Login -> Service Login error: %v", err)
		return
	}

	// Return
	resp.ID = rpcResp.ID
	resp.Name = rpcResp.Name
	resp.Token = rpcResp.Token
	resp.ExpireAt = rpcResp.ExpireAt
	return
}
