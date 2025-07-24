package database

import (
	"time"

	"nso-server/internal/pkg/logger"

	"gorm.io/gorm"
)

// NewGormConfig trả về config GORM tối ưu cho game server
func NewGormConfig() *gorm.Config {
	return &gorm.Config{
		// ⚡ Tăng hiệu năng: không dùng transaction cho mỗi câu lệnh nhỏ
		SkipDefaultTransaction: true,

		// 🔍 Ghi log SQL query, cảnh báo query chậm
		Logger: logger.NewGormLogger(100 * time.Millisecond),

		// 🧠 (Có thể thêm NamingStrategy nếu cần)
	}
}
