package types

import (
	"reflect"
	"testing"
	"time"
)

type toStringFunc func(x interface{}) string
type numberConversionTest struct {
	name     string
	input    interface{}
	expected interface{}
}
type enumType string

const (
	enumType_a enumType = "a-val"
)

type structType struct {
	a string
}

type numberConversionInput struct {
	_int64           int64
	_int32           int32
	_int16           int32
	_int8            int8
	_int             int
	_uint64          uint64
	_uint32          uint32
	_uint16          uint16
	_uint8           uint8
	_uint            uint
	_byte            byte
	_float64         float64
	_float32         float32
	_bool_true       bool
	_bool_false      bool
	_string_true     string
	_string_TRUE     string
	_string_on       string
	_string_enabled  string
	_string_false    string
	_string_off      string
	_string_disabled string
	_int_as_string   string
	_float_as_string string
	_random_string   string
	_map             map[string]string
	_int_1           int
	_struct          structType
	_enum            enumType
}

var numberConversionInputData = numberConversionInput{
	int64(1234567),
	int32(1234567),
	int32(1234567),
	int8(100),
	int(1234567),
	uint64(1234567),
	uint32(1234567),
	uint16(12345),
	uint8(100),
	uint(1234567),
	byte(100),
	float64(1234.12345678),
	float32(1234.1234),
	true,
	false,
	"true",
	"TRUE",
	"on",
	"enabled",
	"false",
	"off",
	"disabled",
	"1234567",
	"1234567.12345678",
	"invalid",
	map[string]string{
		"a": "b",
	},
	1234567,
	structType{"A"},
	enumType_a,
}

func executeToStringTest(t *testing.T, test numberConversionTest, conv toStringFunc) {
	defer func() {
		if r := recover(); r != nil {
			if test.expected != "PANIC" {
				t.Errorf(`Test: '%s'' FAILED : unexpected panic %v`, test.name, r)
			}
		}
	}()

	output := conv(test.input)
	if !reflect.DeepEqual(output, test.expected) {
		t.Errorf(`Test: '%s'' FAILED : expected %s, got %s`, test.name, test.expected, output)
	}
}

var testCasesToString = []numberConversionTest{
	{"int64", numberConversionInputData._int64, "1234567"},
	{"int32", numberConversionInputData._int32, "1234567"},
	{"int16", numberConversionInputData._int16, "1234567"},
	{"int8", numberConversionInputData._int8, "100"},
	{"int", numberConversionInputData._int, "1234567"},
	{"uint64", numberConversionInputData._uint64, "1234567"},
	{"uint32", numberConversionInputData._uint32, "1234567"},
	{"uint16", numberConversionInputData._uint16, "12345"},
	{"uint8", numberConversionInputData._uint8, "100"},
	{"uint", numberConversionInputData._uint, "1234567"},
	{"byte", numberConversionInputData._byte, "100"},
	{"float64", numberConversionInputData._float64, "1234.12345678"},
	{"float32", numberConversionInputData._float32, "1234.1234"},
	{"*int64", &numberConversionInputData._int64, "1234567"},
	{"*int32", &numberConversionInputData._int32, "1234567"},
	{"*int16", &numberConversionInputData._int16, "1234567"},
	{"*int8", &numberConversionInputData._int8, "100"},
	{"*int", &numberConversionInputData._int, "1234567"},
	{"*uint64", &numberConversionInputData._uint64, "1234567"},
	{"*uint32", &numberConversionInputData._uint32, "1234567"},
	{"*uint16", &numberConversionInputData._uint16, "12345"},
	{"*uint8", &numberConversionInputData._uint8, "100"},
	{"*uint", &numberConversionInputData._uint, "1234567"},
	{"*byte", &numberConversionInputData._byte, "100"},
	{"*float64", &numberConversionInputData._float64, "1234.12345678"},
	{"*float32", &numberConversionInputData._float32, "1234.1234"},
	{"map", numberConversionInputData._map, "map[a:b]"},
	{"struct", numberConversionInputData._struct, "{a:A}"},
	{"enum", numberConversionInputData._enum, "a-val"},

	// Nil pointer cases
	{"nil *string", (*string)(nil), ""},
	{"nil *bool", (*bool)(nil), "<nil>"},
	{"nil *int", (*int)(nil), "0"},
	{"nil *float64", (*float64)(nil), "0"},
	{"nil *time.Time", (*time.Time)(nil), ""},
}

