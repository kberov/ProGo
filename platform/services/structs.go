package services

import (
	"context"
	"errors"
	"reflect"
)

// Adding Support for Resolving Struct Fields

// These functions inspect the fields defined by a struct and attempt to
// resolve them using the defined services. Any fields whose type is not an
// interface or for which there is no service are skipped. The
// PopulateForContextWithExtras function allows additional values to be
// provided for struct fields.  Listing 32-25 defines a struct whose fields
// declare dependencies on services.

func Populate(target any) error {
	return PopulateForContext(context.Background(), target)
}
func PopulateForContext(c context.Context, target any) (err error) {
	return PopulateForContextWithExtras(c, target,
		make(map[reflect.Type]reflect.Value))
}
func PopulateForContextWithExtras(c context.Context, target any,
	extras map[reflect.Type]reflect.Value) (err error) {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() == reflect.Ptr &&
		targetValue.Elem().Kind() == reflect.Struct {
		targetValue = targetValue.Elem()
		for i := 0; i < targetValue.Type().NumField(); i++ {
			fieldVal := targetValue.Field(i)
			if fieldVal.CanSet() {
				if extra, ok := extras[fieldVal.Type()]; ok {
					fieldVal.Set(extra)
				} else {
					_ = resolveServiceFromValue(c, fieldVal.Addr())
				}
			}
		}
	} else {
		err = errors.New("Type cannot be used as target")
	}
	return
}
