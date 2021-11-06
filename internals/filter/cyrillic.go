package filter

import (
	"reflect"
	"strings"
)

func removeCyrillic(s string) string {
	var validS strings.Builder
	for _, v := range s {
		if v <= 1024 || v >= 1279 {
			validS.WriteString(string(v))
		}
	}
	return validS.String()
}
func SomeFunc(v interface{}) {
	val := reflect.ValueOf(v).Elem()
	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		switch f.Kind() {
		case reflect.String:
			f.SetString(removeCyrillic(f.String()))
		case reflect.Ptr:
			if f.Elem().Kind() == reflect.String {
				f.Elem().SetString(removeCyrillic(f.Elem().String()))
			} else if f.Elem().Kind() == reflect.Struct {
				SomeFunc(f.Interface())
			}
		case reflect.Struct:
			SomeFunc(f.Addr().Interface())
		}
	}
}
