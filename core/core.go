package core

import (
	"net/http"
	"sync"

	"github.com/castillobg/ping/brokers"
)

var pendingResponses = make([]http.ResponseWriter, 0)
var pendingResponsesLock = new(sync.Mutex)

func Listen(broker brokers.BrokerAdapter, pongs chan []byte, responses chan http.ResponseWriter) {
	go func() {
		// Listens for pong events
		for pong := range pongs {
			// If a pong arrives, writes to the oldest response and removes it from pendingResponses.
			pendingResponsesLock.Lock()
			pendingResLength := len(pendingResponses)
			if pendingResLength > 0 {
				pendingResponses[pendingResLength-1].Write(pong)
				pendingResponses = pendingResponses[:pendingResLength-1]
			}
			pendingResponsesLock.Unlock()
		}
	}()

	// Listens for new api call events.
	go func() {
		for response := range responses {
			// If there's a new api call (ping), append the ResponseWriter to pendingResponses, and
			// publish a ping to the broker.
			pendingResponsesLock.Lock()
			pendingResponses = append(pendingResponses, response)
			broker.Publish("ping", "pings")
			pendingResponsesLock.Unlock()
		}
	}()
}
