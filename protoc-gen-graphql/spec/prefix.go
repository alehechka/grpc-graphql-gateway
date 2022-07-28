package spec

import (
	"strings"
)

// PrefixType adds prefix to avoid conflicting name
func PrefixType(name string) string {
	return "Gql__type_" + name + "()"
}

// PrefixEnum adds prefix to avoid conflicting name
func PrefixEnum(name string) string {
	return "Gql__enum_" + name + "()"
}

// PrefixInput adds prefix to avoid conflicting name
func PrefixInput(name string) string {
	return "Gql__input_" + name + "()"
}

// PrefixInterface adds prefix to avoid conflicting name
func PrefixInterface(name string) string {
	return "Gql__interface_" + name + "()"
}

func IsGooglePackage(p PackageGetter) bool {
	pkg := p.Package()

	if strings.HasPrefix(pkg, "google.protobuf") {
		return true
	}

	return strings.HasPrefix(pkg, "google.type")
}
