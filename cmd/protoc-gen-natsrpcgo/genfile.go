package main

import (
	natsrpcv1 "github.com/leetm4n/nats-proto-rpc-go/api/proto/nats/rpc/v1"
	"github.com/leetm4n/nats-proto-rpc-go/internal/generator"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func generateFile(gen *protogen.Plugin, protoFile *protogen.File, isValidatorEnabled bool) {
	// if file is not generation target then skip
	if !isFileGenerationTarget(protoFile) {
		return
	}

	generationTargetServices := getGenerationTargetServices(protoFile.Services)
	// if no services are generation target then skip
	if len(generationTargetServices) == 0 {
		return
	}

	filename := protoFile.GeneratedFilenamePrefix + ".pb.natsrpc.go"
	generatedFile := gen.NewGeneratedFile(filename, protoFile.GoImportPath)

	generator := generator.New(
		generationTargetServices,
		protoFile,
		generatedFile,
		isValidatorEnabled,
	)

	generator.Generate()
}

func getGenerationTargetServices(services []*protogen.Service) []*protogen.Service {
	targetServices := []*protogen.Service{}
	for _, service := range services {
		if isServiceGenerationTarget(service) {
			targetServices = append(targetServices, service)
		}
	}

	return targetServices
}

func isFileGenerationTarget(file *protogen.File) bool {
	return proto.GetExtension(file.Proto.Options, natsrpcv1.E_IsFileGenerationTarget.TypeDescriptor().Type()).(bool)
}

func isServiceGenerationTarget(service *protogen.Service) bool {
	return proto.GetExtension(service.Desc.Options(), natsrpcv1.E_IsServiceGenerationTarget.TypeDescriptor().Type()).(bool)
}
