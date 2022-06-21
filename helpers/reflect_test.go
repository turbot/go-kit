package helpers

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	input    string
	expected interface{}
}

////  GetNestedFieldValueFromInterface ////
type GetNestedFieldValueFromInterfaceTest struct {
	input     interface{}
	fieldName string
	expected  interface{}
}

type S1 struct {
	S      string
	I      int
	M      map[string]int
	M2     map[string]*string
	Str    S2
	Pb     *bool
	NilMap map[string]string
}
type S2 struct {
	S         string
	I         int
	M         map[string]int
	Str       S3
	Arr       []interface{}
	NestedArr []S3
}
type S3 struct {
	S string
	I int
	M map[string]int
}
type S4 struct {
	Arr []S2
}

var (
	a = "11"
	b = "12"
)

var s3Instance = S3{S: "foo S3", I: 3, M: map[string]int{"a": 31, "b": 32}}
var s2Instance = S2{S: "foo S2", I: 2, M: map[string]int{"a": 21, "b": 22}, Str: s3Instance, Arr: []interface{}{"arr 0", "arr 1", "arr 2", []interface{}{"12"}}}
var s2InstanceNested = S2{S: "foo S2", I: 2, M: map[string]int{"a": 21, "b": 22}, Str: s3Instance, Arr: []interface{}{"arr 0", "arr 1", "arr 2"}, NestedArr: []S3{s3Instance}}
var s1Instance = S1{S: "foo S1", I: 1, M: map[string]int{"a": 11, "b": 12}, M2: map[string]*string{"a": &a, "b": &b}, Str: s2Instance, Pb: nil}
var s4Instance = S4{Arr: []S2{s2Instance, s2InstanceNested}}

var testCasesIsFieldArray = map[string]Test{
	"correct segment": {
		input: "arr[10]",
		expected: map[string]interface{}{
			"name":  "arr",
			"index": 10,
		},
	},
	"invalid segment": {
		input: "arr[10][12]",
		expected: map[string]interface{}{
			"name":  "",
			"index": 0,
		},
	},
	"snakecase segment": {
		input: "ip_config[10]",
		expected: map[string]interface{}{
			"name":  "ip_config",
			"index": 10,
		},
	},
	"camelcase segment": {
		input: "ipConfig[10]",
		expected: map[string]interface{}{
			"name":  "ipConfig",
			"index": 10,
		},
	},
	"alphaNumeric segment": {
		input: "ipConfig-cidr[10]",
		expected: map[string]interface{}{
			"name":  "ipConfig-cidr",
			"index": 10,
		},
	},
}

var testCasesGetNestedFieldValueFromInterface = map[string]GetNestedFieldValueFromInterfaceTest{
	"missing property": {
		input:     &s1Instance,
		fieldName: "MISSING",
		expected:  "MISSING",
	},

	"struct pointer double level map": {
		input:     &s1Instance,
		fieldName: "M.a",
		expected:  11,
	},
	"struct pointer double level *string map": {
		input:     &s1Instance,
		fieldName: "M2.a",
		expected:  &a,
	},
	"struct double level string": {
		input:     s1Instance,
		fieldName: "Str.S",
		expected:  "foo S2",
	},
	"struct double level string escaped": {
		input:     s1Instance,
		fieldName: fmt.Sprintf("%s.%s", EscapePropertyName("Str"), EscapePropertyName("S")),
		expected:  "foo S2",
	},
	"Array property path: Nested multi-dimensional failure": {
		input:     s1Instance,
		fieldName: "Str.Arr[0][3]",
		expected:  "MISSING",
	},
	"Array property path: struct double level": {
		input:     s1Instance,
		fieldName: "Str.Arr[0]",
		expected:  "arr 0",
	},
	"Array property path: nested struct": {
		input:     s4Instance,
		fieldName: "Arr[0].Arr[0]",
		expected:  "arr 0",
	},
	"Array property path: deep-nested struct": {
		input:     s4Instance,
		fieldName: "Arr[1].NestedArr[0].S",
		expected:  "foo S3",
	},
	"Array property path: outOfBound": {
		input:     s4Instance,
		fieldName: "Arr[0].Arr[4]",
		expected:  "MISSING",
	},
	"struct double level int": {
		input:     s1Instance,
		fieldName: "Str.I",
		expected:  2,
	},
	"struct triple level map": {
		input:     s1Instance,
		fieldName: "Str.M.a",
		expected:  21,
	},
	"struct triple level string": {
		input:     s1Instance,
		fieldName: "Str.Str.S",
		expected:  "foo S3",
	},
	"struct triple level int": {
		input:     s1Instance,
		fieldName: "Str.Str.I",
		expected:  3,
	},
	"struct quad level map": {
		input:     s1Instance,
		fieldName: "Str.Str.M.a",
		expected:  31,
	},
	"struct single level string": {
		input:     s1Instance,
		fieldName: "S",
		expected:  "foo S1",
	},
	"struct pointer single level string": {
		input:     &s1Instance,
		fieldName: "S",
		expected:  "foo S1",
	},
	"struct single level int": {
		input:     s1Instance,
		fieldName: "I",
		expected:  1,
	},
	"struct single level int (escaped)": {
		input:     s1Instance,
		fieldName: EscapePropertyName("I"),
		expected:  1,
	},
	"struct pointer single level int": {
		input:     &s1Instance,
		fieldName: "I",
		expected:  1,
	},
}

