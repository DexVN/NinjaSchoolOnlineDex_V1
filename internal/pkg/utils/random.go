package utils

import (
	"strings"
	"time"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)


func GenRandomToken() string {
	return strings.ReplaceAll(time.Now().Format("20060102150405.000"), ".", "")
}

func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}