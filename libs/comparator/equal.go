package comparator

import (
	"reflect"
)

func getValue(x interface{}) reflect.Value {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}

// IsVariableValueEqual : check if 2 variable values are the same
func IsVariableValueEqual(actual, expected interface{}) bool {
	v1 := getValue(actual)
	v2 := getValue(expected)
	return reflect.DeepEqual(v1.Interface(), v2.Interface())
}

// IsStructValueEqual : check if 2 structs values are the same
func IsStructValueEqual(actual, expected interface{}, excludedFields ...string) bool {
	// 1. Get reflect value of 2 input argument.
	v1 := getValue(actual)
	v2 := getValue(expected)

	// 2. Iterate over all fields in the struct
	for i := 0; i < v1.NumField(); i++ {
		fieldName := v1.Type().Field(i).Name
		if SliceContainsString(excludedFields, fieldName) {
			continue
		}
		if v1.Field(i).IsZero() && !v2.Field(i).IsZero() {
			return false
		}
		if !v1.Field(i).IsZero() && v2.Field(i).IsZero() {
			return false
		}
		if v1.Field(i).IsZero() && v2.Field(i).IsZero() {
			continue
		}
		field1 := getValue(v1.Field(i).Interface())
		field2 := getValue(v2.Field(i).Interface())
		if !reflect.DeepEqual(field1.Interface(), field2.Interface()) {
			return false
		}
	}
	return true
}
