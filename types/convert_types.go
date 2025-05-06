package types

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/btubbs/datetime"
	"github.com/dustin/go-humanize"
	"github.com/mitchellh/mapstructure"
)

// SafeString converts a value from *string to string. If it is already a string, just return it
// if the interface is not a string or *string, return empty string
func SafeString(x interface{}) string {
	str, _ := CastString(x)
	return str
}

// CastString casts an interface to a string
// if the interface is not a string or *string, return empty string and a 'false' result
// This is the same as SafeString except it also returns a success flag
func CastString(x interface{}) (string, bool) {
	switch v := x.(type) {
	case *string:
		return StringValue(v), true
	case string:
		return v, true
	}
	return "", false
}

// ToString converts most types value to a string
func ToString(x interface{}) string {
	switch v := x.(type) {
	case *string:
		return StringValue(v)
	case *bool:
		if v != nil {
			return strconv.FormatBool(*v)
		}
	case *int8:
		return strconv.FormatInt(int64(Int8Value(v)), 10)
	case *int16:
		return strconv.FormatInt(int64(Int16Value(v)), 10)
	case *int32:
		return strconv.FormatInt(int64(Int32Value(v)), 10)
	case *int64:
		return strconv.FormatInt(Int64Value(v), 10)
	case *int:
		return strconv.Itoa(IntValue(v))
	case *uint8:
		return strconv.FormatUint(uint64(Uint8Value(v)), 10)
	case *uint16:
		return strconv.FormatUint(uint64(Uint16Value(v)), 10)
	case *uint32:
		return strconv.FormatUint(uint64(Uint32Value(v)), 10)
	case *uint64:
		return strconv.FormatUint(Uint64Value(v), 10)
	case *uint:
		return strconv.FormatUint(uint64(UintValue(v)), 10)
	case *float32:
		return strconv.FormatFloat(float64(Float32Value(v)), 'f', -1, 32)
	case *float64:
		return strconv.FormatFloat(Float64Value(v), 'f', -1, 64)
	case *time.Time:
		if v != nil {
			return v.Format(time.RFC3339)
		}
		return ""
	case string:
		return v
	case bool:
		return strconv.FormatBool(v)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case time.Time:
		return v.Format(time.RFC3339)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case []byte:
		return string(v)
	}
	// fall back to sprintf
	return fmt.Sprintf("%+v", x)
}

// ToHumanisedString converts most types value to a humanised string
func ToHumanisedString(x interface{}) string {
	switch v := x.(type) {
	case *int8:
		return humanize.Comma(int64(Int8Value(v)))
	case *int16:
		return humanize.Comma(int64(Int16Value(v)))
	case *int32:
		return humanize.Comma(int64(Int32Value(v)))
	case *int64:
		return humanize.Comma(Int64Value(v))
	case *int:
		return humanize.Comma(int64(IntValue(v)))
	case *uint8:
		return humanize.Comma(int64(Uint8Value(v)))
	case *uint16:
		return humanize.Comma(int64(Uint16Value(v)))
	case *uint32:
		return humanize.Comma(int64(Uint32Value(v)))
	case *uint64:
		return humanize.Comma(int64(Uint64Value(v)))
	case *uint:
		return humanize.Comma(int64(UintValue(v)))
	case *float32:
		return humanize.Commaf(float64(Float32Value(v)))
	case *float64:
		return humanize.Commaf(Float64Value(v))
	case float64:
		return humanize.Commaf(v)
	case float32:
		return humanize.Commaf(float64(v))
	case int8:
		return humanize.Comma(int64(v))
	case int16:
		return humanize.Comma(int64(v))
	case int32:
		return humanize.Comma(int64(v))
	case int:
		return humanize.Comma(int64(v))
	case int64:
		return humanize.Comma(v)
	case uint8:
		return humanize.Comma(int64(v))
	case uint16:
		return humanize.Comma(int64(v))
	case uint32:
		return humanize.Comma(int64(v))
	case uint64:
		return humanize.Comma(int64(v))
	case uint:
		return humanize.Comma(int64(v))
	default:
		// call ToString to get a string representation of the value
		return ToString(v)
	}
}

