package servicelogic

import (
	"context"
	"errors"
	"github.com/JopenChen/zero-damai/common/errx"
	"github.com/JopenChen/zero-damai/common/global"
	"github.com/JopenChen/zero-damai/common/jwt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"service/internal/svc"
	"service/service_pb"
	"time"

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
			return nil, errx.NewErrCode(errx.UserNotRegisterError)
		}
		return
	}

	// 检查用户是否被禁用
	if userItem.Status == global.StatusForbidden {
		return nil, errx.NewErrCode(errx.UserForbiddenError)
	}

	// 判断登录方式
	switch in.LoginType {
	case global.LoginCodeType:
		// TODO
	case global.LoginPasswordType:
		if in.Password != userItem.Password {
			err = errx.NewErrCode(errx.LoginPasswordIncorrectError)
			return
		}

	case global.LoginZhiFubaoType:
		// TODO
	case global.LoginWeixinType:
		// TODO
	case global.LoginXinlangType:
		// TODO
	case global.LoginQQType:
		// TODO
	default:
		err = errx.NewErrCode(errx.LoginTypeNotSupportError)
		return
	}

	// 获取 JWT Token
	tn := time.Now().Unix()
	claims := make(map[string]interface{})
	claims["id"] = userItem.Id
	claims["name"] = userItem.Name
	token, err := jwt.GenToken(tn, l.svcCtx.Config.Jwt.Secret, claims, l.svcCtx.Config.Jwt.Expire)
	if err != nil {
		err = errx.NewErrCode(errx.UserLoginFailError)
		return
	}

	// Return
	resp.ID = userItem.Id
	resp.Name = userItem.Name
	resp.Token = token
	resp.ExpireAt = tn + l.svcCtx.Config.Jwt.Expire
	return
}
