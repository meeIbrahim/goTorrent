package client

import (
	"crypto/rand"
	"net"
	"sync"
)

type Client struct {
	Id       [20]byte
	Listener *net.Listener
}

var (
	instance *Client
	once     sync.Once
)

func GetClient() *Client {
	once.Do(func() {
		instance = &Client{}
		uuid := make([]byte, 20)
		rand.Read(uuid)
		instance.Id = [20]byte(uuid)
		listener, err := net.Listen("tcp", "127.0.0.1:0")
		instance.Listener = &listener
		if err != nil {
			panic(err)
		}
	})
	return instance
}