// FloatToString converts interface to a string representation of a float.
// Panic if interface is not convertible or parseable as a float
func FloatToString(x interface{}) string {
	switch v := x.(type) {
	case *float32:
		return strconv.FormatFloat(float64(Float32Value(v)), 'f', -1, 32)
	case *float64:
		return strconv.FormatFloat(Float64Value(v), 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case *string:
		i, err := strconv.ParseFloat(StringValue(v), 64)
		if err == nil {
			return strconv.FormatFloat(i, 'f', -1, 64)
		}
		// if there is a parse error, just fall through to panic
		// dereference pointer for panic message
		x = StringValue(v)
	case string:
		i, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return strconv.FormatFloat(i, 'f', -1, 64)
		}
		// if there is a parse error, just fall through to panic
	}
	panic(fmt.Sprintf(`invalid value '%v' passed to FloatToString`, x))
}

// IntToString converts interface to a string representation of an int.
// Panic if interface is not convertible or parseable as a int
func IntToString(x interface{}) string {
	switch v := x.(type) {
	case *int8:
		return strconv.FormatInt(int64(Int8Value(v)), 10)
	case *int16:
		return strconv.FormatInt(int64(Int16Value(v)), 10)
	case *int32:
		return strconv.Itoa(int(Int32Value(v)))
	case *int64:
		return strconv.FormatInt(Int64Value(v), 10)
	case *int:
		return strconv.Itoa(IntValue(v))
	case *uint8:
		return strconv.FormatUint(uint64(Uint8Value(v)), 10)
	case *uint16:
		return strconv.FormatUint(uint64(Uint16Value(v)), 10)
	case *uint32:
		return strconv.FormatUint(uint64(Uint32Value(v)), 10)
	case *uint64:
		return strconv.FormatUint(Uint64Value(v), 10)
	case *uint:
		return strconv.FormatUint(uint64(UintValue(v)), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.Itoa(int(v))
	case int64:
		return strconv.FormatInt(v, 10)
	case int:
		return strconv.Itoa(v)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case *string:
		i, err := strconv.ParseInt(StringValue(v), 10, 64)
		if err == nil {
			return strconv.FormatInt(i, 10)
		}
		// if there is a parse error, just fall through to panic
		// dereference pointer for panic message
		x = StringValue(v)
	case string:
		i, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			return strconv.FormatInt(i, 10)
		}
	}
	panic(fmt.Sprintf(`invalid value '%v' passed to IntToString`, x))
}

// ToInt64 converts interface to an int64
// if interface is not convertible or parseable as an int64, return an error
func ToInt64(i interface{}) (int64Val int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	switch v := i.(type) {
	case *string:
		int64Val, err = strconv.ParseInt(*v, 10, 64)
	case string:
		int64Val, err = strconv.ParseInt(v, 10, 64)
	default:
		var intVal int
		err = mapstructure.Decode(i, &intVal)
		int64Val = int64(intVal)
	}

	return
}

// ToFloat64 converts interface to an float64
// if interface is not convertible or parseable as a float64, return an error
func ToFloat64(i interface{}) (doubleVal float64, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	switch v := i.(type) {
	case *string:
		doubleVal, err = strconv.ParseFloat(*v, 64)
	case string:
		doubleVal, err = strconv.ParseFloat(v, 64)
	default:
		err = mapstructure.Decode(i, &doubleVal)
	}

	return
}

// ToTime converts interface to an time
// if interface is not convertible or parseable as a time, return an error
func ToTime(i interface{}) (timeVal time.Time, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	switch v := i.(type) {
	case *time.Time:
		timeVal = *v
	case time.Time:
		timeVal = v
	case string:
		timeVal, err = datetime.Parse(SafeString(v), time.UTC)
	default:
		err = fmt.Errorf("cannot convert %v to a time.Time", v)
	}

	return
}

// ToBoolPtr converts interface to a bool pointer
// if the value is nil, or there is a conversion error, return nil
func ToBoolPtr(i interface{}) *bool {
	if i == nil {
		return nil
	}
	res, err := ToBool(i)
	if err != nil {
		return nil
	}
	return &res
}

