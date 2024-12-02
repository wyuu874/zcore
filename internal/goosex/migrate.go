package goosex

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
)

// Migrator 迁移器
type Migrator struct {
	db        *sql.DB
	directory string
}

// New 创建迁移器
func New(dsn string, directory string) (*Migrator, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败: %v", err)
	}

	if err := goose.SetDialect("mysql"); err != nil {
		return nil, fmt.Errorf("设置数据库类型失败: %v", err)
	}

	return &Migrator{
		db:        db,
		directory: directory,
	}, nil
}

// Up 执行向上迁移
func (m *Migrator) Up() error {
	return goose.Up(m.db, m.directory)
}

// Down 回滚最后一次迁移
func (m *Migrator) Down() error {
	return goose.Down(m.db, m.directory)
}

// Status 获取迁移状态
func (m *Migrator) Status() error {
	return goose.Status(m.db, m.directory)
}

// Create 创建新的迁移文件
func (m *Migrator) Create(name string) error {
	return goose.Create(m.db, m.directory, name, "sql")
}

// Close 关闭数据库连接
func (m *Migrator) Close() error {
	return m.db.Close()
}
