package not_login

import (
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"nso-server/internal/model"
	"nso-server/internal/net"
	"nso-server/internal/pkg/di"
	"nso-server/internal/pkg/utils"
	"nso-server/internal/proto"
)

type RegisterHandler struct {
	Deps *di.Dependencies
}

func NewRegisterHandler(deps *di.Dependencies) *RegisterHandler {
	return &RegisterHandler{Deps: deps}
}

func (h *RegisterHandler) Handle(msg *proto.Message, s *net.Session) {
	r := msg.Reader()
	lang := h.Deps.I18n
	username, err := r.ReadUTF()
	if err != nil {
		h.Deps.Log.Warn("❌ Failed to read username", err)
		sendRegisterFail(s, lang.Get("common.error_occurred"))
		return
	}
	password, _ := r.ReadUTF()
	email, _ := r.ReadUTF()

	username = strings.TrimSpace(username)
	email = strings.TrimSpace(email)

	if username == "" || password == "" || email == "" {
		sendRegisterFail(s, lang.Get("account.register_incomplete"))
		return
	}
	if len(username) < 5 {
		sendRegisterFail(s, lang.Get("account.register_username_too_short"))
		return
	}
	if !utils.IsValidEmail(email) {
		sendRegisterFail(s, lang.Get("account.register_invalid_email"))
		return
	}

	err = h.Deps.DB.Transaction(func(tx *gorm.DB) error {
		var server model.Server
		if err := tx.Where("code = ?", h.Deps.Config.ServerCode).First(&server).Error; err != nil {
			h.Deps.Log.Error("❌ Không tìm thấy server theo SERVER_CODE", err)
			return errors.New(lang.Get("common.error_occurred"))
		}

		var count int64
		tx.Model(&model.Character{}).
			Where("server_id = ?", server.ID).
			Count(&count)
		if count >= int64(server.MaxPlayers) {
			h.Deps.Log.Warnf("❌ Đã đạt giới hạn người chơi (%d/%d)", count, server.MaxPlayers)
			return errors.New(lang.Get("server.reach_max_players"))
		}

		var existing model.Account
		if err := tx.Where("username = ?", username).First(&existing).Error; err == nil {
			return errors.New(lang.Get("account.register_username_exists"))
		}
		if err := tx.Where("email = ?", email).First(&existing).Error; err == nil {
			return errors.New(lang.Get("account.register_email_exists"))
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			h.Deps.Log.Error("❌ Register: lỗi hash password", err)
			return errors.New(lang.Get("common.error_occurred"))
		}

		account := model.Account{
			Username:    username,
			Password:    string(hashedPassword),
			Email:       email,
			RandomToken: utils.GenRandomToken(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		if err := tx.Create(&account).Error; err != nil {
			h.Deps.Log.Error("❌ Register: lỗi tạo account", err)
			return errors.New(lang.Get("common.error_occurred"))
		}

		h.Deps.Log.Infof("✅ Đăng ký thành công: %s (ID: %d)", username, account.ID)
		return nil
	})

	if err != nil {
		sendRegisterFail(s, err.Error())
		return
	}

	w := proto.NewMessage(proto.CmdServerDialog)
	w.WriteUTF(lang.Get("account.register_success"))
	s.SendMessage(w)
}

func sendRegisterFail(s *net.Session, reason string) {
	w := proto.NewMessage(proto.CmdServerDialog)
	w.WriteUTF(reason)
	s.SendMessage(w)
}
