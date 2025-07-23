package model

import (
	"time"

	"gorm.io/gorm"
)

type SkillOption struct {
	SkillID               int `gorm:"primaryKey;column:skill_id;not null"`
	SkillOptionTemplateID int `gorm:"primaryKey;column:skill_option_template_id;not null"`

	Param int `gorm:"column:param;not null"`

	Skill               Skill               `gorm:"foreignKey:SkillID;constraint:OnDelete:CASCADE"               json:"-"`
	SkillOptionTemplate SkillOptionTemplate `gorm:"foreignKey:SkillOptionTemplateID;constraint:OnDelete:CASCADE" json:"-"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"                json:"deleted_at,omitempty"`
}
