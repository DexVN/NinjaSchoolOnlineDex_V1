package not_login

import (
	"log"
	"strings"

	"nso-server/internal/infra"
	"nso-server/internal/model"
	"nso-server/internal/net"
	"nso-server/internal/proto"
)

func HandleClientInfo(msg *proto.Message, s *net.Session) {
	r := msg.Reader()

	// Đọc lần lượt từng field và kiểm tra lỗi nếu cần
	clientType, _ := r.ReadByte()
	zoomLevel, _ := r.ReadByte()
	isGprs, _ := r.ReadBool()
	width, _ := r.ReadInt32()
	height, _ := r.ReadInt32()
	isQwerty, _ := r.ReadBool()
	isTouch, _ := r.ReadBool()
	platform, _ := r.ReadUTF()
	versionIP, _ := r.ReadInt32()
	_, _ = r.ReadByte() // byte thừa
	lang, _ := r.ReadByte()
	userProvider, _ := r.ReadInt32()
	clientAgent, _ := r.ReadUTF()

	log.Printf("📱 Client info:")
	log.Printf("- Type: %d, Zoom: %d, GPRS: %v", clientType, zoomLevel, isGprs)
	log.Printf("- Screen: %dx%d, QWERTY: %v, Touch: %v", width, height, isQwerty, isTouch)
	log.Printf("- Platform: %s", platform)
	log.Printf("- VersionIP: %d, Lang: %d, Provider: %d", versionIP, lang, userProvider)
	log.Printf("- Agent: %s", clientAgent)

	// Parse IP từ RemoteAddr (thường dạng "ip:port")
	remoteIP := strings.Split(s.Conn().RemoteAddr().String(), ":")[0]

	// Tạo bản ghi session client
	session := model.ClientSession{
		ClientType:   int16(clientType),
		ZoomLevel:    int16(zoomLevel),
		IsGprs:       isGprs,
		Width:        int(width),
		Height:       int(height),
		IsQwerty:     isQwerty,
		IsTouch:      isTouch,
		Platform:     platform,
		VersionIP:    int(versionIP),
		Language:     int16(lang),
		UserProvider: int(userProvider),
		ClientAgent:  clientAgent,
		RemoteAddr:   remoteIP,
	}

	// Lưu vào DB
	if err := infra.DB.Create(&session).Error; err != nil {
		log.Printf("❌ Failed to insert client session: %v", err)
	} else {
		s.ClientSessionID = &session.ID
		log.Printf("✅ Client session saved (ID: %d, IP: %s)", session.ID, session.RemoteAddr)
	}
}
