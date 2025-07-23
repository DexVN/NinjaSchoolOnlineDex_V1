package not_login

import (
	"strings"

	"golang.org/x/crypto/bcrypt"

	logger "nso-server/internal/infra"
	"nso-server/internal/model"
	"nso-server/internal/net"
	"nso-server/internal/proto"
)

func HandleLogin(msg *proto.Message, s *net.Session) {
	r := msg.Reader()

	username, _ := r.ReadUTF()
	password, _ := r.ReadUTF()
	version, _ := r.ReadUTF()
	deviceID, _ := r.ReadUTF()
	otherInfo, _ := r.ReadUTF()
	randomToken, _ := r.ReadUTF()
	serverLogin, _ := r.ReadByte()

	logger.Log.Info("🔐 Login Request:")
	logger.Log.Infof("- Username: %s", username)
	logger.Log.Infof("- Password: %s", password)
	logger.Log.Infof("- Version: %s", version)
	logger.Log.Infof("- DeviceID: %s", deviceID)
	logger.Log.Infof("- Other: %s", otherInfo)
	logger.Log.Infof("- RandomToken: %s", randomToken)
	logger.Log.Infof("- ServerLogin: %d", serverLogin)

	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	if username == "" || password == "" {
		logger.Log.Warn("❌ Login: missing username or password")
		sendLoginFail(s, "Thông tin tài khoản hoặc mật khẩu không chính xác")
		return
	}

	// Tìm account
	var acc model.Account
	err := logger.DB.Where("username = ?", username).First(&acc).Error
	if err != nil {
		logger.Log.WithField("username", username).
			Warn("❌ Login failed: account not found")
		sendLoginFail(s, "Thông tin tài khoản hoặc mật khẩu không chính xác")
		return
	}

	// So sánh password đã mã hóa
	if bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password)) != nil {
		logger.Log.WithField("username", username).
			Warn("❌ Login failed: wrong password")
		sendLoginFail(s, "Thông tin tài khoản hoặc mật khẩu không chính xác")
		return
	}

	logger.Log.Infof("✅ Login success: account ID=%d", acc.ID)

	// Cập nhật token nếu chưa có
	if acc.RandomToken == "" && randomToken != "" {
		acc.RandomToken = randomToken
		logger.DB.Save(&acc)
	}

	// Gắn AccountID vào session nếu có session record
	if s.ClientSessionID != nil {
		logger.DB.Model(&model.ClientSession{}).
			Where("id = ?", *s.ClientSessionID).
			Update("account_id", acc.ID)
		logger.Log.Infof("🔗 Linked session %d → account %d", *s.ClientSessionID, acc.ID)
	}

	// 🔒 Kick session cũ nếu có và gán session mới
	s.OnLoginSuccess(int(acc.ID))

	// ✅ TODO: Gửi danh sách nhân vật nếu muốn
	// sendCharacterList(s, acc.ID)
}

func sendLoginFail(s *net.Session, reason string) {
	logger.Log.Warnf("❌ Login failed: %s", reason)
	w := proto.NewWriter()
	w.WriteUTF(reason)
	s.SendMessageWithCommand(proto.CmdServerDialog, w)
}
