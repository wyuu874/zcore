package gormx

import (
	"gorm.io/gorm"
	"time"
)

// Model 基础模型
type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`              // 主键ID
	CreatedAt time.Time      `gorm:"not null" json:"created_at"`        // 创建时间
	UpdatedAt time.Time      `gorm:"not null" json:"updated_at"`        // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"` // 删除时间（软删除）
}

// TimeModel 时间模型（不带软删除）
type TimeModel struct {
	ID        uint      `gorm:"primarykey" json:"id"`       // 主键ID
	CreatedAt time.Time `gorm:"not null" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"` // 更新时间
}

// Pagination 分页参数
type Pagination struct {
	Page     int `json:"page" form:"page"`           // 当前页码
	PageSize int `json:"page_size" form:"page_size"` // 每页数量
}

// PageResult 分页结果
type PageResult struct {
	List     interface{} `json:"list"`      // 数据列表
	Page     int         `json:"page"`      // 当前页码
	PageSize int         `json:"page_size"` // 每页数量
	Total    int64       `json:"total"`     // 总数据量
}

// Paginate 分页查询
func Paginate(p *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if p.Page <= 0 {
			p.Page = 1
		}
		if p.PageSize <= 0 {
			p.PageSize = 10
		}
		if p.PageSize > 100 {
			p.PageSize = 100
		}

		offset := (p.Page - 1) * p.PageSize
		return db.Offset(offset).Limit(p.PageSize)
	}
}

// GetPage 获取分页结果
func GetPage(db *gorm.DB, p *Pagination, data interface{}) (*PageResult, error) {
	var total int64
	// 克隆一个新的查询，这样不会影响原有的查询条件
	countDB := db.Session(&gorm.Session{})
	err := countDB.Count(&total).Error
	if err != nil {
		return nil, err
	}

	err = db.Scopes(Paginate(p)).Find(data).Error
	if err != nil {
		return nil, err
	}

	return &PageResult{
		List:     data,
		Page:     p.Page,
		PageSize: p.PageSize,
		Total:    total,
	}, nil
}
