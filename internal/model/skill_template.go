package model

import (
	"time"

	"gorm.io/gorm"
)

type SkillTemplate struct {
	ID          int    `gorm:"primaryKey;autoIncrement:false;column:id;not null" json:"id"`
	Name        string `gorm:"column:name;size:100;not null"                     json:"name"`
	MaxPoint    int    `gorm:"column:max_point;not null;default:1"              json:"max_point"`
	Type        int    `gorm:"column:type;not null;default:0"                   json:"type"`
	IconID      int    `gorm:"column:icon_id;not null;default:0"                json:"icon_id"`
	Description string `gorm:"column:description;type:text;not null"            json:"description"`

	NClassID int    `gorm:"column:n_class_id;not null;index"                     json:"n_class_id"`
	NClass   NClass `gorm:"foreignKey:NClassID;constraint:OnDelete:CASCADE"      json:"-"`

	Skills []Skill `gorm:"foreignKey:SkillTemplateID"                            json:"skills,omitempty"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"           json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"           json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"                          json:"deleted_at,omitempty"`
}
