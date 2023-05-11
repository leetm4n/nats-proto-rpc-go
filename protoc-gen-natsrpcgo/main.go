package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

func main() {
	isValidatorEnabled := flag.Bool("validator", false, "")
	protogen.Options{
		ParamFunc: flag.Set,
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
