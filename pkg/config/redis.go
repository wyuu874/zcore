package config

// Redis 配置
type Redis struct {
	Host string `mapstructure:"host" json:"host"` // 服务地址
	Port int    `mapstructure:"port" json:"port"` // 服务端口
	Pass string `mapstructure:"pass" json:"pass"` // 密码
	DB   int    `mapstructure:"db" json:"db"`     // 数据库
}
