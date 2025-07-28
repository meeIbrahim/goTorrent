package torrentfile

import (
	"goTorrent/peer"
	"net/http"
	"time"

	"github.com/jackpal/bencode-go"
)

type bencodeTrackResp struct {
	Interval int    `bencode:"interval"`
	Peers    string `bencode:"peers"`
}

func (t *Torrent) RequestPeers(peerID [20]byte, port uint16) ([]peer.Peer, error) {
	url, err := t.buildTackerUrl(peerID, port)
	if err != nil {
		return nil, err
	}
	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // Defer close until function returns
	trackerResp := bencodeTrackResp{}
	err = bencode.Unmarshal(resp.Body, &trackerResp)
	if err != nil {
		return nil, err
	}
	return peer.Unmarshal([]byte(trackerResp.Peers))

}
