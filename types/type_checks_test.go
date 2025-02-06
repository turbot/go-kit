package types

import "testing"

type typeCheckTest struct {
	name     string
	input    any
	expected bool
}

type typeCheckInput struct {
	_int     int
	_int8    int8
	_int16   int16
	_int32   int32
	_int64   int64
	_uint    uint
	_uint8   uint8
	_uint16  uint16
	_uint32  uint32
	_uint64  uint64
	_float32 float32
	_float64 float64
	_bool    bool
	_slice   []string
	_array   [3]string
	_map     map[string]string
	_struct  structType
	_enum    enumType
	_func    func()
}

var typeCheckInputData = typeCheckInput{
	_int:     100,
	_int8:    int8(100),
	_int16:   int16(100),
	_int32:   int32(100),
	_int64:   int64(100),
	_uint:    uint(100),
	_uint8:   uint8(100),
	_uint16:  uint16(100),
	_uint32:  uint32(100),
	_uint64:  uint64(100),
	_float32: float32(100.1234),
	_float64: float64(100.123456),
	_bool:    true,
	_slice:   []string{"hello", "world"},
	_array:   [3]string{"one", "two", "three"},
	_map: map[string]string{
		"hello": "world",
	},
	_enum:   enumType_a,
	_func:   func() {},
	_struct: structType{"Bob"},
}

var testCasesIsSimpleType = []typeCheckTest{
	{"int", typeCheckInputData._int, true},
	{"int8", typeCheckInputData._int8, true},
	{"int16", typeCheckInputData._int16, true},
	{"int32", typeCheckInputData._int32, true},
	{"int64", typeCheckInputData._int64, true},
	{"uint", typeCheckInputData._uint, true},
	{"uint8", typeCheckInputData._uint8, true},
	{"uint16", typeCheckInputData._uint16, true},
	{"uint32", typeCheckInputData._uint32, true},
	{"uint64", typeCheckInputData._uint64, true},
	{"float32", typeCheckInputData._float32, true},
	{"float64", typeCheckInputData._float64, true},
	{"bool", typeCheckInputData._bool, true},
	{"slice", typeCheckInputData._slice, false},
	{"array", typeCheckInputData._array, false},
	{"map", typeCheckInputData._map, false},
	{"enum", typeCheckInputData._enum, true},
	{"func", typeCheckInputData._func, false},
	{"struct", typeCheckInputData._struct, false},
}

var testCasesIsCollectionType = []typeCheckTest{
	{"int", typeCheckInputData._int, false},
	{"int8", typeCheckInputData._int8, false},
	{"int16", typeCheckInputData._int16, false},
	{"int32", typeCheckInputData._int32, false},
	{"int64", typeCheckInputData._int64, false},
	{"uint", typeCheckInputData._uint, false},
	{"uint8", typeCheckInputData._uint8, false},
	{"uint16", typeCheckInputData._uint16, false},
	{"uint32", typeCheckInputData._uint32, false},
	{"uint64", typeCheckInputData._uint64, false},
	{"float32", typeCheckInputData._float32, false},
	{"float64", typeCheckInputData._float64, false},
	{"bool", typeCheckInputData._bool, false},
	{"slice", typeCheckInputData._slice, true},
	{"array", typeCheckInputData._array, true},
	{"map", typeCheckInputData._map, true},
	{"enum", typeCheckInputData._enum, false},
	{"func", typeCheckInputData._func, false},
	{"struct", typeCheckInputData._struct, false},
}

func TestIsSimpleType(t *testing.T) {
	for _, test := range testCasesIsSimpleType {
		output := IsSimpleType(test.input)
		if output != test.expected {
			t.Errorf("Test: '%s' FAILED: expected: %t, got %t", test.name, test.expected, output)
		}
	}
}

func TestIsCollectionType(t *testing.T) {
	for _, test := range testCasesIsCollectionType {
		output := IsCollectionType(test.input)
		if output != test.expected {
			t.Errorf("Test: '%s' FAILED: expected: %t, got %t", test.name, test.expected, output)
		}
	}
}
