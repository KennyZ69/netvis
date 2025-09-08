package netvis

import (
	"encoding/json"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type NetEvent struct {
	SrcIP string `json:"src"`
	DstIP string `json:"dst"`
	Proto string `json:"proto"`
	Len   int    `json:"len"`
}

type PacketInfo struct {
	Timestamp  time.Time `json:"timestamp"`
	SrcIP      string    `json:"src_ip,omitempty"`
	SrcMAC     string    `json:"src_mac,omitempty"`
	DestIP     string    `json:"dest_ip,omitempty"`
	DestMAC    string    `json:"dest_mac,omitempty"`
	Len        int       `json:"len"`
	Prot       string    `json:"prot,omitempty"`
	SrcPort    uint16    `json:"src_port,omitempty"`
	DestPort   uint16    `json:"dest_port,omitempty"`
	Payload    []byte    `json:"payload"`
	PayloadLen int       `json:"payload_len"`
}

func (p PacketInfo) JSON() ([]byte, error) {
	return json.Marshal(p)
}

func decodePacket(p gopacket.Packet) PacketInfo {
	info := PacketInfo{
		Timestamp: p.Metadata().Timestamp,
		Len:       len(p.Data()),
	}

	if appLayer := p.ApplicationLayer(); appLayer != nil {
		info.PayloadLen = len(p.ApplicationLayer().Payload())
	}

	if eth := p.Layer(layers.LayerTypeEthernet); eth != nil {
		e := eth.(*layers.Ethernet)
		info.SrcMAC = e.SrcMAC.String()
		info.DestMAC = e.DstMAC.String()
	}

	if ip4 := p.Layer(layers.LayerTypeIPv4); ip4 != nil {
		ip := ip4.(*layers.IPv4)
		info.SrcIP = ip.SrcIP.String()
		info.DestIP = ip.DstIP.String()
		info.Prot = ip.Protocol.String()
	} else if ip6 := p.Layer(layers.LayerTypeIPv6); ip6 != nil {
		ip := ip6.(*layers.IPv6)
		info.SrcIP = ip.SrcIP.String()
		info.DestIP = ip.DstIP.String()
		info.Prot = ip.NextHeader.String()
	}

	if tcp := p.Layer(layers.LayerTypeTCP); tcp != nil {
		t := tcp.(*layers.TCP)
		info.Prot = "TCP"
		info.SrcPort = uint16(t.SrcPort)
		info.DestPort = uint16(t.DstPort)
	} else if udp := p.Layer(layers.LayerTypeUDP); udp != nil {
		u := udp.(*layers.UDP)
		info.Prot = "UDP"
		info.SrcPort = uint16(u.SrcPort)
		info.DestPort = uint16(u.DstPort)
	}

	if icmp4 := p.Layer(layers.LayerTypeICMPv4); icmp4 != nil {
		// icmp := icmp4.(*layers.ICMPv4)
		info.Prot = "ICMPv4"
	} else if icmp6 := p.Layer(layers.LayerTypeICMPv6); icmp6 != nil {
		// icmp := icmp6.(*layers.ICMPv4)
		info.Prot = "ICMPv6"
	}

	if arp := p.Layer(layers.LayerTypeARP); arp != nil {
		info.Prot = "ARP"
	}

	return info
}
