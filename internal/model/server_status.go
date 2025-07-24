package model

type ServerStatus int8

const (
	ServerOffline ServerStatus = iota // 0
	ServerOnline                      // 1
	ServerMaintenance                 // 2
)

func (s ServerStatus) String() string {
	switch s {
	case ServerOffline:
		return "Offline"
	case ServerOnline:
		return "Online"
	case ServerMaintenance:
		return "Maintenance"
	default:
		return "Unknown"
	}
}
