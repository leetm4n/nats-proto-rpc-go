package runnable

import (
	"context"

	"github.com/leetm4n/nats-proto-rpc-go/pkg/encoder"
	"github.com/leetm4n/nats-proto-rpc-go/pkg/subject"
	"github.com/leetm4n/nats-proto-rpc-go/pkg/telemetry"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

type Runnable interface {
	Run(ctx context.Context) error
	GetNatsMicroService() micro.Service
}

type Hooks struct {
	ErrorHandler       micro.ErrHandler
	DoneHandler        micro.DoneHandler
	StartHandler       micro.DoneHandler
	EndpointAddHandler EndpointAddHandlerFn
}

type Options struct {
	NatsConnection      *nats.Conn
	Encoder             encoder.Encoder
	IsValidationEnabled bool
	ErrorEncoder        ErrorEncoderFn
	SubjectPrefix       string
	GetSubject          subject.GetSubjectFn
	Telemetry           telemetry.TelemetryOptions
	Hooks               Hooks
}

type EndpointAddHandlerFn func(service micro.Service, methodName, subject string)

type ErrorEncoderFn func(err error) (code string, description string)
