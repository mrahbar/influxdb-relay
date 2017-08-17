package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/mrahbar/influxdb-relay/relay"
)

var configJson string

func main() {
	flag.StringVar(&configJson, "config", "", "Configuration json object to use")
	flag.Parse()

	if configJson == "" {
		configJson = os.Getenv("CONFIG")

		if configJson == "" {
			fmt.Fprintln(os.Stderr, "Missing configuration file")
			flag.PrintDefaults()
			os.Exit(1)
		}
	}

	cfg, err := relay.LoadConfigJson(configJson)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Problem loading config file:", err)
	}

	if cfg.Debug {
		log.Println("[D] Debug is on")
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

	log.Println("Starting relays...")
	r.Run()
}
