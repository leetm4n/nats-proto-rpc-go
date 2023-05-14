package runnable

import (
	"context"

	"github.com/leetm4n/nats-proto-rpc-go/pkg/encoder"
	"github.com/leetm4n/nats-proto-rpc-go/pkg/subject"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

type Runnable interface {
	Run(ctx context.Context) error
	GetNatsMicroService() *micro.Service
}

type Options struct {
	NatsConnection      *nats.Conn
	Encoder             encoder.Encoder
	IsValidationEnabled bool
	ErrorMapper         ErrorMapper
	ErrorHandler        micro.ErrHandler
	DoneHandler         micro.DoneHandler
	SubjectPrefix       string
	GetSubject          subject.GetSubjectFn
}

type ErrorMapper func(err error) (code string, description string)
