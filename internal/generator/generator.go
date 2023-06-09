package generator

import (
	"fmt"
	"strings"
	"unicode"

	natsrpcv1 "github.com/leetm4n/nats-proto-rpc-go/api/proto/nats/rpc/v1"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

type Generator struct {
	services           []*protogen.Service
	protoFile          *protogen.File
	gen                *protogen.GeneratedFile
	isValidatorEnabled bool
}

func (g *Generator) Generate() {
	// this is the outline of generated code
	g.addHeaderComment()
	g.addEmptyLine()
	g.addPackageName()
	g.addEmptyLine()
	g.addImports(
		[]string{
			"context",
			"time",
			"encoding/json",
			"github.com/alecthomas/jsonschema",
			"github.com/nats-io/nats.go",
			"github.com/nats-io/nats.go/micro",
			"go.opentelemetry.io/otel/propagation",
			"go.opentelemetry.io/otel/trace",
			"go.opentelemetry.io/otel/codes",
			"go.opentelemetry.io/otel/attribute",
			"github.com/leetm4n/nats-proto-rpc-go/pkg/service",
			"github.com/leetm4n/nats-proto-rpc-go/pkg/client",
			"github.com/leetm4n/nats-proto-rpc-go/pkg/encoder",
			"github.com/leetm4n/nats-proto-rpc-go/pkg/subject",
			"github.com/leetm4n/nats-proto-rpc-go/pkg/telemetry",
		},
	)
	g.addEmptyLine()
	g.addServices()
}

func New(
	services []*protogen.Service,
	protoFile *protogen.File,
	gen *protogen.GeneratedFile,
	isValidatorEnabled bool,
) *Generator {
	return &Generator{
		services:           services,
		protoFile:          protoFile,
		gen:                gen,
		isValidatorEnabled: isValidatorEnabled,
	}
}

func (g *Generator) addServices() {
	for _, service := range g.services {
		g.addServiceHandlerInterface(service)
		g.addEmptyLine()
		g.addServiceClientInterface(service)
		g.addEmptyLine()
		g.addServiceClientImplementation(service)
		g.addEmptyLine()
		g.addServiceNatsMicroService(service)
	}
}

func (g *Generator) getServiceVersion(service *protogen.Service) string {
	serviceVersion, ok := proto.GetExtension(service.Desc.Options(), natsrpcv1.E_ServiceVersion.TypeDescriptor().Type()).(string)

	if !ok {
		return ""
	}

	return serviceVersion
}

func (g *Generator) getServiceSubject(service *protogen.Service) string {
	serviceSubject := service.GoName

	serviceSubjectOverride, ok := proto.GetExtension(service.Desc.Options(), natsrpcv1.E_ServiceSubject.TypeDescriptor().Type()).(string)

	if ok && serviceSubjectOverride != "" {
		serviceSubject = serviceSubjectOverride
	}

	return strings.ToLower(serviceSubject)
}

func (g *Generator) getMethodSubject(method *protogen.Method) string {
	methodSubject := method.GoName

	methodSubjectOverride, ok := proto.GetExtension(method.Desc.Options(), natsrpcv1.E_MethodSubject.TypeDescriptor().Type()).(string)

	if ok && methodSubjectOverride != "" {
		methodSubject = methodSubjectOverride
	}

	return strings.ToLower(methodSubject)
}

func (g *Generator) addImports(imports []string) {
	g.gen.P("import (")

	for _, importValue := range imports {
		g.gen.P(fmt.Sprintf("  \"%s\"", importValue))
	}

	g.gen.P(")")
}

func (g *Generator) addEmptyLine() {
	g.gen.P()
}

func (g *Generator) addHeaderComment() {
	g.gen.P("// Code generated by protoc-gen-natsrpc-go. DO NOT EDIT.")
	g.gen.P(fmt.Sprintf("// source: %s", *g.protoFile.Proto.Name))
}

func (g *Generator) addPackageName() {
	g.gen.P(fmt.Sprintf("package %s", g.protoFile.GoPackageName))
}

func (g *Generator) addServiceHandlerInterface(service *protogen.Service) {
	g.gen.P(fmt.Sprintf("// %sHandler should be implemented", service.GoName))
	g.gen.P(fmt.Sprintf("type %sHandler interface {", service.GoName))

	for _, method := range service.Methods {
		g.gen.P(fmt.Sprintf("  %s(ctx context.Context, req *%s) (res *%s, err error)", method.GoName, method.Input.GoIdent.GoName, method.Output.GoIdent.GoName))
	}

	g.gen.P("}")
	g.addEmptyLine()
}

func (g *Generator) addServiceClientInterface(service *protogen.Service) {
	g.gen.P(fmt.Sprintf("type %sClient interface {", service.GoName))

	for _, method := range service.Methods {
		g.gen.P(fmt.Sprintf("  %s(ctx context.Context, req *%s, subjectPrefix string) (res *%s, err error)", method.GoName, method.Input.GoIdent.GoName, method.Output.GoIdent.GoName))
	}

	g.gen.P("}")
}

func (g *Generator) addServiceClientImplementation(service *protogen.Service) {
	g.addServiceClientImplStruct(service)
	g.addServiceClientImplHandlers(service)
	g.addEmptyLine()
	g.addServiceClientImplementationConstructor(service)
}

func (g *Generator) addServiceClientImplHandlers(service *protogen.Service) {
	for _, method := range service.Methods {
		g.addEmptyLine()
		g.gen.P(fmt.Sprintf("func (c *%sNatsClient) %s(ctx context.Context, req *%s, subjectPrefix string) (*%s, error) {", service.GoName, method.GoName, method.Input.GoIdent.GoName, method.Output.GoIdent.GoName))
		g.gen.P("  header := telemetry.NatsHeaderCarrier{}")
		g.gen.P(fmt.Sprintf("  ctx, span := c.tracer.Start(ctx, \"%s\", trace.WithSpanKind(trace.SpanKindClient))", g.getMethodSubject(method)))
		g.gen.P("  defer span.End()")
		g.gen.P("  c.propagator.Inject(ctx, header)")
		g.gen.P(fmt.Sprintf("  subject := c.getSubject(\"%s\", \"%s\", subjectPrefix)", g.getServiceSubject(service), g.getMethodSubject(method)))
		g.gen.P("  if span.IsRecording() {")
		g.gen.P("    span.SetAttributes(")
		g.gen.P(fmt.Sprintf("      attribute.String(\"nats.serviceVersion\", \"%s\"),", g.getServiceVersion(service)))
		g.gen.P(fmt.Sprintf("      attribute.String(\"nats.service\", \"%s\"),", g.getServiceSubject(service)))
		g.gen.P(fmt.Sprintf("      attribute.String(\"nats.endpoint\", \"%s\"),", g.getMethodSubject(method)))
		g.gen.P("      attribute.String(\"nats.subject\", subject),")
		g.gen.P("    )")
		g.gen.P("  }")
		g.gen.P(fmt.Sprintf("  handler := func(ctx context.Context) (*%s, error) {", method.Output.GoIdent.GoName))

		if g.isValidatorEnabled {
			g.gen.P("    if (c.isValidationEnabled) {")
			g.gen.P("      if err := req.ValidateAll(); err != nil {")
			g.gen.P("        return nil, err")
			g.gen.P("      }")
			g.gen.P("    }")
		}

		g.gen.P("    encoded, err := c.encoder.Encode(subject, req)")
		g.gen.P("    if err != nil {")
		g.gen.P("      return nil, err")
		g.gen.P("    }")
		g.gen.P("    timeoutCtx, cancel := context.WithTimeout(ctx, c.timeout)")
		g.gen.P("    defer cancel()")
		g.gen.P("	   msg, err := c.natsConnection.RequestMsgWithContext(timeoutCtx, &nats.Msg{ Subject: subject, Data: encoded, Header: header.GetNatsHeader() })")
		g.gen.P("    if err != nil {")
		g.gen.P("      return nil, err")
		g.gen.P("    }")
		g.gen.P("    if err := client.DecodeErrorWithDecoder(msg.Header, msg.Data, c.errorDecoder); err != nil {")
		g.gen.P("      return nil, err")
		g.gen.P("    }")
		g.gen.P(fmt.Sprintf("    res := new(%s)", method.Output.GoIdent.GoName))
		g.gen.P("    if err := c.encoder.Decode(subject, msg.Data, res); err != nil {")
		g.gen.P("      return nil, err")
		g.gen.P("    }")

		if g.isValidatorEnabled {
			g.gen.P("    if c.isValidationEnabled {")
			g.gen.P("      if err := res.ValidateAll(); err != nil {")
			g.gen.P("        return nil, err")
			g.gen.P("      }")
			g.gen.P("    }")
		}

		g.gen.P("    return res, nil")
		g.gen.P("  }")
		g.gen.P("  result, err := handler(ctx)")
		g.gen.P("  if err != nil {")
		g.gen.P("    span.RecordError(err)")
		g.gen.P("    span.SetStatus(codes.Error, err.Error())")
		g.gen.P("  return nil, err")
		g.gen.P("  }")
		g.gen.P("  return result, nil")
		g.gen.P("}")
	}
}

func (g *Generator) addServiceClientImplStruct(service *protogen.Service) {
	g.gen.P("// type check.")
	g.gen.P(fmt.Sprintf("var _ %sClient = (*%sNatsClient)(nil)", service.GoName, service.GoName))
	g.gen.P(fmt.Sprintf("type %sNatsClient struct {", service.GoName))
	g.gen.P("  natsConnection *nats.Conn")
	g.gen.P("  encoder encoder.Encoder")
	g.gen.P("  isValidationEnabled bool")
	g.gen.P("  getSubject subject.GetSubjectFn")
	g.gen.P("  timeout time.Duration")
	g.gen.P("  errorDecoder client.ErrorDecoderFn")
	g.gen.P("  tracer trace.Tracer")
	g.gen.P("  propagator propagation.TextMapPropagator")
	g.gen.P("}")
}

func (g *Generator) addServiceClientImplementationConstructor(service *protogen.Service) {
	g.gen.P(fmt.Sprintf("func New%sNatsClient(options client.Options) *%sNatsClient {", service.GoName, service.GoName))
	g.gen.P(fmt.Sprintf("  return &%s{", fmt.Sprintf("%sNatsClient", service.GoName)))
	g.gen.P("    natsConnection: options.NatsConnection,")
	g.gen.P("    encoder: options.Encoder,")
	g.gen.P("    isValidationEnabled: options.IsValidationEnabled,")
	g.gen.P("    getSubject: options.GetSubject,")
	g.gen.P("    timeout: options.Timeout,")
	g.gen.P("    errorDecoder: options.ErrorDecoder,")
	g.gen.P("    tracer: options.Telemetry.Tracer,")
	g.gen.P("    propagator: options.Telemetry.Propagator,")
	g.gen.P("  }")
	g.gen.P("}")
}

func (g *Generator) addServiceNatsMicroService(service *protogen.Service) {
	g.addServiceHandlerNatsMicroServiceImplStruct(service)
	g.addServiceHandlerNatsMicroServiceImpl(service)
	g.addEmptyLine()
	g.addServiceHandlerNatsMicroServiceConstructor(service)
}

func (g *Generator) addServiceHandlerNatsMicroServiceImplStruct(service *protogen.Service) {
	g.gen.P(fmt.Sprintf("type %sNatsMicroServiceWrapper struct {", service.GoName))
	g.gen.P(fmt.Sprintf("  %sHandler  %sHandler", firstLetterToLower(service.GoName), service.GoName))
	g.gen.P("  service micro.Service")
	g.gen.P("  natsConnection      *nats.Conn")
	g.gen.P("  encoder             encoder.Encoder")
	g.gen.P("  isValidationEnabled bool")
	g.gen.P("  errorEncoder        service.ErrorEncoderFn")
	g.gen.P("  subscriptionErrorHandler        micro.ErrHandler")
	g.gen.P("  doneHandler         micro.DoneHandler")
	g.gen.P("  startHandler        micro.DoneHandler")
	g.gen.P("  endpointAddHandler  service.EndpointAddHandlerFn")
	g.gen.P("  subjectPrefix       string")
	g.gen.P("  getSubject          subject.GetSubjectFn")
	g.gen.P("  tracer              trace.Tracer")
	g.gen.P("  propagator          propagation.TextMapPropagator")
	g.gen.P("  responseErrorHandler  service.ResponseErrorHandlerFn")
	g.gen.P("}")
}

func (g *Generator) addServiceHandlerNatsMicroServiceImpl(service *protogen.Service) {
	g.gen.P(fmt.Sprintf("func (s *%sNatsMicroServiceWrapper) Start(ctx context.Context) error {", service.GoName))
	g.gen.P("  service, err := micro.AddService(s.natsConnection, micro.Config{")
	g.gen.P(fmt.Sprintf("    Name:    \"%s\",", g.getServiceSubject(service)))
	g.gen.P(fmt.Sprintf("	   Version: \"%s\",", g.getServiceVersion(service)))
	g.gen.P("	   ErrorHandler: s.subscriptionErrorHandler,")
	g.gen.P("	   DoneHandler: s.doneHandler,")
	g.gen.P("  })")
	g.gen.P("  if err != nil {")
	g.gen.P("    return err")
	g.gen.P("  }")
	g.gen.P("  s.service = service")
	g.gen.P("  if s.startHandler != nil {")
	g.gen.P("    s.startHandler(service)")
	g.gen.P("  }")

	for _, method := range service.Methods {
		g.gen.P(fmt.Sprintf("  %sSubject := s.getSubject(\"%s\", \"%s\", s.subjectPrefix)", firstLetterToLower(method.GoName), g.getServiceSubject(service), g.getMethodSubject(method)))
		g.gen.P(fmt.Sprintf("  %sSchema, err := json.Marshal(*jsonschema.Reflect(%s{}))", firstLetterToLower(method.Input.GoIdent.GoName), method.Input.GoIdent.GoName))
		g.gen.P("  if err != nil {")
		g.gen.P("    return err")
		g.gen.P("  }")
		g.gen.P(fmt.Sprintf("  %sSchema, err := json.Marshal(*jsonschema.Reflect(%s{}))", firstLetterToLower(method.Output.GoIdent.GoName), method.Output.GoIdent.GoName))
		g.gen.P("  if err != nil {")
		g.gen.P("    return err")
		g.gen.P("  }")
		g.gen.P("  if err := service.AddEndpoint(")
		g.gen.P(fmt.Sprintf("\"%s\",", g.getMethodSubject(method)))
		g.gen.P("    micro.ContextHandler(")
		g.gen.P("      ctx,")
		g.gen.P("      func(ctx context.Context, request micro.Request) {")
		g.gen.P("        ctxWithTrace := s.propagator.Extract(ctx, telemetry.NatsHeaderCarrierFromNatsHeader(request.Headers()))")
		g.gen.P(fmt.Sprintf("        ctxWithSpan, span := s.tracer.Start(ctxWithTrace, \"%s\", trace.WithSpanKind(trace.SpanKindServer))", g.getMethodSubject(method)))
		g.gen.P("        defer span.End()")
		g.gen.P("        if span.IsRecording() {")
		g.gen.P("          span.SetAttributes(")
		g.gen.P(fmt.Sprintf("            attribute.String(\"nats.serviceVersion\", \"%s\"),", g.getServiceVersion(service)))
		g.gen.P(fmt.Sprintf("            attribute.String(\"nats.service\", \"%s\"),", g.getServiceSubject(service)))
		g.gen.P(fmt.Sprintf("            attribute.String(\"nats.endpoint\", \"%s\"),", g.getMethodSubject(method)))
		g.gen.P(fmt.Sprintf("            attribute.String(\"nats.subject\", %sSubject),", firstLetterToLower(method.GoName)))
		g.gen.P("          )")
		g.gen.P("        }")
		g.gen.P("        handler := func (ctx context.Context, request micro.Request) ([]byte, error) {")
		g.gen.P("          select {")
		g.gen.P("          case <-ctx.Done():")
		g.gen.P("             return nil, ctx.Err()")
		g.gen.P("          default:")
		g.gen.P(fmt.Sprintf("            req := new(%s)", method.Input.GoIdent.GoName))
		g.gen.P(fmt.Sprintf("            if err := s.encoder.Decode(%sSubject, request.Data(), req); err != nil {", firstLetterToLower(method.GoName)))
		g.gen.P("              return nil, err")
		g.gen.P("            }")

		if g.isValidatorEnabled {
			g.gen.P("            if s.isValidationEnabled {")
			g.gen.P("              if err := req.ValidateAll(); err != nil {")
			g.gen.P("                return nil, err")
			g.gen.P("              }")
			g.gen.P("            }")
		}

		g.gen.P(fmt.Sprintf("            res, err := s.%sHandler.%s(ctx, req)", firstLetterToLower(service.GoName), method.GoName))
		g.gen.P("            if err != nil {")
		g.gen.P("              return nil, err")
		g.gen.P("            }")

		if g.isValidatorEnabled {
			g.gen.P("            if s.isValidationEnabled {")
			g.gen.P("              if err := res.ValidateAll(); err != nil {")
			g.gen.P("                return nil, err")
			g.gen.P("              }")
			g.gen.P("            }")
		}

		g.gen.P(fmt.Sprintf("            payload, err := s.encoder.Encode(%sSubject, res)", firstLetterToLower(method.GoName)))
		g.gen.P("            if err != nil {")
		g.gen.P("						   return nil, err")
		g.gen.P("					   }")
		g.gen.P("            return payload, nil")
		g.gen.P("          }")
		g.gen.P("        }")
		g.gen.P("        payload, err := handler(ctxWithSpan, request)")
		g.gen.P("        if err != nil {")
		g.gen.P("          span.RecordError(err)")
		g.gen.P("          span.SetStatus(codes.Error, err.Error())")
		g.gen.P("          code, description, payload := s.errorEncoder(err)")
		g.gen.P("          if err := request.Error(code, description, payload); err != nil {")
		g.gen.P("            s.responseErrorHandler(err)")
		g.gen.P("            span.RecordError(err)")
		g.gen.P("            span.SetStatus(codes.Error, err.Error())")
		g.gen.P("          }")
		g.gen.P("          return")
		g.gen.P("        }")
		g.gen.P("        if err := request.Respond(payload); err != nil {")
		g.gen.P("          s.responseErrorHandler(err)")
		g.gen.P("          span.RecordError(err)")
		g.gen.P("          span.SetStatus(codes.Error, err.Error())")
		g.gen.P("        }")
		g.gen.P("      },")
		g.gen.P("    ),")
		g.gen.P("    micro.WithEndpointSchema(&micro.Schema{")
		g.gen.P(fmt.Sprintf("      Request: string(%sSchema),", firstLetterToLower(method.Input.GoIdent.GoName)))
		g.gen.P(fmt.Sprintf("      Response: string(%sSchema),", firstLetterToLower(method.Output.GoIdent.GoName)))
		g.gen.P("    }),")
		g.gen.P(fmt.Sprintf("    micro.WithEndpointSubject(%sSubject),", firstLetterToLower(method.GoName)))
		g.gen.P("  ); err != nil {")
		g.gen.P("    return err")
		g.gen.P("  }")
		g.gen.P("  if s.endpointAddHandler != nil {")
		g.gen.P(fmt.Sprintf("    s.endpointAddHandler(service, \"%s\", %sSubject)", g.getMethodSubject(method), firstLetterToLower(method.GoName)))
		g.gen.P("  }")
	}

	g.gen.P("  return nil")
	g.gen.P("}")
	g.addEmptyLine()
	g.gen.P(fmt.Sprintf("func (s *%sNatsMicroServiceWrapper) GetNatsMicroService() micro.Service {", service.GoName))
	g.gen.P("  return s.service")
	g.gen.P("}")
}

func (g *Generator) addServiceHandlerNatsMicroServiceConstructor(service *protogen.Service) {
	g.gen.P(fmt.Sprintf("func New%s(", fmt.Sprintf("%sNatsMicroServiceWrapper", service.GoName)))
	g.gen.P(fmt.Sprintf("  %sHandler %sHandler,", firstLetterToLower(service.GoName), service.GoName))
	g.gen.P("  options service.Options,")
	g.gen.P(fmt.Sprintf(") *%sNatsMicroServiceWrapper {", service.GoName))
	g.gen.P(fmt.Sprintf("  return &%sNatsMicroServiceWrapper{", service.GoName))
	g.gen.P(fmt.Sprintf("    %sHandler: %sHandler,", firstLetterToLower(service.GoName), firstLetterToLower(service.GoName)))
	g.gen.P("    natsConnection: options.NatsConnection,")
	g.gen.P("    encoder: options.Encoder,")
	g.gen.P("    isValidationEnabled: options.IsValidationEnabled,")
	g.gen.P("    errorEncoder: options.ErrorEncoder,")
	g.gen.P("    subscriptionErrorHandler: options.Hooks.SubscriptionErrorHandler,")
	g.gen.P("    doneHandler: options.Hooks.DoneHandler,")
	g.gen.P("    startHandler: options.Hooks.StartHandler,")
	g.gen.P("    subjectPrefix: options.SubjectPrefix,")
	g.gen.P("    endpointAddHandler: options.Hooks.EndpointAddHandler,")
	g.gen.P("    responseErrorHandler: options.Hooks.ResponseErrorHandler,")
	g.gen.P("    getSubject: options.GetSubject,")
	g.gen.P("    propagator: options.Telemetry.Propagator,")
	g.gen.P("    tracer: options.Telemetry.Tracer,")
	g.gen.P("  }")
	g.gen.P("}")
}

func firstLetterToLower(s string) string {
	if len(s) == 0 {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToLower(r[0])

	return string(r)
}
