package services

import (
	"context"
	"errors"
	"reflect"
)

func GetService(target any) error {
	return GetServiceForContext(context.Background(), target)
}

// The GetServiceForContext accepts a context and a pointer to a value that can
// be set using reflection.  For convenience, the GetService function resolves
// a service using the background context.
func GetServiceForContext(c context.Context, target any) (err error) {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() == reflect.Ptr &&
		targetValue.Elem().CanSet() {
		err = resolveServiceFromValue(c, targetValue)
	} else {
		err = errors.New("Type cannot be used as target")
	}
	return
}
