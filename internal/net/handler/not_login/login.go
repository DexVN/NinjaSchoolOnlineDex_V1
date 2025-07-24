package not_login

import (
	"strings"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"nso-server/internal/model"
	"nso-server/internal/net"
	"nso-server/internal/pkg/di"
	"nso-server/internal/proto"
)

type LoginHandler struct {
	Deps *di.Dependencies
}

func NewLoginHandler(deps *di.Dependencies) *LoginHandler {
	return &LoginHandler{Deps: deps}
}

func (h *LoginHandler) Handle(msg *proto.Message, s *net.Session) {
	r := msg.Reader()

	username, _ := r.ReadUTF()
	password, _ := r.ReadUTF()
	version, _ := r.ReadUTF()
	deviceID, _ := r.ReadUTF()
	otherInfo, _ := r.ReadUTF()
	randomToken, _ := r.ReadUTF()
	serverLogin, _ := r.ReadByte()

	log := h.Deps.Log

	log.Info("🔐 Login Request:")
	log.Infof("- Username: %s", username)
	log.Infof("- Password: %s", password)
	log.Infof("- Version: %s", version)
	log.Infof("- DeviceID: %s", deviceID)
	log.Infof("- Other: %s", otherInfo)
	log.Infof("- RandomToken: %s", randomToken)
	log.Infof("- ServerLogin: %d", serverLogin)

	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	if username == "" || password == "" {
		log.Warn("❌ Login: missing username or password")
		sendLoginFail(s, log, "Thông tin tài khoản hoặc mật khẩu không chính xác")
		return
	}

	acc, err := h.findAccountByUsername(username)
	if err != nil {
		log.Warnf("❌ Login failed: account not found (username=%s)", username)
		sendLoginFail(s, log, "Thông tin tài khoản hoặc mật khẩu không chính xác")
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password)) != nil {
		log.Warnf("❌ Login failed: wrong password (username=%s)", username)
		sendLoginFail(s, log, "Thông tin tài khoản hoặc mật khẩu không chính xác")
		return
	}

	log.Infof("✅ Login success: account ID=%d", acc.ID)

	if acc.RandomToken == "" && randomToken != "" {
		acc.RandomToken = randomToken
		h.Deps.DB.Save(&acc)
	}

	if s.ClientSessionID != nil {
		h.Deps.DB.Model(&model.ClientSession{}).
			Where("id = ?", *s.ClientSessionID).
			Update("account_id", acc.ID)
		log.Infof("🔗 Linked session %d → account %d", *s.ClientSessionID, acc.ID)
	}

	s.OnLoginSuccess(int(acc.ID))
}

func (h *LoginHandler) findAccountByUsername(username string) (*model.Account, error) {
	var acc model.Account
	err := h.Deps.DB.Where("username = ?", username).First(&acc).Error
	if err != nil {
		return nil, err
	}
	return &acc, nil
}

func sendLoginFail(s *net.Session, log *zap.SugaredLogger, reason string) {
	log.Warn("❌ Login failed: " + reason)
	w := proto.NewWriter()
	w.WriteUTF(reason)
	s.SendMessageWithCommand(proto.CmdServerDialog, w)
}
