package user

import (
	"context"
	"github.com/JopenChen/zero-damai/service/client/service"

	"github.com/JopenChen/zero-damai/gateway/internal/svc"
	"github.com/JopenChen/zero-damai/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddLogic) UserAdd(req *types.UserAddReq) (resp *types.UserAddResp, err error) {
	resp = new(types.UserAddResp)

	rpcResp, err := l.svcCtx.ServiceRpc.UserAdd(l.ctx, &service.UserAddReq{
		Name:     req.Name,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Mobile:   req.Mobile,
		Password: req.Password,
		Mail:     req.Mail,
		Identity: req.Identity,
		Gender:   req.Gender,
		Nation:   req.Nation,
		Birthday: req.Birthday,
		Address:  req.Address,
	})
	if err != nil {
		l.Logger.Errorf("API UserAdd -> RPC UserAdd error: %v", err)
		return
	}

	// Return
	resp.Id = rpcResp.Id
	return
}
