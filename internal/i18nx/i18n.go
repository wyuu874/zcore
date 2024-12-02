package i18nx

import (
	"fmt"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	"golang.org/x/text/language"
	"path/filepath"
	"strings"
)

// LocalizeConfig 本地化配置
type LocalizeConfig = i18n.LocalizeConfig

// Bundle 国际化包
var Bundle *i18n.Bundle

// InitI18n 初始化国际化
func InitI18n(defaultLang string, dir string) {
	// 初始化Bundle
	Bundle = i18n.NewBundle(language.Make(defaultLang))
	Bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	// 加载翻译文件
	loadMessageFile(dir)
}

// InitLocalizer 初始化本地化
func InitLocalizer(lang string) *i18n.Localizer {
	return i18n.NewLocalizer(Bundle, lang)
}

// loadMessageFile 加载翻译文件
func loadMessageFile(dir string) {
	// 判断目录是否存在
	if fileutil.IsExist(dir) && fileutil.IsDir(dir) {
		// 读取目录下的所有文件
		fileNames, _ := fileutil.ListFileNames(dir)
		for _, fileName := range fileNames {
			if !strings.HasSuffix(fileName, ".toml") {
				continue
			}

			filePath := filepath.Join(dir, fileName)
			if _, err := Bundle.LoadMessageFile(filePath); err != nil {
				panic(fmt.Sprintf("加载语言文件失败: %s, %s", fileName, err.Error()))
			}
		}
	}
}
