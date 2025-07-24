package seeder

import (
	"encoding/json"
	"os"

	"nso-server/internal/model"
	"nso-server/internal/pkg/database"
	"nso-server/internal/pkg/logger"

	"gorm.io/gorm"
)

type serverSeedFile struct {
	Server []model.Server `json:"server"`
}

func SeedServer(db *database.Database) {
	data, err := os.ReadFile("data/server.json")
	if err != nil {
		logger.Fatalf("❌ Failed to read server.json: %v", err)
	}

	var seed serverSeedFile
	if err := json.Unmarshal(data, &seed); err != nil {
		logger.Fatalf("❌ Failed to parse server.json: %v", err)
	}

	for _, s := range seed.Server {
		var existing model.Server
		if err := db.Where("code = ?", s.Code).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&s).Error; err != nil {
				logger.WithError(err).Warnf("⚠️ Failed to seed server %s", s.Code)
			} else {
				logger.Infof("✅ Seeded server: %s", s.Code)
			}
		} else {
			logger.Info("[SEED] ✅ Server already complete — skipping.")
		}
	}
}
