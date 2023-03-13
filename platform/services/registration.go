package services

import (
	"reflect"
	"sync"
)

func AddTransient(factoryFunc any) (err error) {
	return addService(Transient, factoryFunc)
}
func AddScoped(factoryFunc any) (err error) {
	return addService(Scoped, factoryFunc)
}

// A little more work is required for the singleton lifecycle, and the
// AddSingleton function creates a wrapper around the factory function that
// ensures that it is executed only once, for the first request to resolve the
// service. This ensures that there is only one instance of the implementation
// struct created and that it won’t be created until the first time it is
// needed.
func AddSingleton(factoryFunc any) (err error) {
	factoryFuncVal := reflect.ValueOf(factoryFunc)
	if factoryFuncVal.Kind() == reflect.Func && factoryFuncVal.Type().NumOut() == 1 {
		var results []reflect.Value
		once := sync.Once{}
		wrapper := reflect.MakeFunc(factoryFuncVal.Type(),
			func([]reflect.Value) []reflect.Value {
				once.Do(func() {
					results = invokeFunction(nil, factoryFuncVal)
				})
				return results
			})
		err = addService(Singleton, wrapper.Interface())
	}
	return
}
