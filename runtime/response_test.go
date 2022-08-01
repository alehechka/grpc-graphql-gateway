package runtime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type exampleStruct struct {
	UserId int64  `protobuf:"name=user_id,json=userId" json:"user_id,omitempty"`
	Name   string `protobuf:"name=name" json:"name,omitempty"`

	TestInt   int   `protobuf:"name=test_int,json=testInt" json:"test_int,omitempty"`
	TestInt32 int32 `protobuf:"name=test_int32,json=testInt32" json:"test_int32,omitempty"`
	TestInt64 int64 `protobuf:"name=test_int64,json=testInt64" json:"test_int64,omitempty"`

	TestUint   uint   `protobuf:"name=test_uint,json=testUint" json:"test_uint,omitempty"`
	TestUint32 uint32 `protobuf:"name=test_uint32,json=testUint32" json:"test_uint32,omitempty"`
	TestUint64 uint64 `protobuf:"name=test_uint64,json=testUint64" json:"test_uint64,omitempty"`

	TestFloat32 float32 `protobuf:"name=test_float32,json=testFloat32" json:"test_float32,omitempty"`
	TestFloat64 float64 `protobuf:"name=test_float64,json=testFloat64" json:"test_float64,omitempty"`

	TestBool bool `protobuf:"name=test_bool,json=testBool" json:"test_bool,omitempty"`

	SubStruct *exampleSubStruct   `protobuf:"name=sub_struct,json=subStruct" json:"sub_struct,omitempty"`
	SubSlice  []*exampleSubStruct `protobuf:"name=sub_slice,json=subSlice" json:"sub_slice,omitempty"`
	SubMap    map[string]int64    `protobuf:"name=sub_map,json=subMap" json:"sub_map,omitempty"`
}

type exampleSubStruct struct {
	SomeData string `protobuf:"name=some_data,json=someData" json:"some_data,omitempty"`
}

func assertMap(t *testing.T, actual map[string]interface{}, expects map[string]interface{}) {
	for key, val := range expects {
		v, ok := actual[key]
		assert.True(t, ok, "actual map does not exist for "+key)
		assert.Equal(t, val, v, "assertion failed for "+key)
	}
}

func TestMarshalStruct(t *testing.T) {
	e := &exampleStruct{
		UserId:      100,
		Name:        "example",
		TestInt:     1,
		TestInt32:   1,
		TestInt64:   1,
		TestUint:    1,
		TestUint32:  1,
		TestUint64:  1,
		TestFloat32: 1,
		TestFloat64: 1,
		TestBool:    true,
		SubStruct: &exampleSubStruct{
			SomeData: "some",
		},
		SubSlice: []*exampleSubStruct{
			{SomeData: "some_slice"},
		},
	}
	v := MarshalResponse(e)
	m, ok := v.(map[string]interface{})
	if !assert.True(t, ok) {
		t.FailNow()
	}
	assertMap(t, m, map[string]interface{}{
		"userId":      int64(100),
		"name":        "example",
		"testInt":     1,
		"testInt32":   int32(1),
		"testInt64":   int64(1),
		"testUint":    uint(1),
		"testUint32":  uint32(1),
		"testUint64":  uint64(1),
		"testFloat32": float32(1),
		"testFloat64": float64(1),
		"testBool":    true,
	})

	a, ok := m["subStruct"]
	if !assert.True(t, ok) {
		t.FailNow()
	}
	aa, ok := a.(map[string]interface{})
	if !assert.True(t, ok) {
		t.FailNow()
	}
	assertMap(t, aa, map[string]interface{}{
		"someData": "some",
	})

	b, ok := m["subSlice"]
	if !assert.True(t, ok) {
		t.FailNow()
	}
	bb, ok := b.([]interface{})
	if !assert.True(t, ok) {
		t.FailNow()
	}
	if !assert.Len(t, bb, 1) {
		t.FailNow()
	}
	bbb, ok := bb[0].(map[string]interface{})
	if !assert.True(t, ok) {
		t.FailNow()
	}
	assertMap(t, bbb, map[string]interface{}{
		"someData": "some_slice",
	})
}