func TestToString(t *testing.T) {
	for _, test := range testCasesToString {
		executeToStringTest(t, test, ToString)
	}
}

var testCasesFloatToString = []numberConversionTest{
	{"int64", numberConversionInputData._int64, "PANIC"},
	{"int32", numberConversionInputData._int32, "PANIC"},
	{"int16", numberConversionInputData._int16, "PANIC"},
	{"int8", numberConversionInputData._int8, "PANIC"},
	{"int", numberConversionInputData._int, "PANIC"},
	{"uint64", numberConversionInputData._uint64, "PANIC"},
	{"uint32", numberConversionInputData._uint32, "PANIC"},
	{"uint16", numberConversionInputData._uint16, "PANIC"},
	{"uint8", numberConversionInputData._uint8, "PANIC"},
	{"uint", numberConversionInputData._uint, "PANIC"},
	{"byte", numberConversionInputData._byte, "PANIC"},
	{"float64", numberConversionInputData._float64, "1234.12345678"},
	{"float32", numberConversionInputData._float32, "1234.1234"},
	{"*int64", &numberConversionInputData._int64, "PANIC"},
	{"*int32", &numberConversionInputData._int32, "PANIC"},
	{"*int16", &numberConversionInputData._int16, "PANIC"},
	{"*int8", &numberConversionInputData._int8, "PANIC"},
	{"*int", &numberConversionInputData._int, "PANIC"},
	{"*uint64", &numberConversionInputData._uint64, "PANIC"},
	{"*uint32", &numberConversionInputData._uint32, "PANIC"},
	{"*uint16", &numberConversionInputData._uint16, "PANIC"},
	{"*uint8", &numberConversionInputData._uint8, "PANIC"},
	{"*uint", &numberConversionInputData._uint, "PANIC"},
	{"*byte", &numberConversionInputData._byte, "PANIC"},
	{"*float64", &numberConversionInputData._float64, "1234.12345678"},
	{"*float32", &numberConversionInputData._float32, "1234.1234"},
	{"map", numberConversionInputData._map, "PANIC"},
	{"floatAsString", numberConversionInputData._float_as_string, "1234567.12345678"},
}

func TestFloatToString(t *testing.T) {
	for _, test := range testCasesFloatToString {
		executeToStringTest(t, test, FloatToString)
	}
}

var testCasesIntToString = []numberConversionTest{
	{"int64", numberConversionInputData._int64, "1234567"},
	{"int32", numberConversionInputData._int32, "1234567"},
	{"int16", numberConversionInputData._int16, "1234567"},
	{"int8", numberConversionInputData._int8, "100"},
	{"int", numberConversionInputData._int, "1234567"},
	{"uint64", numberConversionInputData._uint64, "1234567"},
	{"uint32", numberConversionInputData._uint32, "1234567"},
	{"uint16", numberConversionInputData._uint16, "12345"},
	{"uint8", numberConversionInputData._uint8, "100"},
	{"uint", numberConversionInputData._uint, "1234567"},
	{"byte", numberConversionInputData._byte, "100"},
	{"float64", numberConversionInputData._float64, "PANIC"},
	{"float32", numberConversionInputData._float32, "PANIC"},
	{"*int64", &numberConversionInputData._int64, "1234567"},
	{"*int32", &numberConversionInputData._int32, "1234567"},
	{"*int16", &numberConversionInputData._int16, "1234567"},
	{"*int8", &numberConversionInputData._int8, "100"},
	{"*int", &numberConversionInputData._int, "1234567"},
	{"*uint64", &numberConversionInputData._uint64, "1234567"},
	{"*uint32", &numberConversionInputData._uint32, "1234567"},
	{"*uint16", &numberConversionInputData._uint16, "12345"},
	{"*uint8", &numberConversionInputData._uint8, "100"},
	{"*uint", &numberConversionInputData._uint, "1234567"},
	{"*byte", &numberConversionInputData._byte, "100"},
	{"*float64", &numberConversionInputData._float64, "PANIC"},
	{"*float32", &numberConversionInputData._float32, "PANIC"},
	{"map", numberConversionInputData._map, "PANIC"},
	{"intAsString", numberConversionInputData._int_as_string, "1234567"},
	{"floatAsString", numberConversionInputData._float_as_string, "PANIC"},
}

