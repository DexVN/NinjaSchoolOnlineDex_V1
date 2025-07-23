package model

import (
	"time"

	"gorm.io/gorm"
)

type Skill struct {
	ID       int `gorm:"primaryKey;autoIncrement:false;column:id;not null" json:"id"`
	Point    int `gorm:"column:point;not null;default:0"                   json:"point"`
	Level    int `gorm:"column:level;not null;default:1"                   json:"level"`
	ManaUse  int `gorm:"column:mana_use;not null;default:0"               json:"mana_use"`
	CoolDown int `gorm:"column:cool_down;not null;default:0"              json:"cool_down"`
	Dx       int `gorm:"column:dx;not null;default:0"                     json:"dx"`
	Dy       int `gorm:"column:dy;not null;default:0"                     json:"dy"`
	MaxFight int `gorm:"column:max_fight;not null;default:1"              json:"max_fight"`

	SkillTemplateID int           `gorm:"column:skill_template_id;not null;index"                json:"skill_template_id"`
	SkillTemplate   SkillTemplate `gorm:"foreignKey:SkillTemplateID;constraint:OnDelete:CASCADE" json:"-"`

	SkillOptions []SkillOption `gorm:"foreignKey:SkillID" json:"options,omitempty"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"  json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"  json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"                 json:"deleted_at,omitempty"`
}
