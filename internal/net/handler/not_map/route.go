package sub_command

import (
	"nso-server/internal/net"
	"nso-server/internal/proto"
)

func RouteNotMap(subCmd int8, msg *proto.Message, s *net.Session) {
	switch subCmd {
	default:
		// log subCmd chưa được xử lý
	}
}
