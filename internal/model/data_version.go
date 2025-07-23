package model

import (
	"time"

	"gorm.io/gorm"
)

type DataVersion struct {
	Key      string         `gorm:"primaryKey;column:key;size:50;not null" json:"key"`
	Version  int            `gorm:"column:version;default:1;not null"     json:"version"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"     json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"     json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"                    json:"deleted_at,omitempty"`
}
