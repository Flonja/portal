package packet

import (
	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// SendMessageRequest is sent by a connection to send a message to player(s).
type SendMessageRequest struct {
	// PlayerUUIDs are the UUIDs of the players that'll receive the message.
	PlayerUUIDs []uuid.UUID
	// Message is the message of the packet.
	Message string
}

// ID ...
func (pk *SendMessageRequest) ID() uint16 {
	return IDSendMessageRequest
}

// Marshal ...
func (pk *SendMessageRequest) Marshal(w *protocol.Writer) {
	protocol.FuncSlice(w, &pk.PlayerUUIDs, w.UUID)
	w.String(&pk.Message)
}

// Unmarshal ...
func (pk *SendMessageRequest) Unmarshal(r *protocol.Reader) {
	protocol.FuncSlice(r, &pk.PlayerUUIDs, r.UUID)
	r.String(&pk.Message)
}
