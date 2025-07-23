package model

import (
	"time"

	"gorm.io/gorm"
)

type Character struct {
	ID         int    `gorm:"primaryKey;column:id"                           json:"id"`
	SlotIndex  int16  `gorm:"column:slot_index;not null"                    json:"slot_index"`
	Name       string `gorm:"column:name;size:50;not null"                 json:"name"`
	Gender     int16  `gorm:"column:gender;not null"                        json:"gender"`
	ClassName  string `gorm:"column:class_name;not null"                   json:"class_name"`
	Level      int16  `gorm:"column:level;default:1;not null"              json:"level"`
	PartHead   int16  `gorm:"column:part_head;default:0"                   json:"part_head"`
	PartWeapon int16  `gorm:"column:part_weapon;default:0"                 json:"part_weapon"`
	PartBody   int16  `gorm:"column:part_body;default:0"                   json:"part_body"`
	PartLeg    int16  `gorm:"column:part_leg;default:0"                    json:"part_leg"`

	AccountID int     `gorm:"column:account_id;not null;index"             json:"account_id"`
	Account   Account `gorm:"foreignKey:AccountID"                         json:"-"`

	ServerID int    `gorm:"column:server_id;not null;index"               json:"server_id"`
	Server   Server `gorm:"foreignKey:ServerID"                           json:"-"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"     json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"     json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"                    json:"deleted_at,omitempty"`
}
