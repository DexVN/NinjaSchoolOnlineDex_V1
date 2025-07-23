package app

import (
	"encoding/json"
	"os"

	"nso-server/internal/infra"
	"nso-server/internal/model"
)

func SeedSkillOptionTemplate() {
	path := "data/skill_option_template.json"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		infra.Log.Warnf("[SEED] ⚠ File not found: %s", path)
		return
	}

	data, err := os.ReadFile(path)
	if err != nil {
		infra.Log.Fatalf("[SEED] ❌ Failed to read %s: %v", path, err)
	}

	var jsonData struct {
		SkillOptionTemplate []model.SkillOptionTemplate `json:"skill_option_template"`
	}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		infra.Log.Fatalf("[SEED] ❌ Failed to parse JSON: %v", err)
	}

	db := infra.DB

	var existingIDs []int
	if err := db.Model(&model.SkillOptionTemplate{}).Select("id").Find(&existingIDs).Error; err != nil {
		infra.Log.Fatalf("[SEED] ❌ Failed to query existing IDs: %v", err)
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
			infra.Log.Fatalf("[SEED] ❌ Failed to insert: %v", err)
		}
		infra.Log.Infof("[SEED] ✅ Seeded %d SkillOptionTemplates", len(newOptions))
	} else {
		infra.Log.Info("[SEED] ✅ SkillOptionTemplates already complete — skipping.")
	}
}
