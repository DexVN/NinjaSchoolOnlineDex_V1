package lang

import "fmt"

type i18nImpl struct{}

func (i *i18nImpl) Get(key string) string {
	return Get(key)
}

func (i *i18nImpl) Getf(key string, args ...any) string {
	return fmt.Sprintf(Get(key), args...)
}

func NewI18n() I18n {
	return &i18nImpl{}
}
