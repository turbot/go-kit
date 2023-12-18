package types

import "reflect"

func IsSimpleType(v any) bool {
	kind := reflect.TypeOf(v).Kind()
	switch kind {
	case
		reflect.Bool,
		reflect.String,
		reflect.Float32,
		reflect.Float64,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		return true
	default:
		return false
	}
}

func IsCollectionType(v any) bool {
	kind := reflect.TypeOf(v).Kind()
	switch kind {
	case
		reflect.Slice,
		reflect.Array,
		reflect.Map:
		return true
	default:
		return false
	}
}
