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

func getProtoTagName(tag reflect.StructTag) (name string, exists bool) {
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
			exists = true
			break // found `jsonPrefix` so we are able to break from the loop and return
		} else if len(name) == 0 && strings.HasPrefix(option, namePrefix) {
			name = option[len(namePrefix):]
			exists = true
			continue // found `namePrefix`, however `jsonPrefix` is the priority, so we must keep searching
		}
	}

	return
}

func getJsonTagName(tag reflect.StructTag) (name string, exists bool) {
	jsonTag, jsonOK := tag.Lookup("json")
	if !jsonOK {
		return
	}

	name = strings.Split(jsonTag, ",")[0]
	exists = len(name) > 0
	return
}
