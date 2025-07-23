package main

import (
	"fmt"
	torrentfile "goTorrent/torrentFile"
)

func main() {
	torrent, err := torrentfile.GetTorrent("/home/ibrahim/Documents/debian.torrent")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", torrent)
}
