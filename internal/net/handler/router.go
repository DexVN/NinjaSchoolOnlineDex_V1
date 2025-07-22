package handler

import (
	"log"

	"nso-server/internal/net"
	"nso-server/internal/proto"
)

func RouteMessage(msg *proto.Message, session *net.Session) {
	switch msg.Command {
	case proto.CmdNotLogin: // -29
		subCmd, err := msg.Reader().ReadByte()
		if err != nil {
			log.Println("❌ Failed to read sub command:", err)
			return
		}
		log.Printf("🔄 SubCommand (not login): %d", int8(subCmd))

		switch int8(subCmd) {
		case proto.CmdClientInfo:
			HandleClientInfo(msg, session)
		case proto.CmdLogin:
			// HandleLogin(msg, session)
		default:
			log.Printf("⚠️ Unknown sub-command: %d", subCmd)
		}

	default:
		log.Printf("⚠️ Unknown command: %d", msg.Command)
	}
}
