package config

// Database 数据库配置
type Database struct {
	Debug         bool   `mapstructure:"debug" json:"debug"`                   // 调试模式
	Host          string `mapstructure:"host" json:"host"`                     // 服务器地址
	Port          string `mapstructure:"port" json:"port"`                     // 端口
	Username      string `mapstructure:"username" json:"username"`             // 用户名
	Password      string `mapstructure:"password" json:"password"`             // 密码
	Database      string `mapstructure:"database" json:"database"`             // 数据库名
	Charset       string `mapstructure:"charset" json:"charset"`               // 字符集
	MaxIdleConns  int    `mapstructure:"max_idle_conns" json:"max_idle_conns"` // 最大空闲连接数
	MaxOpenConns  int    `mapstructure:"max_open_conns" json:"max_open_conns"` // 最大打开连接数
	MaxLifetime   int    `mapstructure:"max_lifetime" json:"max_lifetime"`     // 连接最大生命周期
	MigrationsDir string `mapstructure:"migrations_dir" json:"migrations_dir"` // 迁移文件目录
}
