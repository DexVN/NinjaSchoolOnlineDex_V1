 package sub_command

import (
	"nso-server/internal/net"
	"nso-server/internal/proto"
	"nso-server/internal/pkg/di"
	"nso-server/internal/pkg/logger"
)

type SubCommandRouter struct {
	Deps *di.Dependencies
}

func NewRouter(deps *di.Dependencies) *SubCommandRouter {
	return &SubCommandRouter{Deps: deps}
}

func (r *SubCommandRouter) Route(subCmd int8, msg *proto.Message, s *net.Session) {
	switch subCmd {
	default:
		logger.Warnf("⚠️ Unknown SubCommand SubCmd: %d", subCmd)
	}
}
