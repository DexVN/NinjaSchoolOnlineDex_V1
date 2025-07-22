package model

import "time"

type SkillOptionTemplate struct {
	ID   int    `gorm:"primaryKey;autoIncrement:false;column:id" json:"id"`
	Name string `gorm:"column:name;size:100;not null" json:"name"`

	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

