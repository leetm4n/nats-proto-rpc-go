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
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/attribute"
	"github.com/leetm4n/nats-proto-rpc-go/pkg/runnable"
	"github.com/leetm4n/nats-proto-rpc-go/pkg/client"
	"github.com/leetm4n/nats-proto-rpc-go/pkg/encoder"
	"github.com/leetm4n/nats-proto-rpc-go/pkg/subject"
	"github.com/leetm4n/nats-proto-rpc-go/pkg/telemetry"
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
	errorDecoder        client.ErrorDecoderFn
	tracer              trace.Tracer
	propagator          propagation.TextMapPropagator
}

func (c *testServiceClient) SendMessage(ctx context.Context, req *SendMessageRequest, subjectPrefix string) (*SendMessageResponse, error) {
	header := telemetry.NatsHeaderCarrier{}
	ctx, span := c.tracer.Start(ctx, "sendmessage", trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()
	c.propagator.Inject(ctx, header)
	subject := c.getSubject("customname", "sendmessage", subjectPrefix)
	if span.IsRecording() {
		span.SetAttributes(
			attribute.String("nats.serviceVersion", "1.0.0"),
			attribute.String("nats.service", "customname"),
			attribute.String("nats.endpoint", "sendmessage"),
			attribute.String("nats.subject", subject),
		)
	}
	handler := func(ctx context.Context) (*SendMessageResponse, error) {
		if c.isValidationEnabled {
			if err := req.ValidateAll(); err != nil {
				return nil, err
			}
		}
		encoded, err := c.encoder.Encode(subject, req)
		if err != nil {
			return nil, err
		}
		timeoutCtx, cancel := context.WithTimeout(ctx, c.timeout)
		defer cancel()
		msg, err := c.natsConnection.RequestMsgWithContext(timeoutCtx, &nats.Msg{Subject: subject, Data: encoded, Header: header.GetNatsHeader()})
		if err != nil {
			return nil, err
		}
		if err := client.DecodeErrorFromHeaderWithDecoder(msg.Header, c.errorDecoder); err != nil {
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
	result, err := handler(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	return result, nil
}

func (c *testServiceClient) GetMessage(ctx context.Context, req *GetMessageRequest, subjectPrefix string) (*GetMessageResponse, error) {
	header := telemetry.NatsHeaderCarrier{}
	ctx, span := c.tracer.Start(ctx, "fetchmessage", trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()
	c.propagator.Inject(ctx, header)
	subject := c.getSubject("customname", "fetchmessage", subjectPrefix)
	if span.IsRecording() {
		span.SetAttributes(
			attribute.String("nats.serviceVersion", "1.0.0"),
			attribute.String("nats.service", "customname"),
			attribute.String("nats.endpoint", "fetchmessage"),
			attribute.String("nats.subject", subject),
		)
	}
	handler := func(ctx context.Context) (*GetMessageResponse, error) {
		if c.isValidationEnabled {
			if err := req.ValidateAll(); err != nil {
				return nil, err
			}
		}
		encoded, err := c.encoder.Encode(subject, req)
		if err != nil {
			return nil, err
		}
		timeoutCtx, cancel := context.WithTimeout(ctx, c.timeout)
		defer cancel()
		msg, err := c.natsConnection.RequestMsgWithContext(timeoutCtx, &nats.Msg{Subject: subject, Data: encoded, Header: header.GetNatsHeader()})
		if err != nil {
			return nil, err
		}
		if err := client.DecodeErrorFromHeaderWithDecoder(msg.Header, c.errorDecoder); err != nil {
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
	result, err := handler(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	return result, nil
}

func NewTestServiceClient(options client.Options) TestServiceClient {
	return &testServiceClient{
		natsConnection:      options.NatsConnection,
		encoder:             options.Encoder,
		isValidationEnabled: options.IsValidationEnabled,
		getSubject:          options.GetSubject,
		timeout:             options.Timeout,
		errorDecoder:        options.ErrorDecoder,
		tracer:              options.Telemetry.Tracer,
		propagator:          options.Telemetry.Propagator,
	}
}

type testServiceServerRunnable struct {
	testServiceServer   TestServiceServer
	service             micro.Service
	natsConnection      *nats.Conn
	encoder             encoder.Encoder
	isValidationEnabled bool
	errorEncoder        runnable.ErrorEncoderFn
	errorHandler        micro.ErrHandler
	doneHandler         micro.DoneHandler
	startHandler        micro.DoneHandler
	endpointAddHandler  runnable.EndpointAddHandlerFn
	subjectPrefix       string
	getSubject          subject.GetSubjectFn
	tracer              trace.Tracer
	propagator          propagation.TextMapPropagator
}

func (s *testServiceServerRunnable) Run(ctx context.Context) error {
	service, err := micro.AddService(s.natsConnection, micro.Config{
		Name:         "customname",
		Version:      "1.0.0",
		ErrorHandler: s.errorHandler,
		DoneHandler:  s.doneHandler,
	})
	if err != nil {
		return err
	}
	s.service = service
	if s.startHandler != nil {
		s.startHandler(service)
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
		"sendmessage",
		micro.ContextHandler(
			ctx,
			func(ctx context.Context, request micro.Request) {
				ctx, span := s.tracer.Start(ctx, "sendmessage", trace.WithSpanKind(trace.SpanKindServer))
				defer span.End()
				ctxWithSpan := s.propagator.Extract(ctx, telemetry.NatsHeaderCarrierFromNatsHeader(request.Headers()))
				if span.IsRecording() {
					span.SetAttributes(
						attribute.String("nats.serviceVersion", "1.0.0"),
						attribute.String("nats.service", "customname"),
						attribute.String("nats.endpoint", "sendmessage"),
						attribute.String("nats.subject", sendMessageSubject),
					)
				}
				handler := func(ctx context.Context, request micro.Request) ([]byte, error) {
					select {
					case <-ctx.Done():
						return nil, ctx.Err()
					default:
						req := new(SendMessageRequest)
						if err := s.encoder.Decode(sendMessageSubject, request.Data(), req); err != nil {
							return nil, err
						}
						if s.isValidationEnabled {
							if err := req.ValidateAll(); err != nil {
								return nil, err
							}
						}
						res, err := s.testServiceServer.SendMessage(ctx, req)
						if err != nil {
							return nil, err
						}
						if s.isValidationEnabled {
							if err := res.ValidateAll(); err != nil {
								return nil, err
							}
						}
						payload, err := s.encoder.Encode(sendMessageSubject, res)
						if err != nil {
							return nil, err
						}
						return payload, nil
					}
				}
				payload, err := handler(ctxWithSpan, request)
				if err != nil {
					span.RecordError(err)
					span.SetStatus(codes.Error, err.Error())
					code, description := s.errorEncoder(err)
					request.Error(code, description, nil)
				}
				request.Respond(payload)
			},
		),
		micro.WithEndpointSchema(&micro.Schema{
			Request:  string(sendMessageRequestSchema),
			Response: string(sendMessageResponseSchema),
		}),
		micro.WithEndpointSubject(sendMessageSubject),
	); err != nil {
		return err
	}
	if s.endpointAddHandler != nil {
		s.endpointAddHandler(service, "sendmessage", sendMessageSubject)
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
		"fetchmessage",
		micro.ContextHandler(
			ctx,
			func(ctx context.Context, request micro.Request) {
				ctx, span := s.tracer.Start(ctx, "fetchmessage", trace.WithSpanKind(trace.SpanKindServer))
				defer span.End()
				ctxWithSpan := s.propagator.Extract(ctx, telemetry.NatsHeaderCarrierFromNatsHeader(request.Headers()))
				if span.IsRecording() {
					span.SetAttributes(
						attribute.String("nats.serviceVersion", "1.0.0"),
						attribute.String("nats.service", "customname"),
						attribute.String("nats.endpoint", "fetchmessage"),
						attribute.String("nats.subject", getMessageSubject),
					)
				}
				handler := func(ctx context.Context, request micro.Request) ([]byte, error) {
					select {
					case <-ctx.Done():
						return nil, ctx.Err()
					default:
						req := new(GetMessageRequest)
						if err := s.encoder.Decode(getMessageSubject, request.Data(), req); err != nil {
							return nil, err
						}
						if s.isValidationEnabled {
							if err := req.ValidateAll(); err != nil {
								return nil, err
							}
						}
						res, err := s.testServiceServer.GetMessage(ctx, req)
						if err != nil {
							return nil, err
						}
						if s.isValidationEnabled {
							if err := res.ValidateAll(); err != nil {
								return nil, err
							}
						}
						payload, err := s.encoder.Encode(getMessageSubject, res)
						if err != nil {
							return nil, err
						}
						return payload, nil
					}
				}
				payload, err := handler(ctxWithSpan, request)
				if err != nil {
					span.RecordError(err)
					span.SetStatus(codes.Error, err.Error())
					code, description := s.errorEncoder(err)
					request.Error(code, description, nil)
				}
				request.Respond(payload)
			},
		),
		micro.WithEndpointSchema(&micro.Schema{
			Request:  string(getMessageRequestSchema),
			Response: string(getMessageResponseSchema),
		}),
		micro.WithEndpointSubject(getMessageSubject),
	); err != nil {
		return err
	}
	if s.endpointAddHandler != nil {
		s.endpointAddHandler(service, "fetchmessage", getMessageSubject)
	}
	return nil
}

func (s *testServiceServerRunnable) GetNatsMicroService() micro.Service {
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
		errorEncoder:        options.ErrorEncoder,
		errorHandler:        options.Hooks.ErrorHandler,
		doneHandler:         options.Hooks.DoneHandler,
		startHandler:        options.Hooks.StartHandler,
		subjectPrefix:       options.SubjectPrefix,
		endpointAddHandler:  options.Hooks.EndpointAddHandler,
		getSubject:          options.GetSubject,
		propagator:          options.Telemetry.Propagator,
		tracer:              options.Telemetry.Tracer,
	}
}
