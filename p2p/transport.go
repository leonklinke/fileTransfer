package p2p

// transport would handle communication between peers, independently of the type (tcp, udp, sockets, etc).
type Transport interface{}

// Peer represent a remote node in the network.
type Peer interface{}
