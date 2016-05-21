package core

import (
	"sync"

	"github.com/castillobg/ping/brokers"
)

var waitingForPong = make([]chan []byte, 0)
var pongListenersLock = new(sync.Mutex)

func Listen(broker brokers.BrokerAdapter, pongs chan []byte, pongListeners chan chan []byte) {
	go func() {
		// Listens for pong events
		for pong := range pongs {
			// If a pong arrives, writes to the oldest response and removes it from pendingResponses.
			pongListenersLock.Lock()
			listenersLength := len(waitingForPong)
			if listenersLength > 0 {
				waitingForPong[0] <- pong
				waitingForPong = waitingForPong[1:listenersLength]
			}
			pongListenersLock.Unlock()
		}
	}()

	// Listens for new api call events.
	go func() {
		for listener := range pongListeners {
			// If there's a new api call (ping), append the ResponseWriter to pendingResponses, and
			// publish a ping to the broker.
			pongListenersLock.Lock()
			waitingForPong = append(waitingForPong, listener)
			broker.Publish("ping", "pings")
			pongListenersLock.Unlock()
		}
	}()
}
