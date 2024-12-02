package config

// Locale 国际化配置
type Locale struct {
	DefaultLang string `mapstructure:"default_lang" json:"default_lang"` // 默认语言
	Dir         string `mapstructure:"dir" json:"dir"`                   // 语言目录
}
