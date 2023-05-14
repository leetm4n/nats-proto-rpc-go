// Code generated by protoc-gen-nrpc. DO NOT EDIT.
// source: example/service/v1/service.proto

package exampleservicev1

import (
	"context"
	"time"
	"encoding/json"
	"github.com/alecthomas/jsonschema"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
	"github.com/leetm4n/nats-proto-rpc-go/pkg/correlation"
	"github.com/leetm4n/nats-proto-rpc-go/pkg/runnable"
	"github.com/leetm4n/nats-proto-rpc-go/pkg/client"
	"github.com/leetm4n/nats-proto-rpc-go/pkg/encoder"
	"github.com/leetm4n/nats-proto-rpc-go/pkg/subject"
)

type TestServiceServer interface {
	SendMessage(ctx context.Context, req *SendMessageRequest) (res *SendMessageResponse, err error)
	GetMessage(ctx context.Context, req *GetMessageRequest) (res *GetMessageResponse, err error)
}

type TestServiceClient interface {
	SendMessage(ctx context.Context, req *SendMessageRequest, subjectPrefix string) (res *SendMessageResponse, err error)
	GetMessage(ctx context.Context, req *GetMessageRequest, subjectPrefix string) (res *GetMessageResponse, err error)
}

type testServiceClient struct {
	natsConnection      *nats.Conn
	encoder             encoder.Encoder
	isValidationEnabled bool
	getSubject          subject.GetSubjectFn
	timeout             time.Duration
}

