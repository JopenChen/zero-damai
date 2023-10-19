package servicelogic

import (
	"context"
	"errors"
	"github.com/JopenChen/zero-damai/common/global"
	"github.com/JopenChen/zero-damai/service/internal/svc"
	"github.com/JopenChen/zero-damai/service/service_pb"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *service_pb.LoginReq) (resp *service_pb.LoginResp, err error) {
	resp = new(service_pb.LoginResp)

	// 校验用户是否已注册
	userRowBuilder := l.svcCtx.UserModel.RowBuilder().
		Where(squirrel.Eq{"mobile": in.Mobile})
	userItem, err := l.svcCtx.UserModel.GetByCondition(l.ctx, userRowBuilder)
	if err != nil {
		l.Logger.Errorf("UserModel.GetByCondition(l.ctx, userRowBuilder) error: %v", err)
		if errors.Is(err, sqlc.ErrNotFound) {
			// TODO 完善对应响应
		}
		return
	}

	// 检查用户是否被禁用
	if userItem.Status == global.StatusForbidden {
		return
	}

	// 获取 JWT Token

	// Return
	resp.ID = userItem.Id
	resp.Name = userItem.Name
	resp.Token = ""
	resp.ExpireAt = 13564742453
	return
}
