package helpers

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

func GetArrayValue(i interface{}, index int) (interface{}, bool) {
	if reflect.TypeOf(i).Kind() != reflect.Slice {
		return nil, false
	}
	arrayItems := reflect.ValueOf(i)
	for idx := 0; idx < arrayItems.Len(); idx++ {
		if idx == index {
			return arrayItems.Index(idx).Interface(), true
		}
	}
	return nil, false
}

// IsZero uses reflection to determine whether the given value is the zero value of it's type
func IsZero(i interface{}) bool {
	if i == nil {
		return true
	}
	v := reflect.ValueOf(i)
	z := reflect.Zero(v.Type())

	switch v.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Array, reflect.Struct:
		return reflect.DeepEqual(v.Interface(), z.Interface())
	default:
		// Compare other types directly:
		return v.Interface() == z.Interface()
	}
}

// IsNil uses reflection to determine whether the given value is nil
// this is needed as a simple 'val == nil' check does not work for an interface
// https://mangatmodi.medium.com/go-check-nil-interface-the-right-way-d142776edef1
func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

// InstantiateType returns a interface representing the zero value for the specified type.
func InstantiateType(t reflect.Type) interface{} {
	return reflect.Zero(t).Interface()
}

// DereferencePointer checks if val is a pointer, and if so, dereferences it
func DereferencePointer(val interface{}) interface{} {
	// If the value is a pointer to a non-struct, get its value and use that.
	reflectVal := reflect.ValueOf(val)
	if reflectVal.Kind() == reflect.Ptr {
		if reflectVal.IsNil() {
			// If the pointer is nil, then the value is just nil
			val = nil
		} else {
			// Otherwise, we dereference the pointer
			reflectVal = reflect.Indirect(reflectVal)
			val = reflectVal.Interface()
		}
	}
	return val
}

/*
		TODO: add support for multi-dimensional arrays
	 	like - arr[1][2] arr[1][2][3] and so on..
*/
func IsFieldArray(fieldName string) (string, int, bool) {
	r := regexp.MustCompile(`^(.*)\[(\d+)\]$`)
	captureGroups := r.FindStringSubmatch(fieldName)
	if len(captureGroups) == 0 {
		return "", 0, false
	}
	arrayName := captureGroups[1]
	// check if arrayName contains index and brackets - arr[12]
	// this would indicate a multi-level array which we do not support at present
	subGroups := r.FindStringSubmatch(arrayName)
	if len(subGroups) > 0 {
		return "", 0, false
	}
	arrayIndex, err := strconv.Atoi(captureGroups[2])
	return arrayName, arrayIndex, err == nil
}

// ExecuteMethod use reflection to invoke method. We do not support functions which expect parameters.
func ExecuteMethod(item interface{}, methodName string) (returnValues []interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	var ptr reflect.Value
	var value reflect.Value
	var finalMethod reflect.Value

	value = reflect.ValueOf(item)

	if value.Type().Kind() == reflect.Ptr {
		ptr = value
		value = ptr.Elem()
	} else {
		ptr = reflect.New(reflect.TypeOf(item))
		temp := ptr.Elem()
		temp.Set(value)
	}

	// check for method on value
	method := value.MethodByName(methodName)
	if method.IsValid() {
		finalMethod = method
	}
	// check for method on pointer
	method = ptr.MethodByName(methodName)
	if method.IsValid() {
		finalMethod = method
	}

	if !finalMethod.IsValid() {
		return nil, fmt.Errorf("method %s is not valid", methodName)
	}

	methodResults := finalMethod.Call([]reflect.Value{})
	returnValues = make([]interface{}, len(methodResults))
	for i, r := range methodResults {
		returnValues[i] = r.Interface()
		switch v := returnValues[i].(type) {
		case error:
			err = v
			return
		}
	}
	return
}

// StructToMap uses reflection to convert a struct to a map
func StructToMap(s any) map[string]any {
	result := make(map[string]any)
	val := reflect.ValueOf(s)

	// We only accept structs
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		fmt.Println("Expecting a struct")
		return nil
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.Field(i).Interface()

		result[field.Name] = value
	}

	return result
}
