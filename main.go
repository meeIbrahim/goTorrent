package main

import (
	"crypto/rand"
	"fmt"
	torrentfile "goTorrent/torrentFile"
	"net"
)

func main() {
	torrent, err := torrentfile.GetTorrent("/home/ibrahim/Documents/debian.torrent")
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}

	defer listener.Close()
	addr := listener.Addr().(*net.TCPAddr)
	port := addr.Port
	uuid := make([]byte, 20)
	rand.Read(uuid)
	peers, err := torrent.RequestPeers([20]byte(uuid), uint16(port))
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(peers); i++ {
		fmt.Println(peers[i].IP)
	}
}
