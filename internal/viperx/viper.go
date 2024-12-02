package viperx

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

// InitConfig 初始化配置
func InitConfig() {
	// 设置配置文件名称和类型
	viper.SetConfigName("app")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil && !errors.Is(err, viper.ConfigFileNotFoundError{}) {
		panic(fmt.Sprintf("读取配置文件失败: %s", err.Error()))
	}
}
