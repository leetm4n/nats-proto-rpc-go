package client

import (
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/encoders/builtin"
)

type WithEncoder interface {
	SetEncoder(nats.Encoder)
	GetEncoder() nats.Encoder
}

type WithClientValidation interface {
	SetClientValidation(bool)
	IsClientValidationEnabled() bool
}

type WithTimeout interface {
	SetTimeout(time.Duration)
	GetTimeout() time.Duration
}

const DefaultTimeout = time.Second * 10

var _ WithEncoder = (*BaseClient)(nil)
var _ WithClientValidation = (*BaseClient)(nil)
var _ WithTimeout = (*BaseClient)(nil)

type BaseClient struct {
	encoder               nats.Encoder
	isClientValidationSet bool
	timeout               time.Duration
}

func (base *BaseClient) SetEncoder(encoder nats.Encoder) {
	base.encoder = encoder
}

func (base *BaseClient) GetEncoder() nats.Encoder {
	return base.encoder
}

func (base *BaseClient) SetClientValidation(isSet bool) {
	base.isClientValidationSet = isSet
}

func (base *BaseClient) IsClientValidationEnabled() bool {
	return base.isClientValidationSet
}

func (base *BaseClient) SetTimeout(timeout time.Duration) {
	base.timeout = timeout
}

func (base *BaseClient) GetTimeout() time.Duration {
	return base.timeout
}

func NewBaseClient() *BaseClient {
	return &BaseClient{
		encoder:               &builtin.JsonEncoder{},
		isClientValidationSet: true,
		timeout:               DefaultTimeout,
	}
}
