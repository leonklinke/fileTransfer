package p2p

import (
	"fmt"
	"net"
)

type TCPPeer struct {
	conn net.Conn
	//outbound would be "true" if this peer dials and retrieve a conn.
	//outbound would be "false" if this peer accepts and returns a conn.
	outbound bool
}

type TCPTransport struct {
	address  string
	listener net.Listener
	peers    map[net.Addr]Peer
}

func NewTCPTransport(address string) *TCPTransport {
	return &TCPTransport{
		address: address,
		peers:   make(map[net.Addr]Peer),
	}
}

func (t *TCPTransport) ListenTCP() error {
	var err error

	t.listener, err = net.Listen("tcp", t.address)
	if err != nil {
		return err
	}

	go t.acceptConnection()

	return nil
}

func (t *TCPTransport) acceptConnection() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("error accepting connection: %s\n", err)
			return
		}

		go t.handleConnection(conn)
	}
}

func (t *TCPTransport) handleConnection(conn net.Conn) {
	fmt.Printf("handling connection from %s\n", conn.RemoteAddr())

	peer := TCPPeer{
		conn:     conn,
		outbound: false,
	}

	t.peers[conn.RemoteAddr()] = &peer
}
