package model

import "time"

type Character struct {
	ID         int    `gorm:"primaryKey;column:id" json:"id"`
	SlotIndex  int16  `gorm:"column:slot_index" json:"slot_index"`
	Name       string `gorm:"column:name;size:50;not null" json:"name"`
	Gender     int16  `gorm:"column:gender" json:"gender"`
	ClassName  string `gorm:"column:class_name;not null" json:"class_name"`
	Level      int16  `gorm:"column:level" json:"level"`
	PartHead   int16  `gorm:"column:part_head" json:"part_head"`
	PartWeapon int16  `gorm:"column:part_weapon" json:"part_weapon"`
	PartBody   int16  `gorm:"column:part_body" json:"part_body"`
	PartLeg    int16  `gorm:"column:part_leg" json:"part_leg"`

	AccountID int     `gorm:"column:account_id" json:"account_id"`
	Account   Account

	ServerID int    `gorm:"column:server_id" json:"server_id"`
	Server   Server `gorm:"foreignKey:ServerID"`

	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
