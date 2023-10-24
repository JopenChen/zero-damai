package servicelogic

import (
	"context"
	"errors"
	"github.com/JopenChen/zero-damai/common/errx"
	"github.com/JopenChen/zero-damai/common/global"
	"github.com/JopenChen/zero-damai/model/mysql/user"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"

	"github.com/JopenChen/zero-damai/service/internal/svc"
	"github.com/JopenChen/zero-damai/service/service_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserAdd 用户注册
func (l *UserAddLogic) UserAdd(in *service_pb.UserAddReq) (resp *service_pb.UserAddResp, err error) {
	resp = new(service_pb.UserAddResp)
	userRowBuilder := l.svcCtx.UserModel.RowBuilder().
		Where(squirrel.Eq{"mobile": in.Mobile})
	_, err = l.svcCtx.UserModel.GetByCondition(l.ctx, userRowBuilder)
	if err != nil {
		if errors.As(err, &sqlx.ErrNotFound) {
			// DO Nothing
			err = nil
		} else {
			l.Logger.Errorf("UserModel.GetByCondition(l.ctx, userRowBuilder) error: %v", err)
			err = errx.NewErrCode(errx.DbError)
			return
		}
	} else {
		err = errx.NewErrCode(errx.UserAlreadyExistError)
		return
	}

	insertResult, err := l.svcCtx.UserModel.Insert(l.ctx, &user.User{
		Name:     in.Name,
		Nickname: in.Nickname,
		Avatar:   in.Avatar,
		Mobile:   in.Mobile,
		Password: in.Password,
		Mail:     in.Mail,
		Identity: in.Identity,
		Gender:   in.Gender,
		Nation:   in.Nation,
		Birthday: in.Birthday,
		Address:  in.Address,
		Audience: global.EmptyArrayString,
		Status:   global.StatusActive,
		LoginAt:  time.Now(),
	})
	if err != nil {
		l.Logger.Errorf("UserModel.Insert error: %v", err)
		err = errx.NewErrCode(errx.DbError)
		return
	}

	id, err := insertResult.LastInsertId()
	if err != nil {
		l.Logger.Errorf("insertResult.LastInsertId() error: %v", err)
		err = errx.NewErrCode(errx.ServerCommonError)
		return
	}

	// Return
	resp.Id = id
	return
}
