package main

import (
	"crypto/rand"
	"fmt"
	"goTorrent/client"
	"goTorrent/peer"
	torrentfile "goTorrent/torrentFile"
	"io"
	"net"
	"time"
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
		thePeer := peers[i]
		conn, err := net.DialTimeout("tcp", thePeer.GetAddress(), 3*time.Second)
		if err != nil {
			continue
		}
		handshakeMessage := peer.Handshake{}
		handshakeMessage.Pstr = "BitTorrent protocol"
		handshakeMessage.PeerID = client.GetClient().Id
		handshakeMessage.InfoHash = torrent.InfoHash

		handshakeSerialized := handshakeMessage.Serialize()
		conn.Write(handshakeSerialized)
		resp, err := peer.ReadHandShake(conn)
		fmt.Printf("Output %+v \n", resp)
		if err != nil {
			if err == io.EOF {
				// Resp
			} else {
				message := fmt.Errorf("Error in Response\n%v\n", err)
				fmt.Print(message)
				continue
			}
		}
	}
}
