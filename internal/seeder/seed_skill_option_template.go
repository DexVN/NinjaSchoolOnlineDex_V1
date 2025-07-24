package seeder

import (
	"encoding/json"
	"os"

	"nso-server/internal/model"
	"nso-server/internal/pkg/database"
	"nso-server/internal/pkg/logger"
)

func SeedSkillOptionTemplate(db *database.Database) {
	path := "data/skill_option_template.json"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		logger.Warnf("[SEED] ⚠ File not found: %s", path)
		return
	}

	data, err := os.ReadFile(path)
	if err != nil {
		logger.Fatalf("[SEED] ❌ Failed to read %s: %v", path, err)
	}

	var jsonData struct {
		SkillOptionTemplate []model.SkillOptionTemplate `json:"skill_option_template"`
	}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		logger.Fatalf("[SEED] ❌ Failed to parse JSON: %v", err)
	}

	var existingIDs []int
	if err := db.Model(&model.SkillOptionTemplate{}).Select("id").Find(&existingIDs).Error; err != nil {
		logger.Fatalf("[SEED] ❌ Failed to query existing IDs: %v", err)
	}

	existingSet := make(map[int]bool)
	for _, id := range existingIDs {
		existingSet[id] = true
	}

	var newOptions []model.SkillOptionTemplate
	for _, opt := range jsonData.SkillOptionTemplate {
		if !existingSet[opt.ID] {
			newOptions = append(newOptions, opt)
		}
	}

	if len(newOptions) > 0 {
		if err := db.Create(&newOptions).Error; err != nil {
			logger.Fatalf("[SEED] ❌ Failed to insert: %v", err)
		}
		logger.Infof("[SEED] ✅ Seeded %d SkillOptionTemplates", len(newOptions))
	} else {
		logger.Info("[SEED] ✅ SkillOptionTemplates already complete — skipping.")
	}
}
