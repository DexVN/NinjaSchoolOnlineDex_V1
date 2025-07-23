package model

import (
	"time"

	"gorm.io/gorm"
)

type SkillOptionTemplate struct {
	ID   int    `gorm:"primaryKey;autoIncrement:false;column:id;not null" json:"id"`
	Name string `gorm:"column:name;size:100;not null"                     json:"name"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"     json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"     json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"                    json:"deleted_at,omitempty"`
}
