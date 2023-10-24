package svc

import (
	"github.com/JopenChen/zero-damai/model/mysql/performance"
	"github.com/JopenChen/zero-damai/model/mysql/performance_seat"
	"github.com/JopenChen/zero-damai/model/mysql/performance_session"
	"github.com/JopenChen/zero-damai/model/mysql/user"
	"github.com/JopenChen/zero-damai/service/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config                  config.Config
	UserModel               user.UserModel
	PerformanceModel        performance.PerformanceModel
	PerformanceSessionModel performance_session.PerformanceSessionModel
	PerformanceSeatModel    performance_seat.PerformanceSeatModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 创建 Mysql 数据库连接
	mysqlConn := sqlx.NewMysql(c.Mysql.Datasource)

	return &ServiceContext{
		Config:                  c,
		UserModel:               user.NewUserModel(mysqlConn),
		PerformanceModel:        performance.NewPerformanceModel(mysqlConn),
		PerformanceSessionModel: performance_session.NewPerformanceSessionModel(mysqlConn),
		PerformanceSeatModel:    performance_seat.NewPerformanceSeatModel(mysqlConn),
	}
}
