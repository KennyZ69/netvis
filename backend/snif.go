package netvis

import (
	"context"
	"fmt"

	hdisc "github.com/KennyZ69/HdiscLib"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// func Snif(ctx context.Context, filter string, stream *UDPStream) error {
// Snif for network packets and handle decoded packets using own handler func;
func Snif(ctx context.Context, filter string, handler func(p PacketInfo)) error {
	ifi, err := hdisc.LocalIface()
	if err != nil {
		return fmt.Errorf("Error getting local ifi: %v\n", err)
	}

	handle, err := pcap.OpenLive(ifi.Name, 1600, true, pcap.BlockForever)
	if err != nil {
		return fmt.Errorf("Error getting the pcap handle: %v\n", err)
	}
	defer handle.Close()

	if filter != "" {
		if err = handle.SetBPFFilter(filter); err != nil {
			return err
		}
	}

	src := gopacket.NewPacketSource(handle, handle.LinkType())
	for {
		select {
		case <-ctx.Done():
			return nil
		case p := <-src.Packets():
			if p == nil {
				continue
			}
			data := decodePacket(p)
			handler(data)
		}
	}
}
