package netvis

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sync"

	hdisc "github.com/KennyZ69/HdiscLib"
)

type UDPServer struct {
	conn *net.UDPConn
	mu   sync.Mutex
	enc  *json.Encoder
}

type udpWriter struct {
	conn *net.UDPConn
	addr *net.UDPAddr
}

func (w *udpWriter) Write(p []byte) (int, error) {
	return w.conn.WriteToUDP(p, w.addr)
}

func (s *UDPServer) Close() error {
	fmt.Println("\nClosing the connection...")
	return s.conn.Close()
}

func NewUDPServer(port int) (*UDPServer, error) {
	ownIP, _, err := hdisc.GetLocalNet()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Listening on %s:%d\n", ownIP.String(), port)

	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("0.0.0.0"),
		// IP: ownIP,
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		return nil, err
	}

	serv := &UDPServer{
		conn: conn,
		// mu: sync.Mutex{},
	}

	return serv, nil
}

// WaitForClient waits for the first message from client to start communicating
func (s *UDPServer) WaitForClient() error {
	buf := make([]byte, 1)
	fmt.Println("Waiting for a client...")
	for {

		n, addr, err := s.conn.ReadFromUDP(buf)
		if err != nil {
			// return err
			continue
		}
		if n > 0 {
			s.mu.Lock()
			// encoder always writing to the client
			s.enc = json.NewEncoder(&udpWriter{s.conn, addr})
			s.mu.Unlock()
			log.Printf("Client connected: %s\n", addr.String())
			break
		}
	}

	return nil
}

func (s *UDPServer) SendPacket(p PacketInfo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.enc == nil {
		return fmt.Errorf("no client connected yet")
	}
	return s.enc.Encode(p)
}
