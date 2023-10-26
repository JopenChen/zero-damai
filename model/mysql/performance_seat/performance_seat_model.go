package performance_seat

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PerformanceSeatModel = (*customPerformanceSeatModel)(nil)

type (
	// PerformanceSeatModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPerformanceSeatModel.
	PerformanceSeatModel interface {
		performanceSeatModel
	}

	customPerformanceSeatModel struct {
		*defaultPerformanceSeatModel
	}
)

// NewPerformanceSeatModel returns a model for the database table.
func NewPerformanceSeatModel(conn sqlx.SqlConn) PerformanceSeatModel {
	return &customPerformanceSeatModel{
		defaultPerformanceSeatModel: newPerformanceSeatModel(conn),
	}
}
