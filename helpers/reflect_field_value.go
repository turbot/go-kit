package helpers

import (
	"reflect"
	"strings"

	"github.com/tkrajina/go-reflector/reflector"
)

// GetFieldValueFromInterface uses reflection to return the value of the given property path
func GetFieldValueFromInterface(i interface{}, fieldName string) (interface{}, bool) {
	// unescape property name
	fieldName = UnescapePropertyName(fieldName)
	obj := reflector.New(i)
	if arrayProperty, index, ok := IsFieldArray(fieldName); ok {
		if arrVal, ok := GetFieldValueFromInterface(i, arrayProperty); ok {
			return GetArrayValue(arrVal, index)
		}
		return nil, false
	}

	if reflect.TypeOf(i).Kind() == reflect.Map {
		return obj.GetByKey(fieldName)
	}

	val, err := obj.Field(fieldName).Get()
	return val, err == nil
}

// GetNestedFieldValueFromInterface uses reflection to return the value of the given nested property path
func GetNestedFieldValueFromInterface(item interface{}, propertyPath string) (interface{}, bool) {
	var value interface{}
	var ok bool
	parent := item
	var pathSegments = strings.Split(propertyPath, ".")
	for _, fieldName := range pathSegments {
		// if there are any dots encoded in this segment, decode
		fieldName = UnescapePropertyName(fieldName)
		value, ok = GetFieldValueFromInterface(parent, fieldName)
		if !ok {
			return nil, false
		}
		// update parent for next iteration
		parent = value
	}
	return value, true
}

const propertyPathDotEscape = "$steampipe_escaped_dot$"

// helpers to support property names containing a dot

// EscapePropertyName replaces any '.' characters in the property name with propertyPathDotEscape
func EscapePropertyName(name string) string {
	return strings.Replace(name, ".", propertyPathDotEscape, -1)
}

// UnescapePropertyName replaces any propertyPathDotEscape occurrences with "."
func UnescapePropertyName(name string) string {
	return strings.Replace(name, propertyPathDotEscape, ".", -1)
}
