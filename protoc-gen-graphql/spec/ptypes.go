package spec

import (
	"fmt"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

// Before protoc v3.14.0, go_package option name does not have "pb" suffix.
var supportedProtobufTypes = []string{
	"timestamp",
	"wrappers",
	"empty",
}

// After protoc v3.14.0, go_package option name have been changed.
// @see https://github.com/protocolbuffers/protobuf/releases/tag/v3.14.0
var supportedProtobufTypesLaterV3_14_0 = []string{
	"timestamppb",
	"wrapperspb",
	"emptypb",
}

var supportedGoogleTypes = []string{
	"money",
}

func getSupportedPtypeNames(cv *plugin.Version) []string {
	// TODO: buf.build does not currently supply a version and thus must be accounted for to use google.protobuf graphql types correctly
	if cv.GetMajor() == 0 && cv.GetMinor() == 0 && cv.GetPatch() == 0 && cv.GetSuffix() == "" {
		return supportedProtobufTypesLaterV3_14_0
	}

	if cv.GetMajor() >= 3 && cv.GetMinor() >= 14 {
		return supportedProtobufTypesLaterV3_14_0
	}
	return supportedProtobufTypes
}

func getImplementedPtypes(m *Message) (string, error) {
	ptype, _ := ParsePackagePathName(m.GoPackage())

	var found bool
	if pkg := m.Package(); pkg == "google.protobuf" {
		for _, v := range getSupportedPtypeNames(m.CompilerVersion) {
			if ptype == v {
				found = true
			}
		}
	} else if pkg == "google.type" {
		for _, v := range supportedGoogleTypes {
			if ptype == v {
				found = true
			}
		}
	}

	if !found {
		return "", fmt.Errorf("google's ptype \"%s\" does not implement for now.", ptype)
	}

	return ptype, nil
}
