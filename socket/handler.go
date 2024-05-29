package socket

// Handler handles events that are called by a client.
type Handler interface {
	// HandlePluginMessage handles a packet that's sent by the client to send custom "logic" data.
	HandlePluginMessage(c *Client, channel string, message []byte)
}

// NopHandler implements the Handler interface but does not execute any code when an event is called. The
// default handler of sessions is set to NopHandler.
// Users may embed NopHandler to avoid having to implement each method.
type NopHandler struct{}

// Compile time check to make sure NopHandler implements Handler.
var _ Handler = (*NopHandler)(nil)

// HandlePluginMessage ...
func (NopHandler) HandlePluginMessage(*Client, string, []byte) {}
