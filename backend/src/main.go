package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/KennyZ69/netvis"
)

func main() {
	filter := flag.String("f", "", "BPF filter (optional)")
	target := flag.String("t", "", "Target address to connect to")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// handling exit
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigc
		cancel()
	}()

	s, err := netvis.NewUDPStream(*target)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	if err := netvis.Snif(ctx, *filter, s); err != nil {
		log.Fatal(err)
	}
}
