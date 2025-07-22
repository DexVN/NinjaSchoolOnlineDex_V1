package model

import "time"

type SkillOption struct {
	SkillID              int
	SkillOptionTemplateID int

	Skill               Skill               `gorm:"foreignKey:SkillID;constraint:OnDelete:CASCADE"`
	SkillOptionTemplate SkillOptionTemplate `gorm:"foreignKey:SkillOptionTemplateID;constraint:OnDelete:CASCADE"`

	Param int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	// Khóa chính tổng hợp
	// GORM tự hiểu theo 2 field nếu không dùng tag `gorm:"primaryKey"`
}

func (SkillOption) TableName() string {
	return "skill_options"
}