func (c *testServiceClient) SendMessage(ctx context.Context, req *SendMessageRequest, subjectPrefix string) (*SendMessageResponse, error) {
	subject := c.getSubject("customname", "sendmessage", subjectPrefix)
	if c.isValidationEnabled {
		if err := req.ValidateAll(); err != nil {
			return nil, err
		}
	}
	encoded, err := c.encoder.Encode(subject, req)
	if err != nil {
		return nil, err
	}
	correlationID := correlation.CorrelationIDFromContext(ctx)
	var header nats.Header
	if correlationID != "" {
		header.Add(correlation.RequestIDHeaderKey, correlationID)
	}
	msg, err := c.natsConnection.RequestMsg(&nats.Msg{Subject: subject, Data: encoded, Header: header}, c.timeout)
	if err != nil {
		return nil, err
	}
	res := new(SendMessageResponse)
	if err := c.encoder.Decode(subject, msg.Data, res); err != nil {
		return nil, err
	}
	if c.isValidationEnabled {
		if err := res.ValidateAll(); err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (c *testServiceClient) GetMessage(ctx context.Context, req *GetMessageRequest, subjectPrefix string) (*GetMessageResponse, error) {
	subject := c.getSubject("customname", "fetchmessage", subjectPrefix)
	if c.isValidationEnabled {
		if err := req.ValidateAll(); err != nil {
			return nil, err
		}
	}
	encoded, err := c.encoder.Encode(subject, req)
	if err != nil {
		return nil, err
	}
	correlationID := correlation.CorrelationIDFromContext(ctx)
	var header nats.Header
	if correlationID != "" {
		header.Add(correlation.RequestIDHeaderKey, correlationID)
	}
	msg, err := c.natsConnection.RequestMsg(&nats.Msg{Subject: subject, Data: encoded, Header: header}, c.timeout)
	if err != nil {
		return nil, err
	}
	res := new(GetMessageResponse)
	if err := c.encoder.Decode(subject, msg.Data, res); err != nil {
		return nil, err
	}
	if c.isValidationEnabled {
		if err := res.ValidateAll(); err != nil {
			return nil, err
		}
	}
	return res, nil
}

func NewTestServiceClient(options client.Options) TestServiceClient {
	return &testServiceClient{
		natsConnection:      options.NatsConnection,
		encoder:             options.Encoder,
		isValidationEnabled: options.IsValidationEnabled,
		getSubject:          options.GetSubject,
		timeout:             options.Timeout,
	}
}

type testServiceServerRunnable struct {
	testServiceServer   TestServiceServer
	service             *micro.Service
	natsConnection      *nats.Conn
	encoder             encoder.Encoder
	isValidationEnabled bool
	errorMapper         runnable.ErrorMapper
	errorHandler        micro.ErrHandler
	doneHandler         micro.DoneHandler
	subjectPrefix       string
	getSubject          subject.GetSubjectFn
}

func (s *testServiceServerRunnable) Run(ctx context.Context) error {
	service, err := micro.AddService(s.natsConnection, micro.Config{
		Name:         "customname",
		Version:      "v1",
		ErrorHandler: s.errorHandler,
		DoneHandler:  s.doneHandler,
	})
	if err != nil {
		return err
	}
	sendMessageSubject := s.getSubject("customname", "sendmessage", s.subjectPrefix)
	sendMessageRequestSchema, err := json.Marshal(*jsonschema.Reflect(SendMessageRequest{}))
	if err != nil {
		return err
	}
	sendMessageResponseSchema, err := json.Marshal(*jsonschema.Reflect(SendMessageResponse{}))
	if err != nil {
		return err
	}
	if err := service.AddEndpoint(
		sendMessageSubject,
		micro.ContextHandler(
			ctx,
			func(ctx context.Context, request micro.Request) {
				select {
				case <-ctx.Done():
					code, description := s.errorMapper(ctx.Err())
					request.Error(code, description, nil)
					return
				default:
					req := new(SendMessageRequest)
					if err := s.encoder.Decode(sendMessageSubject, request.Data(), req); err != nil {
						code, description := s.errorMapper(err)
						request.Error(code, description, nil)
						return
					}
					if s.isValidationEnabled {
						if err := req.ValidateAll(); err != nil {
							code, description := s.errorMapper(err)
							request.Error(code, description, nil)
							return
						}
					}
					res, err := s.testServiceServer.SendMessage(ctx, req)
					if err != nil {
						code, description := s.errorMapper(err)
						request.Error(code, description, nil)
						return
					}
					if s.isValidationEnabled {
						if err := res.ValidateAll(); err != nil {
							code, description := s.errorMapper(err)
							request.Error(code, description, nil)
							return
						}
					}
					payload, err := s.encoder.Encode(sendMessageSubject, res)
					if err != nil {
						code, description := s.errorMapper(err)
						request.Error(code, description, nil)
						return
					}
					request.Respond(payload)
				}
			},
		),
		micro.WithEndpointSchema(&micro.Schema{
			Request:  string(sendMessageRequestSchema),
			Response: string(sendMessageResponseSchema),
		}),
	); err != nil {
		return err
	}
	getMessageSubject := s.getSubject("customname", "fetchmessage", s.subjectPrefix)
	getMessageRequestSchema, err := json.Marshal(*jsonschema.Reflect(GetMessageRequest{}))
	if err != nil {
		return err
	}
	getMessageResponseSchema, err := json.Marshal(*jsonschema.Reflect(GetMessageResponse{}))
	if err != nil {
		return err
	}
	if err := service.AddEndpoint(
		getMessageSubject,
		micro.ContextHandler(
			ctx,
			func(ctx context.Context, request micro.Request) {
				select {
				case <-ctx.Done():
					code, description := s.errorMapper(ctx.Err())
					request.Error(code, description, nil)
					return
				default:
					req := new(GetMessageRequest)
					if err := s.encoder.Decode(getMessageSubject, request.Data(), req); err != nil {
						code, description := s.errorMapper(err)
						request.Error(code, description, nil)
						return
					}
					if s.isValidationEnabled {
						if err := req.ValidateAll(); err != nil {
							code, description := s.errorMapper(err)
							request.Error(code, description, nil)
							return
						}
					}
					res, err := s.testServiceServer.GetMessage(ctx, req)
					if err != nil {
						code, description := s.errorMapper(err)
						request.Error(code, description, nil)
						return
					}
					if s.isValidationEnabled {
						if err := res.ValidateAll(); err != nil {
							code, description := s.errorMapper(err)
							request.Error(code, description, nil)
							return
						}
					}
					payload, err := s.encoder.Encode(getMessageSubject, res)
					if err != nil {
						code, description := s.errorMapper(err)
						request.Error(code, description, nil)
						return
					}
					request.Respond(payload)
				}
			},
		),
		micro.WithEndpointSchema(&micro.Schema{
			Request:  string(getMessageRequestSchema),
			Response: string(getMessageResponseSchema),
		}),
	); err != nil {
		return err
	}
	return nil
}

func (s *testServiceServerRunnable) GetNatsMicroService() *micro.Service {
	return s.service
}

func NewTestServiceRunnable(
	testServiceServer TestServiceServer,
	options runnable.Options,
) runnable.Runnable {
	return &testServiceServerRunnable{
		testServiceServer:   testServiceServer,
		natsConnection:      options.NatsConnection,
		encoder:             options.Encoder,
		isValidationEnabled: options.IsValidationEnabled,
		errorMapper:         options.ErrorMapper,
		errorHandler:        options.ErrorHandler,
		doneHandler:         options.DoneHandler,
		subjectPrefix:       options.SubjectPrefix,
		getSubject:          options.GetSubject,
	}
}
