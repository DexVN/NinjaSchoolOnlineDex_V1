package model

import "time"

type DataVersion struct {
	Key     string `gorm:"primaryKey;column:key;size:50" json:"key"`
	Version int    `gorm:"column:version;default:1" json:"version"`

	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