func TestIntToString(t *testing.T) {

	for _, test := range testCasesIntToString {
		executeToStringTest(t, test, IntToString)
	}
}

var testCasesToBool = []numberConversionTest{
	{"bool true", numberConversionInputData._bool_true, true},
	{"bool false", numberConversionInputData._bool_false, false},
	{"string true", numberConversionInputData._string_true, true},
	{"string TRUE", numberConversionInputData._string_TRUE, true},
	{"string on", numberConversionInputData._string_on, true},
	{"string enabled", numberConversionInputData._string_enabled, true},
	{"string false", numberConversionInputData._string_false, false},
	{"string off", numberConversionInputData._string_off, false},
	{"string disabled", numberConversionInputData._string_disabled, false},
	{"int 1", numberConversionInputData._int_1, "ERROR"},
	{"random string", numberConversionInputData._random_string, "ERROR"},
}

func TestToBool(t *testing.T) {
	for _, test := range testCasesToBool {
		output, err := ToBool(test.input)
		if err != nil {
			if test.expected != "ERROR" {
				t.Errorf(`Test: '%s'' FAILED : unexpected error %v`, test.name, err)
			}
			continue

		}
		if test.expected == "ERROR" {
			t.Errorf(`Test: '%s'' FAILED : expected error but did not get one`, test.name)
			continue
		}

		if output != test.expected.(bool) {
			t.Errorf(`Test: '%s'' FAILED : expected %v, got %v`, test.name, test.expected, output)
		}
	}
}

var testCasesToBoolPtr = []numberConversionTest{
	{"bool true", numberConversionInputData._bool_true, true},
	{"bool false", numberConversionInputData._bool_false, false},
	{"string true", numberConversionInputData._string_true, true},
	{"string TRUE", numberConversionInputData._string_TRUE, true},
	{"string on", numberConversionInputData._string_on, true},
	{"string enabled", numberConversionInputData._string_enabled, true},
	{"string false", numberConversionInputData._string_false, false},
	{"string off", numberConversionInputData._string_off, false},
	{"string disabled", numberConversionInputData._string_disabled, false},
	{"int 1", numberConversionInputData._int_1, nil},
	{"random string", numberConversionInputData._random_string, nil},
	{"random string", nil, nil},
}

func TestToBoolPtr(t *testing.T) {
	for _, test := range testCasesToBoolPtr {
		output := ToBoolPtr(test.input)
		if output == nil {
			if test.expected != nil {
				t.Errorf(`Test: '%s'' FAILED : unexpected nil result - expected %v`, test.name, test.expected)
			}
			return
		}

		if test.expected == nil && output != nil {
			t.Errorf(`Test: '%s'' FAILED : expected error but did not get one`, test.name)
		}

		if *output != test.expected.(bool) {
			t.Errorf(`Test: '%s'' FAILED : expected %v, got %v`, test.name, test.expected, output)
		}
	}
}

