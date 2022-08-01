package runtime

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/gogo/protobuf/jsonpb"
	"google.golang.org/protobuf/runtime/protoiface"
)

// TODO: get this working correctly
func MarshalProtoResponse(message protoiface.MessageV1) (interface{}, error) {
	marshaler := jsonpb.Marshaler{EnumsAsInts: true}

	buf := new(bytes.Buffer)
	if err := marshaler.Marshal(buf, message); err != nil {
		return nil, err
	}

	var data interface{}
	err := json.NewDecoder(buf).Decode(&data)

	fmt.Printf("%#v", data)
	return data, err
}

func derefValue(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}

func MarshalResponse(resp interface{}) interface{} {
	// If response is nil, nothing to do.
	if resp == nil {
		return nil
	}
	v := derefValue(reflect.ValueOf(resp))
	switch v.Kind() {
	case reflect.Struct:
		return marshalStruct(v)
	case reflect.Map:
		return marshalMap(v)
	case reflect.Slice:
		return marshalSlice(v)
	default:
		return primitive(v)
	}
}

// Marshal reflect value to []interface{} with lower camel case field
func marshalSlice(v reflect.Value) []interface{} {
	size := v.Len()
	ret := make([]interface{}, size)

	for i := 0; i < size; i++ {
		vv := derefValue(v.Index(i))
		switch vv.Kind() {
		case reflect.Struct:
			ret[i] = marshalStruct(vv)
		case reflect.Map:
			ret[i] = marshalMap(vv)
		case reflect.Slice:
			ret[i] = marshalSlice(vv)
		default:
			ret[i] = primitive(vv)
		}
	}
	return ret
}

// Marshal reflect value to map[string]interface{} with lower camel case field
func marshalStruct(v reflect.Value) map[string]interface{} {
	ret := make(map[string]interface{})
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		name := getTagName(t.Field(i).Tag)
		if len(name) == 0 {
			continue
		}

		vv := derefValue(v.Field(i))

		switch vv.Kind() {
		case reflect.Struct:
			ret[name] = marshalStruct(vv)
		case reflect.Map:
			ret[name] = marshalMap(vv)
		case reflect.Slice:
			ret[name] = marshalSlice(vv)
		default:
			ret[name] = primitive(vv)
		}
	}
	return ret
}

func getTagName(tag reflect.StructTag) (name string) {

	protoTag := tag.Get("protobuf")
	options := strings.Split(protoTag, ",")

	for _, option := range options {
		if strings.HasPrefix(option, "name") {
			name = option[5:]
		}
		if strings.HasPrefix(option, "json") {
			name = option[5:]
			break
		}
	}

	return
}

type mapValue struct {
	Key   interface{} `json:"key"`
	Value interface{} `json:"value"`
}

// Marshal reflect value to []mapValue with lower camel case field
// Note that in GraphQL, Protocol Buffers map structure should be marshaled to an array of key-value object
func marshalMap(v reflect.Value) []mapValue {
	var ret []mapValue

	iter := v.MapRange()
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		kk := derefValue(key)
		vv := derefValue(value)

		mapItem := mapValue{}

		switch kk.Kind() {
		case reflect.Struct:
			mapItem.Key = marshalStruct(kk)
		case reflect.Map:
			mapItem.Key = marshalMap(kk)
		case reflect.Slice:
			mapItem.Key = marshalSlice(kk)
		default:
			mapItem.Key = primitive(kk)
		}

		switch vv.Kind() {
		case reflect.Struct:
			mapItem.Value = marshalStruct(vv)
		case reflect.Map:
			mapItem.Value = marshalMap(vv)
		case reflect.Slice:
			mapItem.Value = marshalSlice(vv)
		default:
			mapItem.Value = primitive(vv)
		}
		ret = append(ret, mapItem)
	}

	return ret
}

// Get primitive type value
// Protobuf in Go only use a few scalar types.
// See: https://developers.google.com/protocol-buffers/docs/proto3#scalar
func primitive(v reflect.Value) interface{} {
	// Guard by cheking IsValid due to prevent panic in runtime
	if !v.IsValid() {
		return nil
	}

	switch v.Type().Kind() {
	case reflect.String:
		return v.String()
	case reflect.Bool:
		return v.Bool()
	case reflect.Int:
		return int(v.Int())
	case reflect.Int32:
		return int32(v.Int())
	case reflect.Int64:
		return int64(v.Int())
	case reflect.Uint:
		return uint(v.Uint())
	case reflect.Uint32:
		return uint32(v.Uint())
	case reflect.Uint64:
		return uint64(v.Uint())
	case reflect.Float32:
		return float32(v.Float())
	case reflect.Float64:
		return float64(v.Float())
	default:
		return nil
	}
}
