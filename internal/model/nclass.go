package model

import "time"

type NClass struct {
	ID   int    `gorm:"primaryKey;autoIncrement:false;column:id" json:"id"`
	Name string `gorm:"column:name;size:50;not null" json:"name"`

	SkillTemplates []SkillTemplate `gorm:"foreignKey:NClassID"`

	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
