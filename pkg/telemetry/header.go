package telemetry

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

type NatsHeaderCarrier nats.Header

func (nhc NatsHeaderCarrier) Get(key string) string {
	val := nhc[key]
	if len(val) == 0 {
		return ""
	}
	return val[0]
}

func (nhc NatsHeaderCarrier) Set(key string, value string) {
	nhc[key] = []string{value}
}

func (nhc NatsHeaderCarrier) Keys() []string {
	var keys []string
	for key := range nhc {
		keys = append(keys, key)
	}

	return keys
}

func (nhc NatsHeaderCarrier) GetNatsHeader() nats.Header {
	natsHeader := nats.Header{}
	for key, val := range nhc {
		natsHeader[key] = val
	}
	return natsHeader
}

func NatsHeaderCarrierFromNatsHeader(header micro.Headers) NatsHeaderCarrier {
	natsHeaderCarrier := NatsHeaderCarrier{}
	for key, val := range header {
		natsHeaderCarrier[key] = val
	}
	return natsHeaderCarrier
}
