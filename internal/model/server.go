package model

import "time"

type Server struct {
	ID        int        `gorm:"primaryKey;column:id" json:"id"`
	Name      string     `gorm:"column:name;size:100;not null" json:"name"`
	Code      string     `gorm:"column:code;size:20;unique;not null" json:"code"` // e.g., "0", "1", "2"
	IP        string     `gorm:"column:ip" json:"ip"`
	Port      int        `gorm:"column:port" json:"port"`
	Language  int8       `gorm:"column:language" json:"language"` // 0: VI, 1: EN
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
