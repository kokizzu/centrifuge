package centrifuge

import (
	"context"
)

// SessionEvent contains additional data for SessionHandler.
type SessionEvent struct {
	// ClientID that was generated by library for client connection.
	ClientID string
}

// SessionReply is a reply of SessionResover with instructions to library.
// Disconnect if provived allows to close connection, Data is a custom data
// to send in Session push.
type SessionReply struct {
	// Context allows to return new context most probobly derived from previous context.
	Context context.Context
	// Disconnect client.
	Disconnect *Disconnect
	// Data to send in Session push.
	Data Raw
}

// SessionHandler allows to resolve connection Session data in custom way.
type SessionHandler func(context.Context, Transport, SessionEvent) SessionReply

// AuthEvent contains fields related to auth event.
type AuthEvent struct {
	// ClientID that was generated by library for client connection.
	ClientID string
	// Token received from client as part of Connect Command.
	Token string
	// Data received from client as part of Connect Command.
	Data Raw
}

// AuthReply contains fields determining the reaction on auth event.
type AuthReply struct {
	// Context allows to return new context most probobly derived from previous context.
	Context context.Context
	// Error for connect command reply.
	Error *Error
	// Disconnect client.
	Disconnect *Disconnect
	// Credentials should be set if app wants to authenticate connection.
	// This field still optional as auth could be provided through HTTP middleware
	// or via JWT token.
	Credentials *Credentials
	// Data allows to set custom data in connect reply.
	Data Raw
}

// AuthHandler called when new client authenticates on server.
type AuthHandler func(context.Context, Transport, AuthEvent) AuthReply

// ConnectHandler called when new client connects to server.
type ConnectHandler func(context.Context, *Client)

// DisconnectEvent contains fields related to disconnect event.
type DisconnectEvent struct {
	Disconnect *Disconnect
}

// DisconnectReply contains fields determining the reaction on disconnect event.
type DisconnectReply struct{}

// DisconnectHandler called when client disconnects from server.
type DisconnectHandler func(DisconnectEvent) DisconnectReply

// SubscribeEvent contains fields related to subscribe event.
type SubscribeEvent struct {
	Channel string
}

// SubscribeReply contains fields determining the reaction on subscribe event.
type SubscribeReply struct {
	Error       *Error
	Disconnect  *Disconnect
	ExpireAt    int64
	ChannelInfo Raw
}

// SubscribeHandler called when client wants to subscribe on channel.
type SubscribeHandler func(SubscribeEvent) SubscribeReply

// UnsubscribeEvent contains fields related to unsubscribe event.
type UnsubscribeEvent struct {
	Channel string
}

// UnsubscribeReply contains fields determining the reaction on unsubscribe event.
type UnsubscribeReply struct {
}

// UnsubscribeHandler called when client unsubscribed from channel.
type UnsubscribeHandler func(UnsubscribeEvent) UnsubscribeReply

// PublishEvent contains fields related to publish event.
type PublishEvent struct {
	Channel string
	Data    Raw
	Info    *ClientInfo
}

// PublishReply contains fields determining the reaction on publish event.
type PublishReply struct {
	Error      *Error
	Disconnect *Disconnect
}

// PublishHandler called when client publishes into channel.
type PublishHandler func(PublishEvent) PublishReply

// RefreshEvent contains fields related to refresh event.
type RefreshEvent struct{}

// RefreshReply contains fields determining the reaction on refresh event.
type RefreshReply struct {
	ExpireAt int64
	Info     Raw
}

// RefreshHandler called when it's time to validate client connection and
// update it's expiration time.
type RefreshHandler func(RefreshEvent) RefreshReply

// SubRefreshEvent contains fields related to subscription refresh event.
type SubRefreshEvent struct {
	Channel string
}

// SubRefreshReply contains fields determining the reaction on
// subscription refresh event.
type SubRefreshReply struct {
	Expired  bool
	ExpireAt int64
	Info     Raw
}

// SubRefreshHandler called when it's time to validate client subscription to channel and
// update it's state if needed.
type SubRefreshHandler func(SubRefreshEvent) SubRefreshReply

// RPCEvent contains fields related to rpc request.
type RPCEvent struct {
	Data Raw
}

// RPCReply contains fields determining the reaction on rpc request.
type RPCReply struct {
	Error      *Error
	Disconnect *Disconnect
	Data       Raw
}

// RPCHandler must handle incoming command from client.
type RPCHandler func(RPCEvent) RPCReply

// MessageEvent contains fields related to message request.
type MessageEvent struct {
	Data Raw
}

// MessageReply contains fields determining the reaction on message request.
type MessageReply struct {
	Disconnect *Disconnect
}

// MessageHandler must handle incoming async message from client.
type MessageHandler func(MessageEvent) MessageReply