var testCasesToInt64 = []numberConversionTest{
	{"bool", numberConversionInputData._bool_false, "ERROR"},
	{"int 8", numberConversionInputData._int8, int64(numberConversionInputData._int8)},
	{"int 16", numberConversionInputData._int16, int64(numberConversionInputData._int16)},
	{"int 32", numberConversionInputData._int32, int64(numberConversionInputData._int32)},
	{"int 64", numberConversionInputData._int64, numberConversionInputData._int64},
	{"int", numberConversionInputData._int, int64(numberConversionInputData._int)},
	{"int as string", numberConversionInputData._int_as_string, int64(1234567)},
	{"uint 8", numberConversionInputData._uint8, int64(numberConversionInputData._uint8)},
	{"float", numberConversionInputData._float64, int64(1234)},
	{"invalid string", "FOO", "ERROR"},
}

func TestToInt64(t *testing.T) {
	for _, test := range testCasesToInt64 {
		output, err := ToInt64(test.input)
		if err != nil {
			if test.expected != "ERROR" {
				t.Errorf(`Test: '%s'' FAILED : unexpected error %v`, test.name, err)
			}
			continue
		}
		if test.expected == "ERROR" {
			t.Errorf(`Test: '%s'' FAILED : expected error but did not get one`, test.name)
			continue
		}

		if output != test.expected.(int64) {
			t.Errorf(`Test: '%s'' FAILED : expected %v, got %v`, test.name, test.expected, output)
		}
	}
}

var testCasesToFloat64 = []numberConversionTest{
	{"bool", numberConversionInputData._bool_false, "ERROR"},
	{"float 32", numberConversionInputData._float32, float64(numberConversionInputData._float32)},
	{"int 64", numberConversionInputData._int64, float64(numberConversionInputData._int64)},
	{"float 64", numberConversionInputData._float64, numberConversionInputData._float64},
	{"float as string", numberConversionInputData._float_as_string, float64(1234567.12345678)},
	{"invalid string", "FOO", "ERROR"},
}

func TestToFloat64(t *testing.T) {
	for _, test := range testCasesToFloat64 {
		output, err := ToFloat64(test.input)
		if err != nil {
			if test.expected != "ERROR" {
				t.Errorf(`Test: '%s'' FAILED : unexpected error %v`, test.name, err)
			}
			continue
		}
		if test.expected == "ERROR" {
			t.Errorf(`Test: '%s'' FAILED : expected error but did not get one`, test.name)
			continue
		}

		if output != test.expected.(float64) {
			t.Errorf(`Test: '%s'' FAILED : expected %v, got %v`, test.name, test.expected, output)
		}
	}
}

var testCasesStringSlice = [][]string{
	{"a", "b", "c", "d", "e"},
	{"a", "b", "", "", "e"},
}

