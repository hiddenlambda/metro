package worker

import (
	"github.com/razorpay/metro/pkg/messagebroker"
)

// Config for pushconsumer
type Config struct {
	Broker     Broker
	Interfaces struct {
		API NetworkInterfaces
	}
}

// Broker Config (Kafka/Pulsar)
type Broker struct {
	Variant      string // kafka or pulsar
	BrokerConfig messagebroker.BrokerConfig
}

// NetworkInterfaces contains all exposed interfaces
type NetworkInterfaces struct {
	GrpcServerAddress string
	HTTPServerAddress string
}
