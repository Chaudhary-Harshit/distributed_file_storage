package p2p

// This represents the remote node
type Peer interface{}

// This reperesents the kind of communication protocol we want to use to be able to communicate with the remote node.
// It can be TCP, UDP, Websockets, etc.
type Transport interface {
	ListenAndAccept() error
}
