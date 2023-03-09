package services

import (
	"context"
	"reflect"
)

const ServiceKey = "services"

type serviceMap map[reflect.Type]reflect.Value

func NewServiceContext(c context.Context) context.Context {
	if c.Value(ServiceKey) == nil {
		return context.WithValue(c, ServiceKey, make(serviceMap))
	} else {
		// 871 Chapter 32 ■ Creating a Web Platform
		return c
	}
}
