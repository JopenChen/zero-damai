package performance_session

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PerformanceSessionModel = (*customPerformanceSessionModel)(nil)

type (
	// PerformanceSessionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPerformanceSessionModel.
	PerformanceSessionModel interface {
		performanceSessionModel
	}

	customPerformanceSessionModel struct {
		*defaultPerformanceSessionModel
	}
)

// NewPerformanceSessionModel returns a model for the database table.
func NewPerformanceSessionModel(conn sqlx.SqlConn) PerformanceSessionModel {
	return &customPerformanceSessionModel{
		defaultPerformanceSessionModel: newPerformanceSessionModel(conn),
	}
}
