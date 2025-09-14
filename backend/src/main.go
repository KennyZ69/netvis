package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/KennyZ69/netvis"
)

func main() {
	filter := flag.String("f", "", "BPF filter (optional)")
	port := flag.Int("p", 999, "Target port to serve on")
	flag.Parse()

	serv, err := netvis.NewUDPServer(*port)
	if err != nil {
		log.Fatalf("Error creating the backend server: %v\n", err)
	}
	defer serv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigc
		serv.Close()
		cancel()
		os.Exit(0)
	}()

	if err = serv.WaitForClient(); err != nil {
		serv.Close()
		log.Fatalf("Error getting client connection: %v\n", err)
	}

	if err := netvis.Snif(ctx, *filter, func(p netvis.PacketInfo) {
		if err := serv.SendPacket(p); err != nil {
			fmt.Printf("sending error: %s\n", err)
		}
	}); err != nil {
		fmt.Printf("sniffing error: %s\n", err)
	}
}
