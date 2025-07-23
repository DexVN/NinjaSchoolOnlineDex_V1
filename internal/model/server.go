package model

import (
	"time"

	"gorm.io/gorm"
)

type Server struct {
	ID       int    `gorm:"primaryKey;column:id;autoIncrement"         json:"id"`
	Name     string `gorm:"column:name;size:100;not null"              json:"name"`
	Code     string `gorm:"column:code;size:20;not null;uniqueIndex"   json:"code"` // e.g., "0", "1", "2"
	IP       string `gorm:"column:ip;size:50;not null"                 json:"ip"`
	Port     int    `gorm:"column:port;not null"                       json:"port"`
	Language int8   `gorm:"column:language;not null;default:0"         json:"language"` // 0: VI, 1: EN

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"  json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"  json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"                 json:"deleted_at,omitempty"`
}
