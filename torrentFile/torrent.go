package torrentfile

import (
	"bytes"
	"crypto/sha1"
	"io"
	"net/url"
	"os"
	"strconv"

	"github.com/jackpal/bencode-go"
)

type torrentInfo struct {
	Pieces      string `bencode:"pieces"`
	PieceLength int    `bencode:"piece length"`
	Length      int    `bencode:"length"`
	Name        string `bencode:"name"`
}
type torrentFile struct {
	Announce string       `bencode:"announce"`
	Comment  string       `bencode:"comment"`
	Info     *torrentInfo `becode:"info"`
}
type Torrent struct {
	Announce    string
	InfoHash    [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length      int
	Name        string
}

func (file *torrentFile) getTorrent() (*Torrent, error) {
	torrent := Torrent{}
	torrent.Name = file.Info.Name
	torrent.Announce = file.Announce
	torrent.Length = file.Info.Length
	torrent.PieceLength = file.Info.PieceLength
	byte, err := hashInfo(file.Info)
	if err != nil {
		return nil, err
	}
	torrent.InfoHash = byte
	return &torrent, nil
}

func (t *Torrent) buildTackerUrl(peerID [20]byte, port uint16) (string, error) {
	base, err := url.Parse(t.Announce)
	if err != nil {
		return "", err
	}
	params := url.Values{
		"info_hash":  []string{string(t.InfoHash[:])},
		"peer_id":    []string{string(peerID[:])},
		"port":       []string{strconv.Itoa(int(port))},
		"uploaded":   []string{"0"},
		"downloaded": []string{"0"},
		"compact":    []string{"1"},
		"left":       []string{strconv.Itoa(t.Length)},
	}
	base.RawQuery = params.Encode()
	return base.String(), nil
}

func openFile(r io.Reader) (*torrentFile, error) {
	torrentFile := torrentFile{}
	err := bencode.Unmarshal(r, &torrentFile)
	if err != nil {
		return nil, err
	}
	return &torrentFile, nil
}

func hashInfo(info *torrentInfo) ([20]byte, error) {
	var buffer bytes.Buffer
	err := bencode.Marshal(&buffer, info)
	if err != nil {
		return [20]byte{}, err
	}
	hash := sha1.Sum(buffer.Bytes())
	return hash, nil
}

func GetTorrent(path string) (*Torrent, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	torrentFile, err := openFile(bytes.NewReader(content))
	torrent, err := torrentFile.getTorrent()
	if err != nil {
		return nil, err
	}
	return torrent, nil
}
