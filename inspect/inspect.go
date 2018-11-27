package inspect

import (
	"reflect"
)

func RemovePtr(in interface{}) interface{} {
	return UnderlyingValueOf(in).Interface()
}

func inTypes(in interface{}, types ...reflect.Kind) bool {
	kind := reflect.TypeOf(in).Kind()
	for _, t := range types {
		if kind == t {
			return true
		}
	}

	return false
}

func IsStruct(in interface{}) bool {
	switch v := in.(type) {
	case reflect.Type:
		return v.Kind() == reflect.Struct
	case reflect.Value:
		return v.Kind() == reflect.Struct
	default:
		return UnderlyingTypeOf(in).Kind() == reflect.Struct
	}
}

func IsPtrToStruct(in interface{}) bool {
	switch v := in.(type) {
	case reflect.Type:
		return v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct
	case reflect.Value:
		return v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct
	default:
		t := reflect.TypeOf(in)
		return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
	}
}

func UnderlyingValueOf(in interface{}) (v reflect.Value) {
	v = reflect.ValueOf(in)
	if v.Type().Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}

func UnderlyingTypeOf(in interface{}) (t reflect.Type) { //todo: other elemtype
	t = reflect.TypeOf(in)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return
}

func GetFieldsMap(in interface{}) (m map[string]interface{}) {
	return getFieldsMap(in)
}

func getFieldsMap(in interface{}) (m map[string]interface{}) {
	v := UnderlyingValueOf(in)
	if !IsStruct(in) {
		panic("inspect: getFieldsMap must be called with struct")
	}
	m = make(map[string]interface{})
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		if FieldIsExported(field) {
			m[field.Name] = v.Field(i).Interface()
		}
	}
	return
}


func FieldIsExported(f reflect.StructField) bool {
	return f.PkgPath == ""
}
