package spec

import (
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// EnumValue spec wraps EnumValueDescriptorProto with keeping file definition.
type EnumValue struct {
	descriptor *descriptor.EnumValueDescriptorProto
	*File

	paths       []int
	useJsonName bool
}

func NewEnumValue(
	d *descriptor.EnumValueDescriptorProto,
	f *File,
	useJsonName bool,
	paths ...int,
) *EnumValue {

	return &EnumValue{
		descriptor:  d,
		File:        f,
		useJsonName: useJsonName,
		paths:       paths,
	}
}

func (e *EnumValue) Comment() string {
	return e.File.getComment(e.paths)
}

func (e *EnumValue) Number() int32 {
	return e.descriptor.GetNumber()
}

func (e *EnumValue) Name() string {
	return e.descriptor.GetName()
}

func (e *EnumValue) UseJsonName() bool {
	if e != nil {
		return e.useJsonName
	}
	return false
}
