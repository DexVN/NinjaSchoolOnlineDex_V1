package not_login

import (
	logger "nso-server/internal/infra"
	"nso-server/internal/net"
	"nso-server/internal/proto"
)

func RouteNotLogin(subCmd int8, msg *proto.Message, s *net.Session) {
	switch subCmd {
	case proto.CmdLogin:
		HandleLogin(msg, s)
	case proto.CmdClientInfo:
		HandleClientInfo(msg, s)
	case proto.CmdComfirmAccount:
		HandleRegister(msg, s)
	default:
		logger.Log.Warnf("⚠️ Unknown NotLogin SubCmd: %d", subCmd)
	}
}
