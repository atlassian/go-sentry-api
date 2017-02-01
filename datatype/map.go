package datatype

import (
	"reflect"
)

var mapping = map[string]reflect.Value{
	"message":    reflect.ValueOf(Message{}),
	"stacktrace": reflect.ValueOf(Stacktrace{}),
	"exception":  reflect.ValueOf(Exception{}),
	"request":    reflect.ValueOf(Request{}),
	"template":   reflect.ValueOf(Template{}),
	"user":       reflect.ValueOf(User{}),
	"query":      reflect.ValueOf(Query{}),
}

// GetMapping is what you can use to convert datatype to real datatype
func GetMapping(inter string) (string, interface{}) {
	if value, ok := mapping[inter]; ok {
		return inter, value.Interface()
	}
	return inter, nil
}
