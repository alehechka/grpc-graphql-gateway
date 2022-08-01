package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"

	// nolint: staticcheck
	"github.com/alehechka/grpc-graphql-gateway/protoc-gen-graphql/generator"
	"github.com/alehechka/grpc-graphql-gateway/protoc-gen-graphql/spec"
	"github.com/alehechka/grpc-graphql-gateway/protoc-gen-graphql/template"
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

var version = "dev"
var printVersion = flag.Bool("v", false, "show binary version")

func main() {
	flag.Parse()
	if *printVersion {
		io.WriteString(os.Stdout, version) // nolint: errcheck
		os.Exit(0)
	}

	var genError error

	resp := &plugin.CodeGeneratorResponse{}
	defer func() {
		// If some error has been occurred in generate process,
		// add error message to plugin response
		if genError != nil {
			message := genError.Error()
			resp.Error = &message
		}
		buf, err := proto.Marshal(resp)
		if err != nil {
			log.Fatalln(err)
		}
		os.Stdout.Write(buf)
	}()

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(os.Stdin); err != nil {
		genError = err
		return
	}

	var req plugin.CodeGeneratorRequest
	if err := proto.Unmarshal(buf.Bytes(), &req); err != nil {
		genError = err
		return
	}

	var parameter string
	if req.Parameter != nil {
		parameter = req.GetParameter()
	}
	args, err := spec.NewParams(parameter)
	if err != nil {
		genError = err
		return
	}

	if args.FieldProtoName {
		log.Println("[INFO] field_proto option is provided. All type fields will use the `protobuf` name tag.")
	}

	// We're dealing with each descriptors to out wrapper struct
	// in order to access easily plugin options, package name, comment, etc...
	var files []*spec.File
	for _, f := range req.GetProtoFile() {
		files = append(files, spec.NewFile(f, &spec.FileConfig{
			UseProtoName:    args.FieldProtoName,
			CompilerVersion: req.GetCompilerVersion(),
			PluginVersion:   version,
		}))
	}

	g := generator.New(files, args)
	var ftg []string
	for _, f := range req.GetFileToGenerate() {
		if !args.IsExclude(f) {
			ftg = append(ftg, f)
		}
	}
	if len(ftg) > 0 {
		genFiles, err := g.Generate(template.GoTemplate, ftg)
		if err != nil {
			genError = err
			return
		}
		resp.File = append(resp.File, genFiles...)
	}
}
