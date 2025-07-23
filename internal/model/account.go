package model

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	ID           int         `gorm:"primaryKey;column:id"                      json:"id"`
	Username     string      `gorm:"column:username;uniqueIndex;not null"     json:"username"`
	Password     string      `gorm:"column:password;not null"                 json:"-"`            // Không expose ra JSON
	Email        string      `gorm:"column:email;not null"                    json:"email"`
	RandomToken  string      `gorm:"column:random_token"                      json:"random_token"`
	Characters   []Character `gorm:"foreignKey:AccountID"                     json:"characters,omitempty"`

	CreatedAt    time.Time   `gorm:"column:created_at;not null;autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time   `gorm:"column:updated_at;not null;autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"                     json:"deleted_at,omitempty"`
}
