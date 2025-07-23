package not_login

import (
	"strings"

	logger "nso-server/internal/infra"
	"nso-server/internal/model"
	"nso-server/internal/net"
	"nso-server/internal/proto"
)

func HandleClientInfo(msg *proto.Message, s *net.Session) {
	r := msg.Reader()

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

	logger.Log.Info("📱 Client info:")
	logger.Log.Infof("- Type: %d, Zoom: %d, GPRS: %v", clientType, zoomLevel, isGprs)
	logger.Log.Infof("- Screen: %dx%d, QWERTY: %v, Touch: %v", width, height, isQwerty, isTouch)
	logger.Log.Infof("- Platform: %s", platform)
	logger.Log.Infof("- VersionIP: %d, Lang: %d, Provider: %d", versionIP, lang, userProvider)
	logger.Log.Infof("- Agent: %s", clientAgent)

	remoteIP := strings.Split(s.Conn().RemoteAddr().String(), ":")[0]

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

	if err := logger.DB.Create(&session).Error; err != nil {
		logger.Log.WithError(err).Error("❌ Failed to insert client session")
	} else {
		s.ClientSessionID = &session.ID
		logger.Log.Infof("✅ Client session saved (ID: %d, IP: %s)", session.ID, session.RemoteAddr)
	}
}
