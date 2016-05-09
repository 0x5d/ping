package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/castillobg/ping/api"
	"github.com/castillobg/ping/brokers"
	_ "github.com/castillobg/ping/brokers/rabbit"
)

func main() {
	brokerAddr := flag.String("address", "localhost:9092", "The broker's address.")
	brokerName := flag.String("broker", "rabbit", "The broker to be used. Currently supported: rabbit.")
	port := flag.Int("port", 8080, "The port where the api will listen.")
	flag.Parse()

	brokerFactory, ok := brokers.LookUp(*brokerName)
	if !ok {
		fmt.Printf("No broker with name \"%s\" found.\n", *brokerName)
		os.Exit(1)
	}
	broker, err := brokerFactory.New(*brokerAddr)
	if err != nil {
		fmt.Printf("Error initializing broker client for \"%s\": %s\n", *brokerName, err.Error())
		os.Exit(1)
	}

	go api.Listen(*port)

	fmt.Printf("ping is running on port: %d\n", port)

	broker.Start()
}
