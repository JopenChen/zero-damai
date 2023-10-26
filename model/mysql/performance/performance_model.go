package performance

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PerformanceModel = (*customPerformanceModel)(nil)

type (
	// PerformanceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPerformanceModel.
	PerformanceModel interface {
		performanceModel
	}

	customPerformanceModel struct {
		*defaultPerformanceModel
	}
)

// NewPerformanceModel returns a model for the database table.
func NewPerformanceModel(conn sqlx.SqlConn) PerformanceModel {
	return &customPerformanceModel{
		defaultPerformanceModel: newPerformanceModel(conn),
	}
}
