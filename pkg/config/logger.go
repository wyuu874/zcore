package config

// Logger 日志配置
type Logger struct {
	Channel    string `mapstructure:"channel" json:"channel"`         // 日志输出渠道: console | single | daily
	Path       string `mapstructure:"path" json:"path"`               // 日志文件路径
	Level      string `mapstructure:"level" json:"level"`             // 日志级别: debug | info | warn | error
	MaxSize    int    `mapstructure:"max_size" json:"max_size"`       // 最大日志文件大小
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups"` // 最大日志文件数量
	MaxAge     int    `mapstructure:"max_age" json:"max_age"`         // 最大日志保留天数
}
