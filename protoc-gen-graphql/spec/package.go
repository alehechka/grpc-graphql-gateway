package spec

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/types/pluginpb"
)

const (
	mainPackage = "main"
)

type PackageGetter interface {
	Package() string
	GoPackage() string
	Filename() string
}

type Versionable interface {
	CompilerVersion() *pluginpb.Version
	PluginVersion() string
}

type Package struct {
	Name                    string
	CamelName               string
	Path                    string
	GeneratedFilenamePrefix string
	FileName                string
	CompilerVersion         string
	PluginVersion           string
}

func NewPackage(g PackageGetter) *Package {
	p := &Package{}
	p.GeneratedFilenamePrefix = strings.TrimSuffix(g.Filename(), filepath.Ext(g.Filename()))
	p.FileName = filepath.Base(p.GeneratedFilenamePrefix)

	if c, ok := g.(Versionable); ok {
		p.CompilerVersion = compilerVersionString(c.CompilerVersion())
		p.PluginVersion = c.PluginVersion()
	}

	if pkg := g.GoPackage(); pkg != "" {
		p.Name, p.Path = ParsePackagePathName(pkg)
	} else if pkg := g.Package(); pkg != "" {
		p.Name = pkg
	} else {
		p.Name = strings.ReplaceAll(
			g.Filename(),
			filepath.Ext(g.Filename()),
			"",
		)
	}

	p.CamelName = strcase.ToCamel(p.Name)
	return p
}

func NewGooglePackage(m PackageGetter) *Package {
	name, _ := ParsePackagePathName(m.GoPackage())
	packageSuffix := strings.TrimSuffix(m.Filename(), filepath.Ext(m.Filename()))

	return &Package{
		Name:      "gql_ptypes_" + strings.ToLower(name),
		CamelName: strcase.ToCamel(name),
		Path:      "github.com/alehechka/grpc-graphql-gateway/ptypes/" + strings.ToLower(packageSuffix),
	}
}

func NewGoPackageFromString(pkg string) *Package {
	p := &Package{}
	p.Name, p.Path = ParsePackagePathName(pkg)
	return p
}

// ParsePackagePathName allows support custom package definitions like example.com/path/to/package:packageName
func ParsePackagePathName(pkg string) (name string, path string) {
	if index := strings.Index(pkg, ";"); index > -1 {
		name = pkg[index+1:]
		path = pkg[0:index]
	} else {
		name = filepath.Base(pkg)
		path = pkg
	}

	return
}

func compilerVersionString(cv *pluginpb.Version) string {
	if isBufCompiled(cv) {
		return "(unknown)"
	}

	ver := fmt.Sprintf("v%d.%d.%d", cv.GetMajor(), cv.GetMinor(), cv.GetPatch())
	if suffix := cv.GetSuffix(); len(suffix) > 0 {
		ver += fmt.Sprintf("-%s", suffix)
	}

	return ver
}
