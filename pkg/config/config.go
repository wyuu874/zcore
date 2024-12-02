package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"time"
)

// Get 获取配置
func Get(key string) interface{} {
	return viper.Get(key)
}

// GetBool 获取布尔值配置
func GetBool(key string) bool {
	return viper.GetBool(key)
}

// GetFloat64 获取浮点数配置
func GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

// GetInt 获取整数配置
func GetInt(key string) int {
	return viper.GetInt(key)
}

// GetIntSlice 获取整数切片配置
func GetIntSlice(key string) []int {
	return viper.GetIntSlice(key)
}

// GetString 获取字符串配置
func GetString(key string) string {
	return viper.GetString(key)
}

// GetStringMap 获取字符串映射配置
func GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}

// GetStringMapString 获取字符串映射字符串配置
func GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}

// GetStringSlice 获取字符串切片配置
func GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

// GetTime 获取时间配置
func GetTime(key string) time.Time {
	return viper.GetTime(key)
}

// GetDuration 获取持续时间配置
func GetDuration(key string) time.Duration {
	return viper.GetDuration(key)
}

// IsSet 检查配置是否存在
func IsSet(key string) bool {
	return viper.IsSet(key)
}

// AllSettings 获取所有配置
func AllSettings() map[string]interface{} {
	return viper.AllSettings()
}

// GetConfig 通用的配置获取函数
func GetConfig[T any](key string, config *T) {
	loadConfig := GetStringMap(key)
	mapstructure.Decode(loadConfig, config)
}