func TestGetNestedFieldValueFromInterface(t *testing.T) {
	for name, test := range testCasesGetNestedFieldValueFromInterface {
		result, found := GetNestedFieldValueFromInterface(test.input, test.fieldName)
		log.Print(name, "  ", result)
		if !found {
			fmt.Printf("is nil %v, %v", result, result == nil)
			if test.expected != "MISSING" {
				t.Errorf(`Test: '%s'' FAILED : expected %v, got NOT FOUND`, name, test.expected)
			}
		} else if !reflect.DeepEqual(test.expected, result) {
			t.Errorf(`Test: '%s'' FAILED : expected %v, got %v`, name, test.expected, result)
		}
	}
}

func TestIsFieldArray(t *testing.T) {
	for name, test := range testCasesIsFieldArray {
		arr, idx, _ := IsFieldArray(test.input)
		result := map[string]interface{}{
			"name":  arr,
			"index": idx,
		}
		log.Print(arr, "  ", idx, name)
		assert.ObjectsAreEqual(test.expected, result)
	}
}

//// IsZero ////
type GetIsZeroTest struct {
	input    interface{}
	expected bool
}

var nilArray []string = nil
var emptyArray = []string{}
var nilMap map[string]string = nil
var emptyMap = map[string]string{}
var nilStruct S1
var nilTime time.Time
var nilTimePtr *time.Time = nil

var testCasesIsZero = map[string]GetIsZeroTest{
	"bool false": {
		input:    false,
		expected: true,
	},
	"time nil": {
		input:    nilTime,
		expected: true,
	},
	"time ptr nil": {
		input:    nilTimePtr,
		expected: true,
	},
	"int 0": {
		input:    0,
		expected: true,
	},
	"array nil": {
		input:    nilArray,
		expected: true,
	},
	"array empty": {
		input:    emptyArray,
		expected: false,
	},
	"map nil": {
		input:    nilMap,
		expected: true,
	},
	"map empty": {
		input:    emptyMap,
		expected: false,
	},

	"struct nil": {
		input:    nilStruct,
		expected: true,
	},
	"struct": {
		input:    s1Instance,
		expected: false,
	},
}

func TestIsZero(t *testing.T) {
	for name, test := range testCasesIsZero {
		isZero := IsZero(test.input)
		if isZero != test.expected {
			t.Errorf(`Test: '%s'' FAILED : expected %v, got %v`, name, test.expected, isZero)
		}
	}
}

type GetExecuteMethodTest struct {
	input        interface{}
	functionName string
	expected     interface{}
}

var testCasesExecuteMethod = map[string]GetExecuteMethodTest{
	"get data by value": {
		input:        s3Instance,
		functionName: "GetDataByValue",
		expected:     []interface{}{s3Instance},
	},
	"get data by pointer": {
		input:        &s3Instance,
		functionName: "GetDataByPointer",
		expected:     []interface{}{s3Instance},
	},
	"invalid function name": {
		input:        &s3Instance,
		functionName: "GetData",
		expected:     "ERROR",
	},
	"panic method": {
		input:        &s3Instance,
		functionName: "PanicFunc",
		expected:     "ERROR",
	},
	"method which excepts a paramemter": {
		input:        &s3Instance,
		functionName: "GetDataByParam",
		expected:     "ERROR",
	},
	"missing function name": {
		input:        &s3Instance,
		functionName: "",
		expected:     "ERROR",
	},
	"get multiple values": {
		input:        &s3Instance,
		functionName: "GetMultipleReturns",
		expected:     []interface{}{"value-1", "value-2", nil},
	},
	"return error": {
		input:        &s3Instance,
		functionName: "ReturnError",
		expected:     "ERROR",
	},
}

