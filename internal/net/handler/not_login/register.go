package not_login

import (
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	logger "nso-server/internal/infra"
	"nso-server/internal/lang"
	"nso-server/internal/model"
	"nso-server/internal/net"
	"nso-server/internal/proto"
	"nso-server/internal/utils"
)

func HandleRegister(msg *proto.Message, s *net.Session) {
	r := msg.Reader()

	username, err := r.ReadUTF()
	if err != nil {
		logger.Log.WithError(err).Warn("❌ Failed to read username")
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

	// Kiểm tra định dạng email
	if !utils.IsValidEmail(email) {
		sendRegisterFail(s, lang.Get("account.register_invalid_email"))
		return
	}

	var existing model.Account
	// Kiểm tra xem username đã tồn tại chưa
	if err := logger.DB.Where("username = ?", username).First(&existing).Error; err == nil {
		logger.Log.Warnf("❌ Register failed: Username %s already exists", username)
		sendRegisterFail(s, lang.Get("account.register_username_exists"))
		return
	}

	// Kiểm tra xem email đã tồn tại chưa
	if err := logger.DB.Where("email = ?", email).First(&existing).Error; err == nil {
		logger.Log.Warnf("❌ Register failed: Email %s already exists", email)
		sendRegisterFail(s, lang.Get("account.register_email_exists"))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Log.WithError(err).Error("❌ Register: lỗi hash password")
		sendRegisterFail(s, lang.Get("common.error_occurred"))
		return
	}

	account := model.Account{
		Username:    username,
		Password:    string(hashedPassword),
		Email:       email,
		RandomToken: utils.GenRandomToken(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := logger.DB.Create(&account).Error; err != nil {
		logger.Log.WithError(err).Error("❌ Register: lỗi tạo account")
		sendRegisterFail(s, lang.Get("common.error_occurred"))
		return
	}

	logger.Log.Infof("✅ Đăng ký thành công: %s (ID: %d)", username, account.ID)

	w := proto.NewMessage(proto.CmdServerDialog)
	w.WriteUTF(lang.Get("account.register_success"))
	s.SendMessage(w)
}

func sendRegisterFail(s *net.Session, reason string) {
	logger.Log.Warnf("❌ Register failed: %s", reason)
	w := proto.NewMessage(proto.CmdServerDialog)
	w.WriteUTF(reason)
	s.SendMessage(w)
}
