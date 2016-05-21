package brokers

import "errors"

var registeredFactories = make(map[string]AdapterFactory)

type Event struct {
	Type string
	Body interface{}
}

// AdapterFactory specifies a constructor for BrokerAdapter factories.
type AdapterFactory interface {
	// New builds a BrokerAdapter, which should be a client of a broker listening on the given
	// address, and should return a func to cleanup the connections opened.
	New(address string) (BrokerAdapter, func(), error)
}

// RegistryAdapter specifies the contract a broker adapter (kafka, rabbit) should follow.
type BrokerAdapter interface {
	Listen(queue string, messages chan []byte) error
	Publish(message, queue string) error
}

// Register registers an AdapterFactory for use.
func Register(rf AdapterFactory, name string) error {
	if _, ok := registeredFactories[name]; ok {
		// Should be unique (either "kafka", "rabbit", etc.)
		return errors.New("A broker with the name \"" + name + "\" was already registered.")
	}
	registeredFactories[name] = rf
	return nil
}

// LookUp returns an AdapterFactory registered with a given name.
func LookUp(name string) (AdapterFactory, bool) {
	registry, ok := registeredFactories[name]
	return registry, ok
}