func TestExecuteMethod(t *testing.T) {
	for name, test := range testCasesExecuteMethod {
		returnValues, err := ExecuteMethod(test.input, test.functionName)
		if err != nil {
			if test.expected != "ERROR" {
				t.Errorf(`Test: '%s'' FAILED : unexpected error %v`, name, err)
			}
			continue
		}
		if test.expected == "ERROR" {
			t.Errorf(`Test: '%s'' FAILED : expected an error but didn't receive one`, name)
			continue
		}
		expectedValues := test.expected.([]interface{})
		if !reflect.DeepEqual(expectedValues, returnValues) {
			t.Errorf(`Test: '%s'' FAILED : expected %v, got %v`, name, test.expected, returnValues)
		}

	}
}

func (s S3) GetDataByValue() S3 {
	return s3Instance
}

func (s *S3) GetDataByPointer() S3 {
	return s3Instance
}

func (s S3) PanicFunc() {
	panic("PANIC")
}

func (s *S3) GetMultipleReturns() (string, string, error) {
	return "value-1", "value-2", nil
}

func (s *S3) ReturnError() (*string, error) {
	return nil, errors.New("Error")
}

func (s *S3) GetDataByParam(param string) string {
	return strings.ToUpper(param)
}

////  SetNestedFieldInInterface ////
//type SetNestedFieldInInterfaceTest struct {
//	input     interface{}
//	fieldName string
//	value     interface{}
//	expected  interface{}
//}
//
//var testCasesSetNestedFieldInInterface = map[string]SetNestedFieldInInterfaceTest{
//	"struct single level string": {
//		input:     &S1{},
//		fieldName: "S",
//		value:     "string-val",
//		expected:  &S1{S: "string-val"},
//	},
//	"struct single level int": {
//		input:     &S1{},
//		fieldName: "I",
//		value:     100,
//		expected:  &S1{I: 100},
//	},
//	"struct single level map": {
//		input:     &S1{},
//		fieldName: "M",
//		value:     map[string]int{"A": 1},
//		expected:  &S1{M: map[string]int{"A": 1}},
//	},
//	"struct 2 level string": {
//		input:     &S1{},
//		fieldName: "Str.S",
//		value:     map[string]int{"A": 1},
//		expected:  &S1{Str: S2{S: "string-val"}},
//	},
//}
//
//func TestSetNestedFieldInInterface(t *testing.T) {
//
//	a := s1Instance
//	var b interface{} = a
//	c := &b
//	fmt.Println(c)
//	//
//	//reflect.ValueOf(&s3Instance).Elem().FieldByName("S").Set(reflect.ValueOf("ARGGGGGG"))
//	//
//	//s, _ := GetFieldValueFromInterface(s1Instance, "Str")
//	//str := s.(S2)
//	//str.S = "Foo"
//	//obj := reflector.New(s)
//	//f := obj.Field("S")
//	//c := f.IsSettable()
//	//fmt.Println(c)
//	//
//	//obj = reflector.New(str)
//	//f = obj.Field("S")
//	//c = f.IsSettable()
//	//fmt.Println(c)
//	//
//	//obj = reflector.New(s2Instance)
//	//f = obj.Field("S")
//	//c = f.IsSettable()
//	//fmt.Println(c)
//	//
//	//obj = reflector.New(s1Instance)
//	//f = obj.Field("S")
//	//c = f.IsSettable()
//	//fmt.Println(c)
//	//
//	//s = reflect.New(reflect.TypeOf(S1{}))
//	//
//	////s = InstantiateType(reflect.TypeOf(S1{}))
//	//obj = reflector.New(s)
//	//f = obj.Field("S")
//	//c = f.IsSettable()
//	//err := f.Set("FOOOOOO")
//	//fmt.Println(err)
//
//	for name, test := range testCasesSetNestedFieldInInterface {
//		result, success := SetNestedFieldInInterface(test.input, test.fieldName, test.value)
//
//		if !success {
//			if test.expected != "FAIL" {
//				t.Errorf("`Test: '%s'' FAILED : expected success but SetNestedFieldInInterface failed", name)
//			}
//		} else if !reflect.DeepEqual(test.expected, result) {
//			t.Errorf(`Test: '%s'' FAILED : expected %v, got %v`, name, test.expected, result)
//		}
//	}
//}
