package handler

import (
	"log"

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

	log.Println("🔐 Login Request:")
	log.Printf("- Username: %s", username)
	log.Printf("- Password: %s", password)
	log.Printf("- Version: %s", version)
	log.Printf("- DeviceID: %s", deviceID)
	log.Printf("- Other: %s", otherInfo)
	log.Printf("- RandomToken: %s", randomToken)
	log.Printf("- ServerLogin: %d", serverLogin)

	// Tìm account
	var acc model.Account
	err := infra.DB.Where("username = ?", username).First(&acc).Error
	if err != nil {
		// ❗ Tài khoản chưa tồn tại → tạo mới
		acc = model.Account{
			Username:     username,
			Password:     password,
			RandomToken:  randomToken,
		}
		if err := infra.DB.Create(&acc).Error; err != nil {
			log.Printf("❌ Cannot create account: %v", err)
			return
		}
		log.Printf("✅ Created new account ID=%d", acc.ID)
	} else {
		log.Printf("✅ Found existing account ID=%d", acc.ID)

		// Cập nhật token nếu chưa có
		if acc.RandomToken == "" && randomToken != "" {
			acc.RandomToken = randomToken
			infra.DB.Save(&acc)
		}
	}

	// Gắn AccountID vào session nếu có session record
	if s.ClientSessionID != nil {
		infra.DB.Model(&model.ClientSession{}).
			Where("id = ?", *s.ClientSessionID).
			Update("account_id", acc.ID)
		log.Printf("🔗 Linked session %d → account %d", *s.ClientSessionID, acc.ID)
	}

}
