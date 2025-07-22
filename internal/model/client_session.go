// internal/model/client_session.go
package model

import "time"

type ClientSession struct {
	ID        int  `gorm:"primaryKey;column:id" json:"id"`
	AccountID *int `gorm:"column:account_id"`

	ClientType   int16  `gorm:"column:client_type" json:"client_type"`
	ZoomLevel    int16  `gorm:"column:zoom_level" json:"zoom_level"`
	IsGprs       bool   `gorm:"column:is_gprs" json:"is_gprs"`
	Width        int    `gorm:"column:screen_width" json:"screen_width"`
	Height       int    `gorm:"column:screen_height" json:"screen_height"`
	IsQwerty     bool   `gorm:"column:is_qwerty" json:"is_qwerty"`
	IsTouch      bool   `gorm:"column:is_touch" json:"is_touch"`
	Platform     string `gorm:"column:platform;size:50" json:"platform"`
	VersionIP    int    `gorm:"column:version_ip" json:"version_ip"`
	Language     int16  `gorm:"column:language" json:"language"`
	UserProvider int    `gorm:"column:user_provider" json:"user_provider"`
	ClientAgent  string `gorm:"column:client_agent;size:255" json:"client_agent"`

	RemoteAddr string    `gorm:"column:remote_addr;size:50" json:"remote_addr"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
}
