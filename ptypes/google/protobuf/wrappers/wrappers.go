package wrappers

import (
	"github.com/golang/protobuf/ptypes/wrappers"
)

// Expose Google defined ptypes as this package types
type (
	DoubleValue = wrappers.DoubleValue
	FloatValue  = wrappers.FloatValue
	BoolValue   = wrappers.BoolValue
	Int32Value  = wrappers.Int32Value
	Int64Value  = wrappers.Int64Value
	UInt32Value = wrappers.UInt32Value
	UInt64Value = wrappers.UInt64Value
	StringValue = wrappers.StringValue
	BytesValue  = wrappers.BytesValue
)
