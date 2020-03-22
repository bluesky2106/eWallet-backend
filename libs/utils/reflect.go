package utils

import "reflect"

// ValueOf : returns the origin value of x
func ValueOf(x interface{}) reflect.Value {
	v := reflect.ValueOf(x)
	for {
		if v.Kind() != reflect.Ptr {
			break
		}
		v = v.Elem()
	}
	return v
}

// TypeOf : returns the origin type of x
func TypeOf(x interface{}) reflect.Type {
	v := reflect.ValueOf(x)
	for {
		if v.Kind() != reflect.Ptr {
			break
		}
		v = v.Elem()
	}

	return reflect.TypeOf(v.Interface())
}
