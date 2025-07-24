package not_login

import (
	"nso-server/internal/net"
)

func sendCharacterList(s *net.Session, accountID int) {
	// if s.ServerID == nil {
	// 	logger.Warn("❌ Cannot send character list: missing ServerID in session")
	// 	return
	// }

	// var characters []model.Character
	// err := database.DB.
	// 	Where("account_id = ? AND server_id = ?", accountID, *s.ServerID).
	// 	Order("slot_index ASC").
	// 	Find(&characters).Error
	// if err != nil {
	// 	logger.WithError(err).Error("❌ Failed to load characters")
	// 	return
	// }

	// w := proto.NewWriter()

	// w.WriteByte(byte(len(characters)))
	// w.WriteInt8(proto.CmdSelectPlayer)

	// for _, char := range characters {
	// 	w.WriteByte(byte(char.Gender))
	// 	w.WriteUTF(char.Name)
	// 	w.WriteUTF(char.ClassName)
	// 	w.WriteByte(byte(char.Level))
	// 	w.WriteInt16(char.PartHead)

	// 	wp := char.PartWeapon
	// 	if wp == -1 {
	// 		wp = 15
	// 	}
	// 	w.WriteInt16(wp)

	// 	body := char.PartBody
	// 	if body == -1 {
	// 		body = int16(map[int]int{0: 10, 1: 1}[int(char.Gender)])
	// 	}
	// 	w.WriteInt16(body)

	// 	leg := char.PartLeg
	// 	if leg == -1 {
	// 		leg = int16(map[int]int{0: 9, 1: 0}[int(char.Gender)])
	// 	}
	// 	w.WriteInt16(leg)
	// }

	// s.SendMessageWithCommand(proto.CmdNotMap, w)
	// logger.Infof("📤 Sent %d characters (account=%d, server=%d)", len(characters), accountID, *s.ServerID)
}
