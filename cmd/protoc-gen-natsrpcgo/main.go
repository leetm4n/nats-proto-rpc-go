package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

func main() {
	isValidatorEnabled := flag.Bool(
		"with_pb_validate",
		false,
		"generates validation code into handlers using protoc-gen-validate generated code",
	)

	protogen.Options{
		ParamFunc: flag.Set,
		ImportRewriteFunc: func(path protogen.GoImportPath) protogen.GoImportPath {
			return path
		},
	}.Run(func(gen *protogen.Plugin) error {
		for _, file := range gen.Files {
			// if file does not need to be generated or missing services then skip
			if !file.Generate || len(file.Services) == 0 {
				continue
			}

			// else generate file
			generateFile(gen, file, *isValidatorEnabled)
		}

		gen.SupportedFeatures = SupportedFeatures

		return nil
	})
}
