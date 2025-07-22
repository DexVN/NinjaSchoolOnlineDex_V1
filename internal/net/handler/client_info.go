// internal/net/handler/client_info.go
package handler

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

	// Đọc các trường theo đúng thứ tự gửi từ client
	clientType, _ := r.ReadByte()
	zoomLevel, _ := r.ReadByte()
	isGprs, _ := r.ReadBool()
	width := int(r.ReadInt())
	height := int(r.ReadInt())
	isQwerty, _ := r.ReadBool()
	isTouch, _ := r.ReadBool()
	platform, _ := r.ReadUTF()
	versionIP := int(r.ReadInt())
	_, _ = r.ReadByte() // byte thừa
	lang, _ := r.ReadByte()
	userProvider := int(r.ReadInt())
	clientAgent, _ := r.ReadUTF()

	log.Printf("📱 Client info:")
	log.Printf("- Type: %d, Zoom: %d, GPRS: %v", clientType, zoomLevel, isGprs)
	log.Printf("- Screen: %dx%d, QWERTY: %v, Touch: %v", width, height, isQwerty, isTouch)
	log.Printf("- Platform: %s", platform)
	log.Printf("- VersionIP: %d, Lang: %d, Provider: %d", versionIP, lang, userProvider)
	log.Printf("- Agent: %s", clientAgent)

	// Parse IP từ RemoteAddr (vì có dạng "ip:port")
	remoteIP := strings.Split(s.Conn().RemoteAddr().String(), ":")[0]

	// Tạo client session model để lưu vào DB
	session := model.ClientSession{
		ClientType:   int16(clientType),
		ZoomLevel:    int16(zoomLevel),
		IsGprs:       isGprs,
		Width:        width,
		Height:       height,
		IsQwerty:     isQwerty,
		IsTouch:      isTouch,
		Platform:     platform,
		VersionIP:    versionIP,
		Language:     int16(lang),
		UserProvider: userProvider,
		ClientAgent:  clientAgent,
		RemoteAddr:   remoteIP,
	}

	if err := infra.DB.Create(&session).Error; err != nil {
		log.Printf("❌ Failed to insert client session: %v", err)
	} else {
		s.ClientSessionID = &session.ID
		log.Printf("✅ Client session saved (ID: %d, IP: %s)", session.ID, session.RemoteAddr)
	}
}
