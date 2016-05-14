package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/castillobg/ping/api"
	"github.com/castillobg/ping/brokers"
	_ "github.com/castillobg/ping/brokers/rabbit"
	"github.com/castillobg/ping/core"
)

func main() {
	brokerAddr := flag.String("address", "localhost:5672", "The broker's address.")
	brokerName := flag.String("broker", "rabbit", "The broker to be used. Currently supported: rabbit.")
	port := flag.Int("port", 8080, "The port where the api will listen.")
	flag.Parse()

	brokerFactory, ok := brokers.LookUp(*brokerName)
	if !ok {
		fmt.Printf("No broker with name \"%s\" found.\n", *brokerName)
		os.Exit(1)
	}
	broker, cleanup, err := brokerFactory.New(*brokerAddr)
	defer cleanup()
	if err != nil {
		log.Printf("Error initializing broker client for \"%s\": %s\n", *brokerName, err.Error())
		os.Exit(1)
	}
	messages := make(chan []byte)
	err = broker.Listen("pongs", messages)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	responses := make(chan http.ResponseWriter)
	core.Listen(broker, messages, responses)

	log.Printf("ping is running on port: %d\n", *port)
	api.Listen(*port, responses)
}