// ToBool converts interface to an bool
// if interface is not convertible or parseable as a bool, return an error
func ToBool(i interface{}) (boolValue bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	switch v := i.(type) {
	case *string:
		if v == nil {
			err = fmt.Errorf(`null string pointer passed to ToBool`)
			break
		}
		boolValue = StringToBool(*v)
	case string:
		boolValue = StringToBool(v)
	case bool:
		boolValue = v
	case *bool:
		if v == nil {
			err = fmt.Errorf(`null bool pointer passed to ToBool`)
			break
		}
		boolValue = *v
	default:
		err = fmt.Errorf(`invalid value '%v' passed to ToBool, must one of: bool, *bool, string, *string`, i)
	}
	return
}

// StringToBool :: try to map a string value to true or false, using a dictionary of expected values
func StringToBool(str string) bool {
	str = strings.ToLower(str)
	stringMap := map[string]bool{
		"true":      true,
		"on":        true,
		"enabled":   true,
		"yes":       true,
		"1":         true,
		"false":     false,
		"off":       false,
		"disabled":  false,
		"no":        false,
		"suspended": false,
		"0":         false,
	}
	if res, ok := stringMap[str]; ok {
		return res
	}
	// invalid value
	panic(fmt.Sprintf("invalid value '%s' passed to StringToBool", str))
}

// String :: return a pointer to the string value passed in.
func String(v string) *string {
	return &v
}

// StringValue :: return the value of the string pointer passed in or
// "" if the pointer is nil.
func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