func TestMarshalSlice(t *testing.T) {
	e := []*exampleStruct{
		{
			UserId:      100,
			Name:        "example",
			TestInt:     1,
			TestInt32:   1,
			TestInt64:   1,
			TestUint:    1,
			TestUint32:  1,
			TestUint64:  1,
			TestFloat32: 1,
			TestFloat64: 1,
			TestBool:    true,
			SubStruct: &exampleSubStruct{
				SomeData: "some",
			},
			SubSlice: []*exampleSubStruct{
				{SomeData: "some_slice"},
			},
		},
		{
			UserId:      100,
			Name:        "example",
			TestInt:     1,
			TestInt32:   1,
			TestInt64:   1,
			TestUint:    1,
			TestUint32:  1,
			TestUint64:  1,
			TestFloat32: 1,
			TestFloat64: 1,
			TestBool:    true,
			SubStruct: &exampleSubStruct{
				SomeData: "some",
			},
			SubSlice: []*exampleSubStruct{
				{SomeData: "some_slice"},
			},
		},
	}
	v := MarshalResponse(e)
	vv, ok := v.([]interface{})
	if !assert.True(t, ok) {
		t.FailNow()
	}
	if !assert.Len(t, vv, 2) {
		t.FailNow()
	}
	for _, v := range vv {
		m, ok := v.(map[string]interface{})
		if !assert.True(t, ok) {
			t.FailNow()
		}
		assertMap(t, m, map[string]interface{}{
			"userId":      int64(100),
			"name":        "example",
			"testInt":     1,
			"testInt32":   int32(1),
			"testInt64":   int64(1),
			"testUint":    uint(1),
			"testUint32":  uint32(1),
			"testUint64":  uint64(1),
			"testFloat32": float32(1),
			"testFloat64": float64(1),
			"testBool":    true,
		})

		a, ok := m["subStruct"]
		if !assert.True(t, ok) {
			t.FailNow()
		}
		aa, ok := a.(map[string]interface{})
		if !assert.True(t, ok) {
			t.FailNow()
		}
		assertMap(t, aa, map[string]interface{}{
			"someData": "some",
		})

		b, ok := m["subSlice"]
		if !assert.True(t, ok) {
			t.FailNow()
		}
		bb, ok := b.([]interface{})
		if !assert.True(t, ok) {
			t.FailNow()
		}
		if !assert.Len(t, bb, 1) {
			t.FailNow()
		}
		bbb, ok := bb[0].(map[string]interface{})
		if !assert.True(t, ok) {
			t.FailNow()
		}
		assertMap(t, bbb, map[string]interface{}{
			"someData": "some_slice",
		})
	}
}

func TestMarshalMap(t *testing.T) {
	e := &exampleStruct{
		SubMap: map[string]int64{
			"item01": 1,
			"item02": 2,
		},
	}
	v := MarshalResponse(e)
	m, ok := v.(map[string]interface{})
	if !assert.True(t, ok) {
		t.FailNow()
	}

	a, ok := m["subMap"]
	if !assert.True(t, ok) {
		t.FailNow()
	}
	aa, ok := a.([]mapValue)
	if !assert.True(t, ok) {
		t.FailNow()
	}
	if !assert.Len(t, aa, 2) {
		t.FailNow()
	}
	// Map does not consier key order so we need to assert with contains
	assert.Contains(t, []string{"item01", "item02"}, aa[0].Key)
	assert.Contains(t, []int64{1, 2}, aa[0].Value)
	assert.Contains(t, []string{"item01", "item02"}, aa[1].Key)
	assert.Contains(t, []int64{1, 2}, aa[1].Value)
}
