package model

import "time"

type Account struct {
	ID        int        `gorm:"primaryKey;column:id" json:"id"`
	Username  string     `gorm:"column:username" json:"username"`
	Password  string     `gorm:"column:password" json:"password"`
	Email     string     `gorm:"column:email" json:"email"`
	RandomToken string   `gorm:"column:random_token" json:"random_token"`
	Characters []Character `gorm:"foreignKey:AccountID"`

	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
