package runtime

import (
	"reflect"
	"strings"
)

func getTagName(tag reflect.StructTag) (name string) {

	name, protoOK := getProtoTagName(tag)
	if !protoOK {
		name, _ = getJsonTagName(tag)
	}

	return
}

func getProtoTagName(tag reflect.StructTag) (name string, ok bool) {
	protoTag, protoOK := tag.Lookup("protobuf")
	if !protoOK {
		return
	}

	namePrefix := "name="
	jsonPrefix := "json="

	options := strings.Split(protoTag, ",")
	for _, option := range options {
		if strings.HasPrefix(option, jsonPrefix) {
			name = option[len(jsonPrefix):]
			if ok = len(name) > 0; ok {
				return // found the first valid `jsonPrefix` so we are able to return
			}
		} else if !ok && strings.HasPrefix(option, namePrefix) {
			name = option[len(namePrefix):]
			ok = len(name) > 0
		}
	}

	return
}

func getJsonTagName(tag reflect.StructTag) (name string, ok bool) {
	jsonTag, jsonOK := tag.Lookup("json")
	if !jsonOK {
		return
	}

	name = strings.Split(jsonTag, ",")[0]
	ok = len(name) > 0
	return
}
