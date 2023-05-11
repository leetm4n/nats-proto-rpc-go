package runnable

import (
	"context"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

type Runnable interface {
	Run(ctx context.Context) error
	GetNatsMicroService() *micro.Service
}

type Options struct {
	Encoder             nats.Encoder
	IsValidationEnabled bool
	ErrorMapper         ErrorMapper
	SubjectPrefix       *string
}

type ErrorMapper func(err error) (code string, description string)
