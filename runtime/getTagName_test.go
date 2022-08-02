package runtime

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getJsonTagName_NoCommas(t *testing.T) {
	data := struct {
		item string `json:"item"`
	}{item: "hello"}

	name, exists := getJsonTagName(reflect.ValueOf(data).Type().Field(0).Tag)

	assert.Equal(t, "item", name)
	assert.True(t, exists)
}

func Test_getJsonTagName_OmitEmpty(t *testing.T) {
	data := struct {
		item string `json:"item,omitempty"`
	}{item: "hello"}

	name, exists := getJsonTagName(reflect.ValueOf(data).Type().Field(0).Tag)

	assert.Equal(t, "item", name)
	assert.True(t, exists)
}

func Test_getJsonTagName_StarterComma(t *testing.T) {
	data := struct {
		item string `json:",omitempty"`
	}{item: "hello"}

	name, exists := getJsonTagName(reflect.ValueOf(data).Type().Field(0).Tag)

	assert.Equal(t, "", name)
	assert.False(t, exists)
}

func Test_getJsonTagName_EndComma(t *testing.T) {
	data := struct {
		item string `json:"item,"`
	}{item: "hello"}

	name, exists := getJsonTagName(reflect.ValueOf(data).Type().Field(0).Tag)

	assert.Equal(t, "item", name)
	assert.True(t, exists)
}

func Test_getJsonTagName_MultipleCommas(t *testing.T) {
	data := struct {
		item string `json:"item,omitempty,required"`
	}{item: "hello"}

	name, exists := getJsonTagName(reflect.ValueOf(data).Type().Field(0).Tag)

	assert.Equal(t, "item", name)
	assert.True(t, exists)
}

func Test_getJsonTagName_Empty(t *testing.T) {
	data := struct {
		item string `json:""`
	}{item: "hello"}

	name, exists := getJsonTagName(reflect.ValueOf(data).Type().Field(0).Tag)

	assert.Equal(t, "", name)
	assert.False(t, exists)
}

func Test_getJsonTagName_NoJsonTag(t *testing.T) {
	data := struct {
		item string `graphql:"item,omitempty"`
	}{item: "hello"}

	name, exists := getJsonTagName(reflect.ValueOf(data).Type().Field(0).Tag)

	assert.Equal(t, "", name)
	assert.False(t, exists)
}

func Test_getProtoTagName_NoJson(t *testing.T) {
	data := struct {
		item string `protobuf:"bytes,2,opt,name=item,proto3"`
	}{item: "hello"}

	name, exists := getProtoTagName(reflect.ValueOf(data).Type().Field(0).Tag)

	assert.Equal(t, "item", name)
	assert.True(t, exists)
}

func Test_getProtoTagName_NameAndJson(t *testing.T) {
	data := struct {
		item string `protobuf:"bytes,2,opt,name=item,json=itemName,proto3"`
	}{item: "hello"}

	name, exists := getProtoTagName(reflect.ValueOf(data).Type().Field(0).Tag)

	assert.Equal(t, "itemName", name)
	assert.True(t, exists)
}

func Test_getProtoTagName_NoName(t *testing.T) {
	data := struct {
		item string `protobuf:"bytes,2,opt,json=item,proto3"`
	}{item: "hello"}

	name, exists := getProtoTagName(reflect.ValueOf(data).Type().Field(0).Tag)

	assert.Equal(t, "item", name)
	assert.True(t, exists)
}

func Test_getProtoTagName_NoNameNoJson(t *testing.T) {
	data := struct {
		item string `protobuf:"bytes,2,opt,proto3"`
	}{item: "hello"}

	name, exists := getProtoTagName(reflect.ValueOf(data).Type().Field(0).Tag)

	assert.Equal(t, "", name)
	assert.False(t, exists)
}

func Test_getProtoTagName_NoProtobufTag(t *testing.T) {
	data := struct {
		item string `graphql:"item,omitempty"`
	}{item: "hello"}

	name, exists := getProtoTagName(reflect.ValueOf(data).Type().Field(0).Tag)

	assert.Equal(t, "", name)
	assert.False(t, exists)
}

func Test_getProtoTagName_EmptyTag(t *testing.T) {
	data := struct {
		item string `protobuf:""`
	}{item: "hello"}

	name, exists := getProtoTagName(reflect.ValueOf(data).Type().Field(0).Tag)

	assert.Equal(t, "", name)
	assert.False(t, exists)
}

func Test_getProtoTagName_OnlyCommas(t *testing.T) {
	data := struct {
		item string `protobuf:","`
	}{item: "hello"}

	name, exists := getProtoTagName(reflect.ValueOf(data).Type().Field(0).Tag)

	assert.Equal(t, "", name)
	assert.False(t, exists)
}

func Test_getProtoTagName_FindsFirstName(t *testing.T) {
	data := struct {
		item string `protobuf:"bytes,2,opt,name=item,name=itemName,proto3"`
	}{item: "hello"}

	name, exists := getProtoTagName(reflect.ValueOf(data).Type().Field(0).Tag)

	assert.Equal(t, "item", name)
	assert.True(t, exists)
}

func Test_getProtoTagName_FindsFirstJson(t *testing.T) {
	data := struct {
		item string `protobuf:"bytes,2,opt,json=item,json=itemName,proto3"`
	}{item: "hello"}

	name, exists := getProtoTagName(reflect.ValueOf(data).Type().Field(0).Tag)

	assert.Equal(t, "item", name)
	assert.True(t, exists)
}

func Test_getTagName_ProtoName(t *testing.T) {
	data := struct {
		item string `protobuf:"bytes,2,opt,name=item,proto3"`
	}{item: "hello"}

	name := getTagName(reflect.ValueOf(data).Type().Field(0).Tag)

	assert.Equal(t, "item", name)
}

func Test_getTagName_ProtoJson(t *testing.T) {
	data := struct {
		item string `protobuf:"bytes,2,opt,name=item,json=itemName,proto3"`
	}{item: "hello"}

	name := getTagName(reflect.ValueOf(data).Type().Field(0).Tag)

	assert.Equal(t, "itemName", name)
}

func Test_getTagName_NoProto_Json(t *testing.T) {
	data := struct {
		item string `json:"item"`
	}{item: "hello"}

	name := getTagName(reflect.ValueOf(data).Type().Field(0).Tag)

	assert.Equal(t, "item", name)
}

func Test_getTagName_InvalidProto_Json(t *testing.T) {
	data := struct {
		item string `protobuf:"bytes,2,opt,proto3" json:"item"`
	}{item: "hello"}

	name := getTagName(reflect.ValueOf(data).Type().Field(0).Tag)

	assert.Equal(t, "item", name)
}