func TestStringSlice(t *testing.T) {
	for idx, in := range testCasesStringSlice {
		if in == nil {
			continue
		}
		out := StringSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := StringValueSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesStringValueSlice = [][]*string{
	{String("a"), String("b"), nil, String("c")},
}

func TestStringValueSlice(t *testing.T) {
	for idx, in := range testCasesStringValueSlice {
		if in == nil {
			continue
		}
		out := StringValueSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != "" {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := StringSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != "" {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *in[i], *out2[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesStringMap = []map[string]string{
	{"a": "1", "b": "2", "c": "3"},
}

func TestStringMap(t *testing.T) {
	for idx, in := range testCasesStringMap {
		if in == nil {
			continue
		}
		out := StringMap(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := StringValueMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesBoolSlice = [][]bool{
	{true, true, false, false},
}

func TestBoolSlice(t *testing.T) {
	for idx, in := range testCasesBoolSlice {
		if in == nil {
			continue
		}
		out := BoolSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := BoolValueSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesBoolValueSlice = [][]*bool{}

func TestBoolValueSlice(t *testing.T) {
	for idx, in := range testCasesBoolValueSlice {
		if in == nil {
			continue
		}
		out := BoolValueSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := BoolSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := in[i], out2[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesBoolMap = []map[string]bool{
	{"a": true, "b": false, "c": true},
}

func TestBoolMap(t *testing.T) {
	for idx, in := range testCasesBoolMap {
		if in == nil {
			continue
		}
		out := BoolMap(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := BoolValueMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesUintSlice = [][]uint{
	{1, 2, 3, 4},
}

func TestUintSlice(t *testing.T) {
	for idx, in := range testCasesUintSlice {
		if in == nil {
			continue
		}
		out := UintSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := UintValueSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesUintValueSlice = [][]*uint{}

func TestUintValueSlice(t *testing.T) {
	for idx, in := range testCasesUintValueSlice {
		if in == nil {
			continue
		}
		out := UintValueSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := UintSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := in[i], out2[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesUintMap = []map[string]uint{
	{"a": 3, "b": 2, "c": 1},
}

func TestUintMap(t *testing.T) {
	for idx, in := range testCasesUintMap {
		if in == nil {
			continue
		}
		out := UintMap(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := UintValueMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesIntSlice = [][]int{
	{1, 2, 3, 4},
}

func TestIntSlice(t *testing.T) {
	for idx, in := range testCasesIntSlice {
		if in == nil {
			continue
		}
		out := IntSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := IntValueSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesIntValueSlice = [][]*int{}

func TestIntValueSlice(t *testing.T) {
	for idx, in := range testCasesIntValueSlice {
		if in == nil {
			continue
		}
		out := IntValueSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := IntSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := in[i], out2[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesIntMap = []map[string]int{
	{"a": 3, "b": 2, "c": 1},
}

func TestIntMap(t *testing.T) {
	for idx, in := range testCasesIntMap {
		if in == nil {
			continue
		}
		out := IntMap(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := IntValueMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesInt8Slice = [][]int8{
	{1, 2, 3, 4},
}

func TestInt8Slice(t *testing.T) {
	for idx, in := range testCasesInt8Slice {
		if in == nil {
			continue
		}
		out := Int8Slice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := Int8ValueSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesInt8ValueSlice = [][]*int8{}

func TestInt8ValueSlice(t *testing.T) {
	for idx, in := range testCasesInt8ValueSlice {
		if in == nil {
			continue
		}
		out := Int8ValueSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := Int8Slice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := in[i], out2[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesInt8Map = []map[string]int8{
	{"a": 3, "b": 2, "c": 1},
}

func TestInt8Map(t *testing.T) {
	for idx, in := range testCasesInt8Map {
		if in == nil {
			continue
		}
		out := Int8Map(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := Int8ValueMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesInt16Slice = [][]int16{
	{1, 2, 3, 4},
}

func TestInt16Slice(t *testing.T) {
	for idx, in := range testCasesInt16Slice {
		if in == nil {
			continue
		}
		out := Int16Slice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := Int16ValueSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesInt16ValueSlice = [][]*int16{}

func TestInt16ValueSlice(t *testing.T) {
	for idx, in := range testCasesInt16ValueSlice {
		if in == nil {
			continue
		}
		out := Int16ValueSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := Int16Slice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := in[i], out2[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesInt16Map = []map[string]int16{
	{"a": 3, "b": 2, "c": 1},
}

func TestInt16Map(t *testing.T) {
	for idx, in := range testCasesInt16Map {
		if in == nil {
			continue
		}
		out := Int16Map(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := Int16ValueMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesInt32Slice = [][]int32{
	{1, 2, 3, 4},
}

func TestInt32Slice(t *testing.T) {
	for idx, in := range testCasesInt32Slice {
		if in == nil {
			continue
		}
		out := Int32Slice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := Int32ValueSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesInt32ValueSlice = [][]*int32{}

func TestInt32ValueSlice(t *testing.T) {
	for idx, in := range testCasesInt32ValueSlice {
		if in == nil {
			continue
		}
		out := Int32ValueSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := Int32Slice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := in[i], out2[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesInt32Map = []map[string]int32{
	{"a": 3, "b": 2, "c": 1},
}

func TestInt32Map(t *testing.T) {
	for idx, in := range testCasesInt32Map {
		if in == nil {
			continue
		}
		out := Int32Map(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := Int32ValueMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesInt64Slice = [][]int64{
	{1, 2, 3, 4},
}

func TestInt64Slice(t *testing.T) {
	for idx, in := range testCasesInt64Slice {
		if in == nil {
			continue
		}
		out := Int64Slice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := Int64ValueSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesInt64ValueSlice = [][]*int64{}

func TestInt64ValueSlice(t *testing.T) {
	for idx, in := range testCasesInt64ValueSlice {
		if in == nil {
			continue
		}
		out := Int64ValueSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := Int64Slice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := in[i], out2[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesInt64Map = []map[string]int64{
	{"a": 3, "b": 2, "c": 1},
}

func TestInt64Map(t *testing.T) {
	for idx, in := range testCasesInt64Map {
		if in == nil {
			continue
		}
		out := Int64Map(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := Int64ValueMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesUint8Slice = [][]uint8{
	{1, 2, 3, 4},
}

func TestUint8Slice(t *testing.T) {
	for idx, in := range testCasesUint8Slice {
		if in == nil {
			continue
		}
		out := Uint8Slice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := Uint8ValueSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesUint8ValueSlice = [][]*uint8{}

func TestUint8ValueSlice(t *testing.T) {
	for idx, in := range testCasesUint8ValueSlice {
		if in == nil {
			continue
		}
		out := Uint8ValueSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := Uint8Slice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := in[i], out2[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesUint8Map = []map[string]uint8{
	{"a": 3, "b": 2, "c": 1},
}

func TestUint8Map(t *testing.T) {
	for idx, in := range testCasesUint8Map {
		if in == nil {
			continue
		}
		out := Uint8Map(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := Uint8ValueMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesUint16Slice = [][]uint16{
	{1, 2, 3, 4},
}

func TestUint16Slice(t *testing.T) {
	for idx, in := range testCasesUint16Slice {
		if in == nil {
			continue
		}
		out := Uint16Slice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := Uint16ValueSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesUint16ValueSlice = [][]*uint16{}

func TestUint16ValueSlice(t *testing.T) {
	for idx, in := range testCasesUint16ValueSlice {
		if in == nil {
			continue
		}
		out := Uint16ValueSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := Uint16Slice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := in[i], out2[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesUint16Map = []map[string]uint16{
	{"a": 3, "b": 2, "c": 1},
}

func TestUint16Map(t *testing.T) {
	for idx, in := range testCasesUint16Map {
		if in == nil {
			continue
		}
		out := Uint16Map(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := Uint16ValueMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesUint32Slice = [][]uint32{
	{1, 2, 3, 4},
}

func TestUint32Slice(t *testing.T) {
	for idx, in := range testCasesUint32Slice {
		if in == nil {
			continue
		}
		out := Uint32Slice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := Uint32ValueSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesUint32ValueSlice = [][]*uint32{}

func TestUint32ValueSlice(t *testing.T) {
	for idx, in := range testCasesUint32ValueSlice {
		if in == nil {
			continue
		}
		out := Uint32ValueSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := Uint32Slice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := in[i], out2[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesUint32Map = []map[string]uint32{
	{"a": 3, "b": 2, "c": 1},
}

func TestUint32Map(t *testing.T) {
	for idx, in := range testCasesUint32Map {
		if in == nil {
			continue
		}
		out := Uint32Map(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := Uint32ValueMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesUint64Slice = [][]uint64{
	{1, 2, 3, 4},
}

func TestUint64Slice(t *testing.T) {
	for idx, in := range testCasesUint64Slice {
		if in == nil {
			continue
		}
		out := Uint64Slice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := Uint64ValueSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesUint64ValueSlice = [][]*uint64{}

func TestUint64ValueSlice(t *testing.T) {
	for idx, in := range testCasesUint64ValueSlice {
		if in == nil {
			continue
		}
		out := Uint64ValueSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := Uint64Slice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := in[i], out2[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesUint64Map = []map[string]uint64{
	{"a": 3, "b": 2, "c": 1},
}

func TestUint64Map(t *testing.T) {
	for idx, in := range testCasesUint64Map {
		if in == nil {
			continue
		}
		out := Uint64Map(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := Uint64ValueMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesFloat32Slice = [][]float32{
	{1, 2, 3, 4},
}

func TestFloat32Slice(t *testing.T) {
	for idx, in := range testCasesFloat32Slice {
		if in == nil {
			continue
		}
		out := Float32Slice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := Float32ValueSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesFloat32ValueSlice = [][]*float32{}

func TestFloat32ValueSlice(t *testing.T) {
	for idx, in := range testCasesFloat32ValueSlice {
		if in == nil {
			continue
		}
		out := Float32ValueSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := Float32Slice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := in[i], out2[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesFloat32Map = []map[string]float32{
	{"a": 3, "b": 2, "c": 1},
}

func TestFloat32Map(t *testing.T) {
	for idx, in := range testCasesFloat32Map {
		if in == nil {
			continue
		}
		out := Float32Map(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := Float32ValueMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesFloat64Slice = [][]float64{
	{1, 2, 3, 4},
}

func TestFloat64Slice(t *testing.T) {
	for idx, in := range testCasesFloat64Slice {
		if in == nil {
			continue
		}
		out := Float64Slice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := Float64ValueSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesFloat64ValueSlice = [][]*float64{}

func TestFloat64ValueSlice(t *testing.T) {
	for idx, in := range testCasesFloat64ValueSlice {
		if in == nil {
			continue
		}
		out := Float64ValueSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if out[i] != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := Float64Slice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if *(out2[i]) != 0 {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := in[i], out2[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesFloat64Map = []map[string]float64{
	{"a": 3, "b": 2, "c": 1},
}

func TestFloat64Map(t *testing.T) {
	for idx, in := range testCasesFloat64Map {
		if in == nil {
			continue
		}
		out := Float64Map(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := Float64ValueMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesTimeSlice = [][]time.Time{
	{time.Now(), time.Now().AddDate(100, 0, 0)},
}

func TestTimeSlice(t *testing.T) {
	for idx, in := range testCasesTimeSlice {
		if in == nil {
			continue
		}
		out := TimeSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := TimeValueSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

var testCasesTimeValueSlice = [][]*time.Time{}

func TestTimeValueSlice(t *testing.T) {
	for idx, in := range testCasesTimeValueSlice {
		if in == nil {
			continue
		}
		out := TimeValueSlice(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if in[i] == nil {
				if !out[i].IsZero() {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := *(in[i]), out[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}

		out2 := TimeSlice(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out2 {
			if in[i] == nil {
				if !(out2[i]).IsZero() {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			} else {
				if e, a := in[i], out2[i]; e != a {
					t.Errorf("Unexpected value at idx %d", idx)
				}
			}
		}
	}
}

var testCasesTimeMap = []map[string]time.Time{
	{"a": time.Now().AddDate(-100, 0, 0), "b": time.Now()},
}

func TestTimeMap(t *testing.T) {
	for idx, in := range testCasesTimeMap {
		if in == nil {
			continue
		}
		out := TimeMap(in)
		if e, a := len(out), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		for i := range out {
			if e, a := in[i], *(out[i]); e != a {
				t.Errorf("Unexpected value at idx %d", idx)
			}
		}

		out2 := TimeValueMap(out)
		if e, a := len(out2), len(in); e != a {
			t.Errorf("Unexpected len at idx %d", idx)
		}
		if e, a := in, out2; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected value at idx %d", idx)
		}
	}
}

type TimeValueTestCase struct {
	in        int64
	outSecs   time.Time
	outMillis time.Time
}

var testCasesTimeValue = []TimeValueTestCase{
	{
		in:        int64(1501558289000),
		outSecs:   time.Unix(1501558289, 0),
		outMillis: time.Unix(1501558289, 0),
	},
	{
		in:        int64(1501558289001),
		outSecs:   time.Unix(1501558289, 0),
		outMillis: time.Unix(1501558289, 1*1000000),
	},
}

func TestSecondsTimeValue(t *testing.T) {
	for idx, testCase := range testCasesTimeValue {
		out := SecondsTimeValue(&testCase.in)
		if e, a := testCase.outSecs, out; e != a {
			t.Errorf("Unexpected value for time value at %d", idx)
		}
	}
}

func TestMillisecondsTimeValue(t *testing.T) {
	for idx, testCase := range testCasesTimeValue {
		out := MillisecondsTimeValue(&testCase.in)
		if e, a := testCase.outMillis, out; e != a {
			t.Errorf("Unexpected value for time value at %d", idx)
		}
	}
}

var testCasesToHumanisedString = []numberConversionTest{
	// Basic number cases with humanisation
	{"int64", numberConversionInputData._int64, "1,234,567"},
	{"int32", numberConversionInputData._int32, "1,234,567"},
	{"int16", numberConversionInputData._int16, "1,234,567"},
	{"int8", numberConversionInputData._int8, "100"},
	{"int", numberConversionInputData._int, "1,234,567"},
	{"uint64", numberConversionInputData._uint64, "1,234,567"},
	{"uint32", numberConversionInputData._uint32, "1,234,567"},
	{"uint16", numberConversionInputData._uint16, "12,345"},
	{"uint8", numberConversionInputData._uint8, "100"},
	{"uint", numberConversionInputData._uint, "1,234,567"},
	{"byte", numberConversionInputData._byte, "100"},
	{"float64", numberConversionInputData._float64, "1,234.12345678"},
	{"float32", numberConversionInputData._float32, "1,234.1234130859375"},

	// Pointer cases with humanisation
	{"*int64", &numberConversionInputData._int64, "1,234,567"},
	{"*int32", &numberConversionInputData._int32, "1,234,567"},
	{"*int16", &numberConversionInputData._int16, "1,234,567"},
	{"*int8", &numberConversionInputData._int8, "100"},
	{"*int", &numberConversionInputData._int, "1,234,567"},
	{"*uint64", &numberConversionInputData._uint64, "1,234,567"},
	{"*uint32", &numberConversionInputData._uint32, "1,234,567"},
	{"*uint16", &numberConversionInputData._uint16, "12,345"},
	{"*uint8", &numberConversionInputData._uint8, "100"},
	{"*uint", &numberConversionInputData._uint, "1,234,567"},
	{"*byte", &numberConversionInputData._byte, "100"},
	{"*float64", &numberConversionInputData._float64, "1,234.12345678"},
	{"*float32", &numberConversionInputData._float32, "1,234.1234130859375"},

	// Nil pointer cases
	{"nil *string", (*string)(nil), ""},
	{"nil *bool", (*bool)(nil), "<nil>"},
	{"nil *int", (*int)(nil), "0"},
	{"nil *float64", (*float64)(nil), "0"},
	{"nil *time.Time", (*time.Time)(nil), ""},

	// Other types
	{"string", "test", "test"},
	{"bool", true, "true"},
	{"time.Time", time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), "2024-01-01T00:00:00Z"},
	{"[]byte", []byte("test"), "test"},
	{"map", numberConversionInputData._map, "map[a:b]"},
	{"struct", numberConversionInputData._struct, "{a:A}"},
	{"enum", numberConversionInputData._enum, "a-val"},
}

func TestToHumanisedString(t *testing.T) {
	for _, test := range testCasesToHumanisedString {
		executeToStringTest(t, test, ToHumanisedString)
	}
}
