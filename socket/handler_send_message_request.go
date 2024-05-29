package socket

import (
	"github.com/paroxity/portal/socket/packet"
	minecraftpacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// SendMessageRequestHandler is responsible for handling the SendMessageRequest packet sent by servers.
type SendMessageRequestHandler struct{ requireAuth }

// Handle ...
func (*SendMessageRequestHandler) Handle(p packet.Packet, srv Server, _ *Client) error {
	pk := p.(*packet.SendMessageRequest)
	for _, u := range pk.PlayerUUIDs {
		s, ok := srv.SessionStore().Load(u)
		if !ok {
			continue
		}

		_ = s.Conn().WritePacket(&minecraftpacket.Text{
			TextType: minecraftpacket.TextTypeRaw,
			Message:  pk.Message,
		})
	}

	return nil
}
