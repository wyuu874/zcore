package cache

import (
	"context"
	"github.com/wyuu874/zcore/pkg/rdb"
	"time"
)

// Put 存储缓存
func Put(key string, val interface{}, timeout time.Duration) bool {
	err := rdb.RDB().Set(context.Background(), key, val, timeout).Err()
	return err == nil
}

// Get 获取缓存
func Get(key string) string {
	val, err := rdb.RDB().Get(context.Background(), key).Result()
	if err != nil || val == "" {
		return ""
	}
	return val
}

// IsExist 判断缓存是否存在
func IsExist(key string) bool {
	b, err := rdb.RDB().Exists(context.Background(), key).Result()
	if err != nil {
		return false
	}
	return b > 0
}

// Del 删除缓存
func Del(key string) error {
	return rdb.RDB().Del(context.Background(), key).Err()
}
