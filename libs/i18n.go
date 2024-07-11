package libs

import (
	"bcw/config"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

var bundle *i18n.Bundle

func init() {
	bundle = i18n.NewBundle(language.Chinese)

	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	bundle.MustLoadMessageFile("resources/lang/zh.yaml")
	bundle.MustLoadMessageFile("resources/lang/en.yaml")
}

func GetLanguageMsg(id string) string {
	currentLang := config.GetLang()

	if currentLang == "" {
		currentLang = "zh"
	}
	msg := i18n.NewLocalizer(bundle, currentLang).MustLocalize(&i18n.LocalizeConfig{
		MessageID: id,
	})
	return msg
}
