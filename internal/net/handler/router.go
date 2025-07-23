package handler

import (
	"log"

	"nso-server/internal/net"
	"nso-server/internal/proto"

	notlogin "nso-server/internal/net/handler/not_login"
	subcmd "nso-server/internal/net/handler/sub_command"
)

func RouteMessage(msg *proto.Message, session *net.Session) {
	switch msg.Command {
	case proto.CmdNotLogin:
		subCmd, _ := msg.Reader().ReadInt8()
		notlogin.RouteNotLogin(subCmd, msg, session)
	case proto.CmdSubCommand:
		subCmd, _ := msg.Reader().ReadInt8()
		subcmd.RouteSubCommand(subCmd, msg, session)
	default:
		log.Printf("⚠️ Unknown command: %d", msg.Command)
	}
}
