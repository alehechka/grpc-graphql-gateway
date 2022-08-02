package runtime_test

import (
	"testing"

	"github.com/alehechka/grpc-graphql-gateway/runtime"
	"github.com/alehechka/grpc-graphql-gateway/runtime/testpb"
	"github.com/stretchr/testify/assert"
)

func assertStruct(t *testing.T, s *testpb.A) {
	assert.NotNil(t, s)
	assert.Equal(t, "string", s.StringValue)
	assert.Equal(t, int64(1), s.IntValue)
	if assert.Len(t, s.StringSlice, 3) {
		assert.Equal(t, "A", s.StringSlice[0])
		assert.Equal(t, "B", s.StringSlice[1])
		assert.Equal(t, "C", s.StringSlice[2])
	}
	assert.Equal(t, "string", s.NestedStruct.StringValue)
	assert.Equal(t, int64(1), s.NestedStruct.IntValue)
	if assert.Len(t, s.NestedStruct.StringSlice, 3) {
		assert.Equal(t, "A", s.NestedStruct.StringSlice[0])
		assert.Equal(t, "B", s.NestedStruct.StringSlice[1])
		assert.Equal(t, "C", s.NestedStruct.StringSlice[2])
	}
	if assert.Len(t, s.StructSlice, 2) {
		assert.Equal(t, "string01", s.StructSlice[0].StringValue)
		assert.Equal(t, int64(2), s.StructSlice[0].IntValue)
		assert.Equal(t, "string02", s.StructSlice[1].StringValue)
		assert.Equal(t, int64(3), s.StructSlice[1].IntValue)
	}
}

func TestMarshalRequest(t *testing.T) {
	data := map[string]interface{}{
		"string_value": "string",
		"int_value":    1,
		"string_slice": []string{"A", "B", "C"},
		"nested_struct": map[string]interface{}{
			"string_value": "string",
			"int_value":    1,
			"string_slice": []string{"A", "B", "C"},
		},
		"struct_slice": []testpb.C{
			{
				StringValue: "string01",
				IntValue:    2,
			},
			{
				StringValue: "string02",
				IntValue:    3,
			},
		},
	}
	var v testpb.A
	err := runtime.MarshalRequest(data, &v)
	assert.NoError(t, err)
	assertStruct(t, &v)
}

func TestMarshalRequestWithCamelCaseInput(t *testing.T) {
	data := map[string]interface{}{
		"stringValue": "string",
		"intValue":    1,
		"stringSlice": []string{"A", "B", "C"},
		"nestedStruct": map[string]interface{}{
			"valueString": "string",
			"valueInt":    1,
			"sliceString": []string{"A", "B", "C"},
		},
		"structSlice": []testpb.C{
			{
				StringValue: "string01",
				IntValue:    2,
			},
			{
				StringValue: "string02",
				IntValue:    3,
			},
		},
	}
	var v testpb.A
	err := runtime.MarshalRequest(data, &v)
	assert.NoError(t, err)
	assertStruct(t, &v)
}
