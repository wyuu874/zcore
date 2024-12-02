package config

// App 应用配置
type App struct {
	Debug          bool   `mapstructure:"debug" json:"debug"`                       // 是否开启调试模式
	Host           string `mapstructure:"host" json:"host"`                         // 服务地址
	Port           string `mapstructure:"port" json:"port"`                         // 端口
	ReadTimeout    int    `mapstructure:"read_timeout" json:"read_timeout"`         // 读取超时时间, 单位: 秒
	WriteTimeout   int    `mapstructure:"write_timeout" json:"write_timeout"`       // 写入超时时间, 单位: 秒
	IdleTimeout    int    `mapstructure:"idle_timeout" json:"idle_timeout"`         // 空闲超时时间, 单位: 秒
	MaxHeaderBytes int    `mapstructure:"max_header_bytes" json:"max_header_bytes"` // 最大请求头字节数, 单位: MB
}
