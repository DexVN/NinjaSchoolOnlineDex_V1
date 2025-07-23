package model

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	ID              int         `gorm:"primaryKey;column:id"                      json:"id"`
	Username        string      `gorm:"column:username;uniqueIndex;not null"     json:"username"`
	Password        string      `gorm:"column:password;not null"                 json:"-"` // Không expose ra JSON
	Email           string      `gorm:"column:email;not null"                    json:"email"`
	RandomToken     string      `gorm:"column:random_token"                      json:"random_token"`
	EmailVerified   bool        `gorm:"column:email_verified;default:false"      json:"email_verified"`
	EmailVerifiedAt time.Time   `gorm:"column:email_verified_at"                 json:"email_verified_at,omitempty"`
	LoginLock       bool        `gorm:"column:login_lock;default:false"          json:"login_lock"`
	LoginLockTime   time.Time   `gorm:"column:login_lock_time"                   json:"login_lock_time,omitempty"`
	Characters      []Character `gorm:"foreignKey:AccountID"                     json:"characters,omitempty"`

	CreatedAt time.Time      `gorm:"column:created_at;not null;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"                     json:"deleted_at,omitempty"`
}
