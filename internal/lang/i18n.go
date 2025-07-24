package lang

type I18n interface {
	Get(key string) string
	Getf(key string, args ...any) string
}
