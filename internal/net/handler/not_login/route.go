package not_login

import (
	"log"
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
		log.Printf("⚠️ Unknown NotLogin SubCmd: %d", subCmd)
	}
}
