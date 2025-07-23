package not_login

import (
	"log"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"nso-server/internal/infra"
	"nso-server/internal/model"
	"nso-server/internal/net"
	"nso-server/internal/proto"
	"nso-server/internal/utils"
)

func HandleRegister(msg *proto.Message, s *net.Session) {
	r := msg.Reader()

	username, err := r.ReadUTF()
	if err != nil {
		log.Println("❌ Failed to read username:", err)
		return
	}
	password, _ := r.ReadUTF()
	email, _ := r.ReadUTF()

	username = strings.TrimSpace(username)
	email = strings.TrimSpace(email)

	// Validate đơn giản
	if username == "" || password == "" || email == "" {
		log.Println("❌ Register: missing fields")
		w := proto.NewMessage(21)
		_ = w.Writer().WriteByte(0)
		w.WriteUTF("Vui lòng nhập đầy đủ thông tin")
		s.SendMessage(w)
		return
	}
	if len(username) < 5 {
		log.Println("❌ Register: username too short")
		w := proto.NewMessage(21)
		_ = w.Writer().WriteByte(0)
		w.WriteUTF("Tên tài khoản quá ngắn")
		s.SendMessage(w)
		return
	}

	// Kiểm tra tài khoản tồn tại
	var existing model.Account
	if err := infra.DB.Where("username = ?", username).First(&existing).Error; err == nil {
		log.Printf("❌ Register: username '%s' already exists\n", username)
		w := proto.NewMessage(21)
		_ = w.Writer().WriteByte(0)
		w.WriteUTF("Tài khoản đã tồn tại")
		s.SendMessage(w)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("❌ Register: failed to hash password:", err)
		w := proto.NewMessage(21)
		_ = w.Writer().WriteByte(0)
		w.WriteUTF("Lỗi nội bộ khi tạo tài khoản")
		s.SendMessage(w)
		return
	}

	account := model.Account{
		Username:     username,
		Password:     string(hashedPassword),
		Email:        email,
		RandomToken:  utils.GenRandomToken(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := infra.DB.Create(&account).Error; err != nil {
		log.Println("❌ Register: failed to create account:", err)
		w := proto.NewMessage(21)
		_ = w.Writer().WriteByte(0)
		w.WriteUTF("Lỗi tạo tài khoản")
		s.SendMessage(w)
		return
	}

	log.Printf("✅ Registered account: %s (ID: %d)", username, account.ID)

	w := proto.NewMessage(21)
	_ = w.Writer().WriteByte(1)
	w.WriteUTF("Đăng ký thành công")
	s.SendMessage(w)
}
