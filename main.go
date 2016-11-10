package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/mrahbar/influxdb-relay/relay"
)

var (
	configJson = flag.String("config", "", "Configuration json object to use")
)

func main() {
	flag.Parse()

	if *configJson == "" {
		fmt.Fprintln(os.Stderr, "Missing configuration file")
		flag.PrintDefaults()
		os.Exit(1)
	}

	cfg, err := relay.LoadConfigJson(*configJson)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Problem loading config file:", err)
	}

	r, err := relay.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	go func() {
		<-sigChan
		r.Stop()
	}()

	log.Println("starting relays...")
	r.Run()
}
