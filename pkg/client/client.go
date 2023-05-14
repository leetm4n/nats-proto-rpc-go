package client

import (
	"time"

	"github.com/leetm4n/nats-proto-rpc-go/pkg/encoder"
	"github.com/leetm4n/nats-proto-rpc-go/pkg/subject"
	"github.com/nats-io/nats.go"
)

type Options struct {
	NatsConnection      *nats.Conn
	Encoder             encoder.Encoder
	IsValidationEnabled bool
	GetSubject          subject.GetSubjectFn
	Timeout             time.Duration
}
