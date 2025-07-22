package model

import "time"

type Skill struct {
	ID       int `gorm:"primaryKey;autoIncrement:false;column:id" json:"id"`
	Point    int `gorm:"column:point" json:"point"`
	Level    int `gorm:"column:level" json:"level"`
	ManaUse  int `gorm:"column:mana_use" json:"mana_use"`
	CoolDown int `gorm:"column:cool_down" json:"cool_down"`
	Dx       int `gorm:"column:dx" json:"dx"`
	Dy       int `gorm:"column:dy" json:"dy"`
	MaxFight int `gorm:"column:max_fight" json:"max_fight"`

	SkillTemplateID int           `gorm:"column:skill_template_id" json:"skill_template_id"`
	SkillTemplate   SkillTemplate

	SkillOptions []SkillOption `gorm:"foreignKey:SkillID"`

	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

