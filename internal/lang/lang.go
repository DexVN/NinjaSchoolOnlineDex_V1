package lang

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var messages = make(map[string]map[string]string)
var fallbackMessages = make(map[string]map[string]string)

// Luôn ưu tiên tiếng Việt → fallback tiếng Anh nếu thiếu
func Init(lang string) error {
	// Load tiếng Việt (mặc định)
	if err := loadMessages("vi", &messages); err != nil {
		return err
	}

	// Nếu khác "vi", load thêm fallback
	if lang != "vi" {
		_ = loadMessages(lang, &messages) // override nếu có
	}

	// Load tiếng Anh để fallback
	_ = loadMessages("en", &fallbackMessages)

	return nil
}

func loadMessages(lang string, target *map[string]map[string]string) error {
	path := filepath.Join("internal", "lang", "message", lang+".json")
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, target)
}

// Get chuỗi theo key, fallback tiếng Anh nếu thiếu
func Get(key string) string {
	group, subkey, ok := splitKey(key)
	if !ok {
		return "{" + key + "}"
	}

	if val := lookup(messages, group, subkey); val != "" {
		return val
	}
	if val := lookup(fallbackMessages, group, subkey); val != "" {
		return val
	}
	return "{" + key + "}"
}

func Getf(key string, args ...any) string {
	return fmt.Sprintf(Get(key), args...)
}

// Tách key: "login.success" => "login", "success"
func splitKey(key string) (group string, subkey string, ok bool) {
	parts := strings.SplitN(key, ".", 2)
	if len(parts) != 2 {
		return "", "", false
	}
	return parts[0], parts[1], true
}

func lookup(data map[string]map[string]string, group, subkey string) string {
	if m, ok := data[group]; ok {
		if val, ok := m[subkey]; ok {
			return val
		}
	}
	return ""
}

func GetLangDisplayName(code string) string {
	switch code {
	case "vi":
		return "Vietnamese"
	case "en":
		return "English"
	case "jp":
		return "Japanese"
	case "th":
		return "Thai"
	default:
		return code
	}
}
