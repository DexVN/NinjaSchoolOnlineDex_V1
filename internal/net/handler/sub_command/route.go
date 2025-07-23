package sub_command

import (
	"nso-server/internal/proto"
	"nso-server/internal/net"
)

func RouteSubCommand(subCmd int8, msg *proto.Message, s *net.Session) {
	switch subCmd {
	default:
		// log subCmd chưa được xử lý
	}
}
