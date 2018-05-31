package netmsg

type Kind byte

const (
	MsgUnknown          Kind = 0 + iota
	MsgConnectResponse       = 1
	MsgUpdatePlayer          = 2
	MsgDisconnectPlayer      = 3
)

var kindToString = []string{
	MsgUnknown:          "MsgUnknown",
	MsgConnectResponse:  "MsgConnectResponse",
	MsgUpdatePlayer:     "MsgUpdatePlayer",
	MsgDisconnectPlayer: "MsgDisconnectPlayer",
}

func (kind Kind) String() string {
	kindAsInt := int(kind)
	if kindAsInt >= 0 && kindAsInt < len(kindToString) {
		return kindToString[kind]
	}
	return "MsgUnknown"
}
