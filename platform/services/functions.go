package services

import (
	"context"
	"errors"
	"reflect"
)

// The CallForContext function receives a function and uses services to produce
// the values that are used as arguments to invoke the function. The Call
// function is a convenience for use when a Context isn’t available. The
// implementation of this feature relies on the code used to invoke factory
// functions in Listing 32-22. Listing 32-23 demonstrates how calling functions
// directly can simplify the use of services.

func Call(target any, otherArgs ...any) ([]any, error) {
	return CallForContext(context.Background(), target, otherArgs...)
}

func CallForContext(c context.Context, target any, otherArgs ...any) (results []any, err error) {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() == reflect.Func {
		resultVals := invokeFunction(c, targetValue, otherArgs...)
		results = make([]any, len(resultVals))
		for i := 0; i < len(resultVals); i++ {
			results[i] = resultVals[i].Interface()
		}
	} else {
		err = errors.New("Only functions can be invoked")
	}
	return
}
