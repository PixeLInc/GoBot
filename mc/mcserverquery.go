package mc

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"
	"net"
	"strconv"
)

const latestProtcol = 0x47

type Server struct {
	Version struct {
		Name     string
		Protocol int
	} `json:"version"`
	Players struct {
		Max    int `json:"max"`
		Online int `json:"online"`
		Sample []map[string]string
	} `json:"players"`
	Description interface{} `json:"description"`
	FavIcon     string      `json:"favicon"`
}

func PingServer(conn net.Conn, host string) (*Server, error) {
	if err := SendHandshake(conn, host); err != nil {
		return nil, err
	}

	if err := SendStatusRequest(conn); err != nil {
		return nil, err
	}

	pong, err := ReadPong(conn)
	if err != nil {
		return nil, err
	}

	return pong, nil
}

func makePacket(pl *bytes.Buffer) *bytes.Buffer {
	var buf bytes.Buffer

	buf.Write(encodeVarint(uint64(len(pl.Bytes()))))

	buf.Write(pl.Bytes())

	return &buf
}

func SendHandshake(conn net.Conn, host string) error {
	pl := &bytes.Buffer{}

	// Handshake Packet
	pl.WriteByte(0x00)

	// Protcol Version
	pl.WriteByte(latestProtcol)

	// Server Address
	host, port, err := net.SplitHostPort(host)
	if err != nil {
		panic(err)
	}

	pl.Write(encodeVarint(uint64(len(host))))
	pl.WriteString(host)

	// Port
	iPort, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	binary.Write(pl, binary.BigEndian, int16(iPort))

	// Status Update
	pl.WriteByte(0x01)

	if _, err := makePacket(pl).WriteTo(conn); err != nil {
		return errors.New("Cannot Write the Handshake!")
	}

	return nil
}

func SendStatusRequest(conn net.Conn) error {
	pl := &bytes.Buffer{}

	// Send Request Zero
	pl.WriteByte(0x00)

	if _, err := makePacket(pl).WriteTo(conn); err != nil {
		return errors.New("Cannot write send status req.")
	}

	return nil
}

// https://code.google.com/p/goprotobuf/source/browse/proto/encode.go#83
func encodeVarint(x uint64) []byte {
	var buf [10]byte
	var n int
	for n = 0; x > 127; n++ {
		buf[n] = 0x80 | uint8(x&0x7F)
		x >>= 7
	}
	buf[n] = uint8(x)
	n++
	return buf[0:n]
}

func ReadPong(rd io.Reader) (*Server, error) {
	r := bufio.NewReader(rd)
	nl, err := binary.ReadUvarint(r)
	if err != nil {
		return nil, errors.New("canot read length")
	}

	pl := make([]byte, nl)
	if _, err := io.ReadFull(r, pl); err != nil {
		return nil, errors.New("Cant read length")
	}

	_, n := binary.Uvarint(pl)
	if n <= 0 {
		return nil, errors.New("cant read id")
	}

	_, n2 := binary.Uvarint(pl[n:])
	if n2 <= 0 {
		return nil, errors.New("cant read string var")
	}

	var serv Server
	if err := json.Unmarshal(pl[n+n2:], &serv); err != nil {
		return nil, errors.New("cant read json")
	}

	return &serv, nil
}
