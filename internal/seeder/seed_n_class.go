package seeder

import (
	"encoding/json"
	"os"

	"nso-server/internal/model"
	"nso-server/internal/pkg/database"
	"nso-server/internal/pkg/logger"
)

type SkillOptionTemplateLite struct {
	ID int `json:"id"`
}

// DTOs tương ứng với JSON
type NClassSeed struct {
	ClassId       int                 `json:"class_id"`
	Name          string              `json:"name"`
	SkillTemplate []SkillTemplateSeed `json:"skill_template"`
}

type SkillTemplateSeed struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	MaxPoint    int          `json:"max_point"`
	Type        int          `json:"type"`
	IconId      int          `json:"icon_id"`
	Description string       `json:"description"`
	Skill       []SkillSeed  `json:"skill"`
}

type SkillSeed struct {
	SkillId  int               `json:"skill_id"`
	Point    int               `json:"point"`
	Level    int               `json:"level"`
	ManaUse  int               `json:"mana_use"`
	CoolDown int               `json:"cool_down"`
	Dx       int               `json:"dx"`
	Dy       int               `json:"dy"`
	MaxFight int               `json:"max_fight"`
	Options  []SkillOptionSeed `json:"options"`
}

type SkillOptionSeed struct {
	Param          int                   `json:"param"`
	OptionTemplate SkillOptionTemplateId `json:"option_template"`
}

type SkillOptionTemplateId struct {
	Id int `json:"id"`
}

func SeedNClass(db *database.Database) {
	// ⚠️ Bỏ qua nếu đã có
	var count int64
	if err := db.Model(&model.NClass{}).Count(&count).Error; err != nil {
		logger.Fatalf("[SEED] ❌ Failed to count NClass: %v", err)
	}
	if count > 0 {
		logger.Infof("[SEED] ✅ NClasses already seeded — skipping.")
		return
	}

	// ✅ Load cache SkillOptionTemplate từ file JSON
	optionTemplateMap := loadSkillOptionTemplateMap("data/skill_option_template.json")

	// ✅ Load NClass JSON
	nclassRaw, err := os.ReadFile("data/n_class.json")
	if err != nil {
		logger.Fatalf("[SEED] ❌ Failed to read NClass.json: %v", err)
	}

	var nclassJSON struct {
		NClass []NClassSeed `json:"n_class"`
	}
	if err := json.Unmarshal(nclassRaw, &nclassJSON); err != nil {
		logger.Fatalf("[SEED] ❌ Failed to parse NClass.json: %v", err)
	}

	// ✅ Insert từng NClass
	for _, nc := range nclassJSON.NClass {
		logger.Infof("[SEED] %s", nc.Name)

		nclass := model.NClass{
			ID:   nc.ClassId,
			Name: nc.Name,
		}

		for _, st := range nc.SkillTemplate {
			skillTemplate := model.SkillTemplate{
				ID:          st.Id,
				Name:        st.Name,
				MaxPoint:    st.MaxPoint,
				Type:        st.Type,
				IconID:      st.IconId,
				Description: st.Description,
			}

			for _, sk := range st.Skill {
				skill := model.Skill{
					ID:         sk.SkillId,
					Point:      sk.Point,
					Level:      sk.Level,
					ManaUse:    sk.ManaUse,
					CoolDown:   sk.CoolDown,
					Dx:         sk.Dx,
					Dy:         sk.Dy,
					MaxFight:   sk.MaxFight,
					SkillOptions: []model.SkillOption{},
				}

				for _, opt := range sk.Options {
					if _, ok := optionTemplateMap[opt.OptionTemplate.Id]; !ok {
						continue
					}
					skill.SkillOptions = append(skill.SkillOptions, model.SkillOption{
						SkillID:               sk.SkillId,
						SkillOptionTemplateID: opt.OptionTemplate.Id,
						Param:                 opt.Param,
					})
				}

				skillTemplate.Skills = append(skillTemplate.Skills, skill)
			}

			nclass.SkillTemplates = append(nclass.SkillTemplates, skillTemplate)
		}

		if err := db.Create(&nclass).Error; err != nil {
			logger.Errorf("[SEED] ❌ Failed to insert NClass %s: %v", nc.Name, err)
		}
	}

	logger.Info("[SEED] ✅ NClass seed complete.")
}

func loadSkillOptionTemplateMap(path string) map[int]struct{} {
	file, err := os.ReadFile(path)
	if err != nil {
		logger.Fatalf("[SEED] ❌ Failed to read SkillOptionTemplate JSON: %v", err)
	}

	var raw struct {
		SkillOptionTemplate []SkillOptionTemplateLite `json:"skill_option_template"`
	}
	if err := json.Unmarshal(file, &raw); err != nil {
		logger.Fatalf("[SEED] ❌ Failed to parse SkillOptionTemplate JSON: %v", err)
	}

	result := make(map[int]struct{})
	for _, tpl := range raw.SkillOptionTemplate {
		result[tpl.ID] = struct{}{}
	}

	logger.Infof("[SEED] ✅ Loaded %d SkillOptionTemplate IDs from file", len(result))
	return result
}
