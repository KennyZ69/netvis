package netvis

import (
	"encoding/json"
	"net"
)

type UDPStream struct {
	conn net.Conn
	enc  *json.Encoder
}

func NewUDPStream(addr string) (*UDPStream, error) {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		return nil, err
	}

	return &UDPStream{
		conn: conn,
		enc:  json.NewEncoder(conn),
	}, nil
}

func (u *UDPStream) SendPacket(p PacketInfo) error {
	return u.enc.Encode(p)
}

func (u *UDPStream) Close() error {
	return u.conn.Close()
}

// func sendUDP(p PacketInfo, addr string) (int, error) {
// 	conn, err := net.Dial("udp", addr)
// 	if err != nil {
// 		return 0, fmt.Errorf("Error getting the connection for udp: %v\n", err)
// 	}
// 	defer conn.Close()
//
// 	data, err := json.Marshal(p)
// 	if err != nil {
// 		return 0, err
// 	}
//
// 	n, err := conn.Write(data)
//
// 	return n, err
// }
