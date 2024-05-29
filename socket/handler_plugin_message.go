package socket

import (
	"github.com/paroxity/portal/socket/packet"
)

// PluginMessageHandler is responsible for handling the PluginMessage packet sent by servers.
type PluginMessageHandler struct{ requireAuth }

// Handle ...
func (*PluginMessageHandler) Handle(p packet.Packet, srv Server, c *Client) error {
	pk := p.(*packet.PluginMessage)
	srv.Handler().HandlePluginMessage(c, pk.Channel, pk.Message)
	return nil
}
