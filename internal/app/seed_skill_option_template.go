package app

import (
	"encoding/json"
	"log"
	"os"

	"nso-server/internal/infra"
	"nso-server/internal/model"
)

func SeedSkillOptionTemplate() {
	path := "data/skill_option_template.json"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("[SEED] ⚠ File not found: %s\n", path)
		return
	}

	// Đọc file JSON
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("[SEED] ❌ Failed to read %s: %v", path, err)
	}

	// Parse dữ liệu
	var jsonData struct {
	SkillOptionTemplate []model.SkillOptionTemplate `json:"skill_option_template"`
}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		log.Fatalf("[SEED] ❌ Failed to parse JSON: %v", err)
	}

	db := infra.DB

	// Lấy danh sách ID đã có
	var existingIDs []int
	if err := db.Model(&model.SkillOptionTemplate{}).Select("id").Find(&existingIDs).Error; err != nil {
		log.Fatalf("[SEED] ❌ Failed to query existing IDs: %v", err)
	}
	existingSet := make(map[int]bool)
	for _, id := range existingIDs {
		existingSet[id] = true
	}

	// Lọc những item mới
	var newOptions []model.SkillOptionTemplate
	for _, opt := range jsonData.SkillOptionTemplate {
		if !existingSet[opt.ID] {
			newOptions = append(newOptions, opt)
		}
	}

	// Thêm vào nếu có item mới
	if len(newOptions) > 0 {
		if err := db.Create(&newOptions).Error; err != nil {
			log.Fatalf("[SEED] ❌ Failed to insert: %v", err)
		}
		log.Printf("[SEED] ✅ Seeded %d SkillOptionTemplates\n", len(newOptions))
	} else {
		log.Println("[SEED] ✅ SkillOptionTemplates already complete — skipping.")
	}
}
