package not_login

import (
	"nso-server/internal/net"
	"nso-server/internal/pkg/di"
	"nso-server/internal/pkg/logger"
	"nso-server/internal/proto"
)

type NotLoginRouter struct {
	deps    *di.Dependencies
	reg     *RegisterHandler
	client  *ClientInfoHandler
}

func NewRouter(deps *di.Dependencies, reg *RegisterHandler, client *ClientInfoHandler) *NotLoginRouter {
	return &NotLoginRouter{
		deps:   deps,
		reg:    reg,
		client: client,
	}
}

func (r *NotLoginRouter) Route(subCmd int8, msg *proto.Message, s *net.Session) {
	switch subCmd {
	case proto.CmdConfirmAccount:
		r.reg.Handle(msg, s)
	case proto.CmdClientInfo:
		r.client.Handle(msg, s)
	default:
		logger.Warnf("⚠️ Unknown NotLogin SubCmd: %d", subCmd)
	}
}
