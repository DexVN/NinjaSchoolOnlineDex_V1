package model

type Language int8

const (
	LanguageVI Language = iota // 0
	LanguageEN                 // 1
)

func (l Language) String() string {
	switch l {
	case LanguageVI:
		return "Vietnamese"
	case LanguageEN:
		return "English"
	default:
		return "Unknown"
	}
}
