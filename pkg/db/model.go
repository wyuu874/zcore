package db

import (
	"github.com/wyuu874/zcore/internal/gormx"
)

// Model 基础模型
type Model = gormx.Model

// TimeModel 时间模型（不带软删除）
type TimeModel = gormx.TimeModel

// Pagination 分页参数
type Pagination = gormx.Pagination

// PageResult 分页结果
type PageResult = gormx.PageResult

// Paginate 分页查询
func Paginate(p *Pagination) func(db *gormx.GormEngine) *gormx.GormEngine {
	return gormx.Paginate(p)
}

// GetPage 获取分页结果
func GetPage(db *gormx.GormEngine, p *Pagination, data interface{}) (*PageResult, error) {
	return gormx.GetPage(db, p, data)
}
