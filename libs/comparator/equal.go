package comparator

import (
	"reflect"

	"github.com/bluesky2106/eWallet-backend/libs/utils"
)

// IsVariableValueEqual : check if 2 variable values are the same
func IsVariableValueEqual(actual, expected interface{}) bool {
	v1 := utils.ValueOf(actual)
	v2 := utils.ValueOf(expected)
	return reflect.DeepEqual(v1.Interface(), v2.Interface())
}

// IsStructValueEqual : check if 2 structs values are the same
func IsStructValueEqual(actual, expected interface{}, excludedFields ...string) bool {
	// 1. Get reflect value of 2 input argument.
	v1 := utils.ValueOf(actual)
	v2 := utils.ValueOf(expected)

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
		field1 := utils.ValueOf(v1.Field(i).Interface())
		field2 := utils.ValueOf(v2.Field(i).Interface())
		if !reflect.DeepEqual(field1.Interface(), field2.Interface()) {
			return false
		}
	}
	return true
}
