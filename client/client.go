package client

import (
	"crypto/rand"
	"net"
)

type Client struct {
	id       [20]byte
	listener *net.Listener
}

func GetClient() *Client {
	client := Client{}
	uuid := make([]byte, 20)
	rand.Read(uuid)
	client.id = [20]byte(uuid)
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	client.listener = &listener
	if err != nil {
		panic(err)
	}
	return &client
}
