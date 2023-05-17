package client

import (
	"time"

	"github.com/leetm4n/nats-proto-rpc-go/pkg/encoder"
	"github.com/leetm4n/nats-proto-rpc-go/pkg/subject"
	"github.com/leetm4n/nats-proto-rpc-go/pkg/telemetry"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

type Options struct {
	NatsConnection      *nats.Conn
	Encoder             encoder.Encoder
	IsValidationEnabled bool
	GetSubject          subject.GetSubjectFn
	Timeout             time.Duration
	ErrorDecoder        ErrorDecoderFn
	Telemetry           telemetry.Options
}

type ErrorDecoderFn func(code string, description string) error

func getErrorCodeAndDescriptionFromHeader(header nats.Header) (string, string) {
	var code, description string

	errCode := header[micro.ErrorCodeHeader]
	if len(errCode) == 1 {
		code = errCode[0]
	}

	errDescription := header[micro.ErrorHeader]
	if len(errDescription) == 1 {
		description = errDescription[0]
	}

	return code, description
}

func DecodeErrorFromHeaderWithDecoder(header nats.Header, errorDecoder ErrorDecoderFn) error {
	code, description := getErrorCodeAndDescriptionFromHeader(header)

	if code == "" {
		return nil
	}

	return errorDecoder(code, description)
}
