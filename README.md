# ZCore Web 框架

ZCore 是一个基于 Gin 的轻量级 Web 框架,提供了简单易用的路由注册和中间件管理功能。

## 特性

- 基于 Gin 框架封装
- 简洁的路由注册方式
- 灵活的中间件管理
- 支持路由分组
- 统一的响应格式
- 内置日志和恢复中间件
- 优雅的错误处理
- 数据库迁移工具

## 快速开始

### 安装

```bash
go get github.com/wyuu874/zcore
```

### 数据库迁移

框架集成了数据库迁移工具，基于 goose 实现。

#### 配置数据库

在配置文件中设置数据库信息：

```toml
[database]
host = "127.0.0.1"
port = "3306"
username = "your_username"
password = "your_password"
database = "your_database"
```

#### 使用迁移命令

```bash
# 创建迁移文件
go run cmd/main.go migrate create create_users

# 执行迁移
go run cmd/main.go migrate up

# 回滚迁移
go run cmd/main.go migrate down

# 查看迁移状态
go run cmd/main.go migrate status
```

#### 迁移文件格式

在 migrations_dir 目录下创建的迁移文件格式如下：

```sql
-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
```

## 贡献

欢迎提交 Issue 和 Pull Request。

## 许可证

MIT License