// StringSlice converts a slice of string values into a slice of
// string pointers
func StringSlice(src []string) []*string {
	dst := make([]*string, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// StringValueSlice converts a slice of string pointers into a slice of
// string values
func StringValueSlice(src []*string) []string {
	dst := make([]string, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Float64SliceToStringSlice converts a slice of float64 into a slice of
// string values
func Float64SliceToStringSlice(src []float64) []string {
	dst := make([]string, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = strconv.FormatFloat(src[i], 'f', -1, 64)
	}
	return dst
}

// StringMap converts a string map of string values into a string
// map of string pointers
func StringMap(src map[string]string) map[string]*string {
	dst := make(map[string]*string)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// StringValueMap converts a string map of string pointers into a string
// map of string values
func StringValueMap(src map[string]*string) map[string]string {
	dst := make(map[string]string)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Bool :: return a pointer to the bool value passed in.
func Bool(v bool) *bool {
	return &v
}

// BoolValue :: return the value of the bool pointer passed in or
// false if the pointer is nil.
func BoolValue(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}

// BoolSlice converts a slice of bool values into a slice of
// bool pointers
func BoolSlice(src []bool) []*bool {
	dst := make([]*bool, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// BoolValueSlice converts a slice of bool pointers into a slice of
// bool values
func BoolValueSlice(src []*bool) []bool {
	dst := make([]bool, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// BoolMap converts a string map of bool values into a string
// map of bool pointers
func BoolMap(src map[string]bool) map[string]*bool {
	dst := make(map[string]*bool)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// BoolValueMap converts a string map of bool pointers into a string
// map of bool values
func BoolValueMap(src map[string]*bool) map[string]bool {
	dst := make(map[string]bool)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Int :: return a pointer to the int value passed in.
func Int(v int) *int {
	return &v
}

// IntValue :: return the value of the int pointer passed in or
// 0 if the pointer is nil.
func IntValue(v *int) int {
	if v != nil {
		return *v
	}
	return 0
}

// IntSlice converts a slice of int values into a slice of
// int pointers
func IntSlice(src []int) []*int {
	dst := make([]*int, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// IntValueSlice converts a slice of int pointers into a slice of
// int values
func IntValueSlice(src []*int) []int {
	dst := make([]int, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// IntMap converts a string map of int values into a string
// map of int pointers
func IntMap(src map[string]int) map[string]*int {
	dst := make(map[string]*int)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// IntValueMap converts a string map of int pointers into a string
// map of int values
func IntValueMap(src map[string]*int) map[string]int {
	dst := make(map[string]int)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Uint :: return a pointer to the uint value passed in.
func Uint(v uint) *uint {
	return &v
}

// UintValue :: return the value of the uint pointer passed in or
// 0 if the pointer is nil.
func UintValue(v *uint) uint {
	if v != nil {
		return *v
	}
	return 0
}

// UintSlice converts a slice of uint values uinto a slice of
// uint pointers
func UintSlice(src []uint) []*uint {
	dst := make([]*uint, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// UintValueSlice converts a slice of uint pointers uinto a slice of
// uint values
func UintValueSlice(src []*uint) []uint {
	dst := make([]uint, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// UintMap converts a string map of uint values uinto a string
// map of uint pointers
func UintMap(src map[string]uint) map[string]*uint {
	dst := make(map[string]*uint)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// UintValueMap converts a string map of uint pointers uinto a string
// map of uint values
func UintValueMap(src map[string]*uint) map[string]uint {
	dst := make(map[string]uint)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Int8 :: return a pointer to the int8 value passed in.
func Int8(v int8) *int8 {
	return &v
}

// Int8Value :: return the value of the int8 pointer passed in or
// 0 if the pointer is nil.
func Int8Value(v *int8) int8 {
	if v != nil {
		return *v
	}
	return 0
}

// Int8Slice converts a slice of int8 values into a slice of
// int8 pointers
func Int8Slice(src []int8) []*int8 {
	dst := make([]*int8, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// Int8ValueSlice converts a slice of int8 pointers into a slice of
// int8 values
func Int8ValueSlice(src []*int8) []int8 {
	dst := make([]int8, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Int8Map converts a string map of int8 values into a string
// map of int8 pointers
func Int8Map(src map[string]int8) map[string]*int8 {
	dst := make(map[string]*int8)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// Int8ValueMap converts a string map of int8 pointers into a string
// map of int8 values
func Int8ValueMap(src map[string]*int8) map[string]int8 {
	dst := make(map[string]int8)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Int16 :: return a pointer to the int16 value passed in.
func Int16(v int16) *int16 {
	return &v
}

// Int16Value :: return the value of the int16 pointer passed in or
// 0 if the pointer is nil.
func Int16Value(v *int16) int16 {
	if v != nil {
		return *v
	}
	return 0
}

// Int16Slice converts a slice of int16 values into a slice of
// int16 pointers
func Int16Slice(src []int16) []*int16 {
	dst := make([]*int16, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// Int16ValueSlice converts a slice of int16 pointers into a slice of
// int16 values
func Int16ValueSlice(src []*int16) []int16 {
	dst := make([]int16, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Int16Map converts a string map of int16 values into a string
// map of int16 pointers
func Int16Map(src map[string]int16) map[string]*int16 {
	dst := make(map[string]*int16)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// Int16ValueMap converts a string map of int16 pointers into a string
// map of int16 values
func Int16ValueMap(src map[string]*int16) map[string]int16 {
	dst := make(map[string]int16)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Int32 :: return a pointer to the int32 value passed in.
func Int32(v int32) *int32 {
	return &v
}

// Int32Value :: return the value of the int32 pointer passed in or
// 0 if the pointer is nil.
func Int32Value(v *int32) int32 {
	if v != nil {
		return *v
	}
	return 0
}

// Int32Slice converts a slice of int32 values into a slice of
// int32 pointers
func Int32Slice(src []int32) []*int32 {
	dst := make([]*int32, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// Int32ValueSlice converts a slice of int32 pointers into a slice of
// int32 values
func Int32ValueSlice(src []*int32) []int32 {
	dst := make([]int32, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Int32Map converts a string map of int32 values into a string
// map of int32 pointers
func Int32Map(src map[string]int32) map[string]*int32 {
	dst := make(map[string]*int32)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// Int32ValueMap converts a string map of int32 pointers into a string
// map of int32 values
func Int32ValueMap(src map[string]*int32) map[string]int32 {
	dst := make(map[string]int32)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Int64 :: return a pointer to the int64 value passed in.
func Int64(v int64) *int64 {
	return &v
}

// Int64Value :: return the value of the int64 pointer passed in or
// 0 if the pointer is nil.
func Int64Value(v *int64) int64 {
	if v != nil {
		return *v
	}
	return 0
}

// Int64Slice converts a slice of int64 values into a slice of
// int64 pointers
func Int64Slice(src []int64) []*int64 {
	dst := make([]*int64, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// Int64ValueSlice converts a slice of int64 pointers into a slice of
// int64 values
func Int64ValueSlice(src []*int64) []int64 {
	dst := make([]int64, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Int64Map converts a string map of int64 values into a string
// map of int64 pointers
func Int64Map(src map[string]int64) map[string]*int64 {
	dst := make(map[string]*int64)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// Int64ValueMap converts a string map of int64 pointers into a string
// map of int64 values
func Int64ValueMap(src map[string]*int64) map[string]int64 {
	dst := make(map[string]int64)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Uint8 :: return a pointer to the uint8 value passed in.
func Uint8(v uint8) *uint8 {
	return &v
}

// Uint8Value :: return the value of the uint8 pointer passed in or
// 0 if the pointer is nil.
func Uint8Value(v *uint8) uint8 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint8Slice converts a slice of uint8 values into a slice of
// uint8 pointers
func Uint8Slice(src []uint8) []*uint8 {
	dst := make([]*uint8, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// Uint8ValueSlice converts a slice of uint8 pointers into a slice of
// uint8 values
func Uint8ValueSlice(src []*uint8) []uint8 {
	dst := make([]uint8, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Uint8Map converts a string map of uint8 values into a string
// map of uint8 pointers
func Uint8Map(src map[string]uint8) map[string]*uint8 {
	dst := make(map[string]*uint8)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// Uint8ValueMap converts a string map of uint8 pointers into a string
// map of uint8 values
func Uint8ValueMap(src map[string]*uint8) map[string]uint8 {
	dst := make(map[string]uint8)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Uint16 :: return a pointer to the uint16 value passed in.
func Uint16(v uint16) *uint16 {
	return &v
}

// Uint16Value :: return the value of the uint16 pointer passed in or
// 0 if the pointer is nil.
func Uint16Value(v *uint16) uint16 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint16Slice converts a slice of uint16 values into a slice of
// uint16 pointers
func Uint16Slice(src []uint16) []*uint16 {
	dst := make([]*uint16, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// Uint16ValueSlice converts a slice of uint16 pointers into a slice of
// uint16 values
func Uint16ValueSlice(src []*uint16) []uint16 {
	dst := make([]uint16, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Uint16Map converts a string map of uint16 values into a string
// map of uint16 pointers
func Uint16Map(src map[string]uint16) map[string]*uint16 {
	dst := make(map[string]*uint16)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// Uint16ValueMap converts a string map of uint16 pointers into a string
// map of uint16 values
func Uint16ValueMap(src map[string]*uint16) map[string]uint16 {
	dst := make(map[string]uint16)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Uint32 :: return a pointer to the uint32 value passed in.
func Uint32(v uint32) *uint32 {
	return &v
}

// Uint32Value :: return the value of the uint32 pointer passed in or
// 0 if the pointer is nil.
func Uint32Value(v *uint32) uint32 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint32Slice converts a slice of uint32 values into a slice of
// uint32 pointers
func Uint32Slice(src []uint32) []*uint32 {
	dst := make([]*uint32, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// Uint32ValueSlice converts a slice of uint32 pointers into a slice of
// uint32 values
func Uint32ValueSlice(src []*uint32) []uint32 {
	dst := make([]uint32, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Uint32Map converts a string map of uint32 values into a string
// map of uint32 pointers
func Uint32Map(src map[string]uint32) map[string]*uint32 {
	dst := make(map[string]*uint32)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// Uint32ValueMap converts a string map of uint32 pointers into a string
// map of uint32 values
func Uint32ValueMap(src map[string]*uint32) map[string]uint32 {
	dst := make(map[string]uint32)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Uint64 :: return a pointer to the uint64 value passed in.
func Uint64(v uint64) *uint64 {
	return &v
}

// Uint64Value :: return the value of the uint64 pointer passed in or
// 0 if the pointer is nil.
func Uint64Value(v *uint64) uint64 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint64Slice converts a slice of uint64 values into a slice of
// uint64 pointers
func Uint64Slice(src []uint64) []*uint64 {
	dst := make([]*uint64, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// Uint64ValueSlice converts a slice of uint64 pointers into a slice of
// uint64 values
func Uint64ValueSlice(src []*uint64) []uint64 {
	dst := make([]uint64, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Uint64Map converts a string map of uint64 values into a string
// map of uint64 pointers
func Uint64Map(src map[string]uint64) map[string]*uint64 {
	dst := make(map[string]*uint64)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// Uint64ValueMap converts a string map of uint64 pointers into a string
// map of uint64 values
func Uint64ValueMap(src map[string]*uint64) map[string]uint64 {
	dst := make(map[string]uint64)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Float32 :: return a pointer to the float32 value passed in.
func Float32(v float32) *float32 {
	return &v
}

// Float32Value :: return the value of the float32 pointer passed in or
// 0 if the pointer is nil.
func Float32Value(v *float32) float32 {
	if v != nil {
		return *v
	}
	return 0
}

// Float32Slice converts a slice of float32 values into a slice of
// float32 pointers
func Float32Slice(src []float32) []*float32 {
	dst := make([]*float32, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// Float32ValueSlice converts a slice of float32 pointers into a slice of
// float32 values
func Float32ValueSlice(src []*float32) []float32 {
	dst := make([]float32, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Float32Map converts a string map of float32 values into a string
// map of float32 pointers
func Float32Map(src map[string]float32) map[string]*float32 {
	dst := make(map[string]*float32)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// Float32ValueMap converts a string map of float32 pointers into a string
// map of float32 values
func Float32ValueMap(src map[string]*float32) map[string]float32 {
	dst := make(map[string]float32)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Float64 :: return a pointer to the float64 value passed in.
func Float64(v float64) *float64 {
	return &v
}

// Float64Value :: return the value of the float64 pointer passed in or
// 0 if the pointer is nil.
func Float64Value(v *float64) float64 {
	if v != nil {
		return *v
	}
	return 0
}

// Float64Slice converts a slice of float64 values into a slice of
// float64 pointers
func Float64Slice(src []float64) []*float64 {
	dst := make([]*float64, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// Float64ValueSlice converts a slice of float64 pointers into a slice of
// float64 values
func Float64ValueSlice(src []*float64) []float64 {
	dst := make([]float64, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Float64Map converts a string map of float64 values into a string
// map of float64 pointers
func Float64Map(src map[string]float64) map[string]*float64 {
	dst := make(map[string]*float64)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// Float64ValueMap converts a string map of float64 pointers into a string
// map of float64 values
func Float64ValueMap(src map[string]*float64) map[string]float64 {
	dst := make(map[string]float64)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Time :: return a pointer to the time.Time value passed in.
func Time(v time.Time) *time.Time {
	return &v
}

// TimeValue :: return the value of the time.Time pointer passed in or
// time.Time{} if the pointer is nil.
func TimeValue(v *time.Time) time.Time {
	if v != nil {
		return *v
	}
	return time.Time{}
}

// SecondsTimeValue converts an int64 pointer to a time.Time value
// representing seconds since Epoch or time.Time{} if the pointer is nil.
func SecondsTimeValue(v *int64) time.Time {
	if v != nil {
		return time.Unix(*v/1000, 0)
	}
	return time.Time{}
}

// MillisecondsTimeValue converts an int64 pointer to a time.Time value
// representing milliseconds sinch Epoch or time.Time{} if the pointer is nil.
func MillisecondsTimeValue(v *int64) time.Time {
	if v != nil {
		return time.Unix(0, *v*1000000)
	}
	return time.Time{}
}

// TimeUnixMilli :: return a Unix timestamp in milliseconds from "January 1, 1970 UTC".
// The result is undefined if the Unix time cannot be represented by an int64.
// Which includes calling TimeUnixMilli on a zero Time is undefined.
//
// This utility is useful for service API's such as CloudWatch Logs which require
// their unix time values to be in milliseconds.
//
// See Go stdlib https://golang.org/pkg/time/#Time.UnixNano for more information.
func TimeUnixMilli(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond/time.Nanosecond)
}

// TimeSlice converts a slice of time.Time values into a slice of
// time.Time pointers
func TimeSlice(src []time.Time) []*time.Time {
	dst := make([]*time.Time, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// TimeValueSlice converts a slice of time.Time pointers into a slice of
// time.Time values
func TimeValueSlice(src []*time.Time) []time.Time {
	dst := make([]time.Time, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// TimeMap converts a string map of time.Time values into a string
// map of time.Time pointers
func TimeMap(src map[string]time.Time) map[string]*time.Time {
	dst := make(map[string]*time.Time)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// TimeValueMap converts a string map of time.Time pointers into a string
// map of time.Time values
func TimeValueMap(src map[string]*time.Time) map[string]time.Time {
	dst := make(map[string]time.Time)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}
