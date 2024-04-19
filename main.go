package main

import (
	"log"

	"github.com/leonklinke/fileTransfer/p2p"
)

func main() {
	transport := p2p.NewTCPTransport(":3000")
	if err := transport.ListenTCP(); err != nil {
		log.Fatal(err)
	}

	select {}
}
