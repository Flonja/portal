package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// PluginMessage is sent by a connection or the proxy to have intercommunication between "custom" logic.
type PluginMessage struct {
	// Channel is the type of message the connection or proxy should be decoding.
	Channel string
	// Message is the message of the packet.
	Message []byte
}

// ID ...
func (pk *PluginMessage) ID() uint16 {
	return IDPluginMessage
}

// Marshal ...
func (pk *PluginMessage) Marshal(w *protocol.Writer) {
	w.String(&pk.Channel)
	w.ByteSlice(&pk.Message)
}

// Unmarshal ...
func (pk *PluginMessage) Unmarshal(r *protocol.Reader) {
	r.String(&pk.Channel)
	r.ByteSlice(&pk.Message)
}
