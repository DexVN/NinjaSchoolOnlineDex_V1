package not_login

import (
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"nso-server/internal/infra"
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

	log.Println("Login Request:")
	log.Printf("- Username: %s", username)
	log.Printf("- Password: %s", password)
	log.Printf("- Version: %s", version)
	log.Printf("- DeviceID: %s", deviceID)
	log.Printf("- Other: %s", otherInfo)
	log.Printf("- RandomToken: %s", randomToken)
	log.Printf("- ServerLogin: %d", serverLogin)

	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	if username == "" || password == "" {
		log.Println("\u274C Login: missing username or password")
		sendLoginFail(s, "Tên đăng nhập hoặc mật khẩu không được để trống")
		return
	}

	// Tìm account
	var acc model.Account
	err := infra.DB.Where("username = ?", username).First(&acc).Error
	if err != nil {
		log.Printf("\u274C Login failed: account '%s' not found", username)
		sendLoginFail(s, "Tài khoản không tồn tại")
		return
	}

	// So sánh password đã mã hóa
	if bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password)) != nil {
		log.Printf("\u274C Login failed: wrong password for account '%s'", username)
		sendLoginFail(s, "Mật khẩu không đúng")
		return
	}

	log.Printf("\u2705 Login success: account ID=%d", acc.ID)

	// Cập nhật token nếu chưa có
	if acc.RandomToken == "" && randomToken != "" {
		acc.RandomToken = randomToken
		infra.DB.Save(&acc)
	}

	// Gắn AccountID vào session nếu có session record
	if s.ClientSessionID != nil {
		infra.DB.Model(&model.ClientSession{}).
			Where("id = ?", *s.ClientSessionID).
			Update("account_id", acc.ID)
		log.Printf(" Linked session %d → account %d", *s.ClientSessionID, acc.ID)
	}

	// 🔒 Kick session cũ nếu có và gán session mới
	s.OnLoginSuccess(int(acc.ID))

	// ✅ TODO: Gửi danh sách nhân vật nếu muốn
	// sendCharacterList(s, acc.ID)
}

func sendLoginFail(s *net.Session, reason string) {
	log.Println("❌ Login failed:", reason)
	w := proto.NewWriter()
	w.WriteUTF(reason)
	s.SendMessageWithCommand(proto.CmdServerDialog, w)
}
