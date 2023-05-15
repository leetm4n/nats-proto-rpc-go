package client

import (
	"time"

	"github.com/leetm4n/nats-proto-rpc-go/pkg/encoder"
	"github.com/leetm4n/nats-proto-rpc-go/pkg/subject"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
	"github.com/opentracing/opentracing-go"
)

type Options struct {
	NatsConnection      *nats.Conn
	Encoder             encoder.Encoder
	IsValidationEnabled bool
	GetSubject          subject.GetSubjectFn
	Timeout             time.Duration
	ErrorDecoder        ErrorDecoderFn
	Tracer              opentracing.Tracer
}

type ErrorDecoderFn func(code string, description string) error

func getErrorCodeAndDescriptionFromHeader(header nats.Header) (code string, description string) {
	errCode := header[micro.ErrorCodeHeader]
	if len(errCode) == 1 {
		code = errCode[0]
	}

	errDescription := header[micro.ErrorHeader]
	if len(errDescription) == 1 {
		description = errDescription[0]
	}

	return
}

func DecodeErrorFromHeaderWithDecoder(header nats.Header, errorDecoder ErrorDecoderFn) error {
	code, description := getErrorCodeAndDescriptionFromHeader(header)

	if code == "" {
		return nil
	}

	return errorDecoder(code, description)
}
