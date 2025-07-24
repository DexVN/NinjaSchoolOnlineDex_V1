package not_map

import (
	"nso-server/internal/net"
	"nso-server/internal/proto"
	"nso-server/internal/pkg/di"
	"nso-server/internal/pkg/logger"
)

type NotMapRouter struct {
	deps *di.Dependencies
}

func NewRouter(deps *di.Dependencies) *NotMapRouter {
	return &NotMapRouter{deps: deps}
}

func (r *NotMapRouter) Route(subCmd int8, msg *proto.Message, s *net.Session) {
	switch subCmd {

	// case proto.SubCmdExample:
	// 	   r.handleSomething(msg, s)

	default:
		logger.Warnf("⚠️ Unknown NotMap SubCmd: %d", subCmd)
	}
}
