package model

import "time"

type SkillTemplate struct {
	ID          int    `gorm:"primaryKey;autoIncrement:false;column:id" json:"id"`
	Name        string `gorm:"column:name;size:100;not null" json:"name"`
	MaxPoint    int    `gorm:"column:max_point" json:"max_point"`
	Type        int    `gorm:"column:type" json:"type"`
	IconID      int    `gorm:"column:icon_id" json:"icon_id"`
	Description string `gorm:"column:description;type:text" json:"description"`

	NClassID int    `gorm:"column:n_class_id" json:"n_class_id"`
	NClass   NClass

	Skills []Skill `gorm:"foreignKey:SkillTemplateID"`

	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}


