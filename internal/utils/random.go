package utils

import (
	"strings"
	"time"
)

func GenRandomToken() string {
	return strings.ReplaceAll(time.Now().Format("20060102150405.000"), ".", "")
}
