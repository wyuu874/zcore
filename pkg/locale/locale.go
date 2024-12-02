package locale

import (
	"fmt"
	"github.com/wyuu874/zcore/internal/i18nx"
)

// Translate 翻译指定的key
func Translate(lang string, messageID string, templateData map[string]interface{}) string {
	loc := i18nx.InitLocalizer(lang)
	msg, err := loc.Localize(&i18nx.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: templateData,
	})
	if err != nil {
		fmt.Println(err.Error())
		return messageID
	}
	return msg
}
