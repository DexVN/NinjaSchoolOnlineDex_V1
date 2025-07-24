package handler

import (
	"nso-server/internal/net"
	"nso-server/internal/net/handler/not_login"
	"nso-server/internal/net/handler/not_map"
	"nso-server/internal/net/handler/sub_command"
	"nso-server/internal/proto"
	"nso-server/internal/pkg/logger"
)

type SubRouter interface {
	Route(subCmd int8, msg *proto.Message, s *net.Session)
}

type Router struct {
	commandRouters map[int8]SubRouter
}

func NewRouter(
	notLogin *not_login.NotLoginRouter,
	notMap *not_map.NotMapRouter,
	subCommand *sub_command.SubCommandRouter,
) *Router {
	return &Router{
		commandRouters: map[int8]SubRouter{
			proto.CmdNotLogin:   notLogin,
			proto.CmdNotMap:     notMap,
			proto.CmdSubCommand: subCommand,
		},
	}
}

func (r *Router) Handle(msg *proto.Message, s *net.Session) {
	subCmd, _ := msg.Reader().ReadInt8()
	if router, ok := r.commandRouters[msg.Command]; ok {
		router.Route(subCmd, msg, s)
	} else {
		logger.Warnf("⚠️ Unknown command: %d", msg.Command)
	}
}
