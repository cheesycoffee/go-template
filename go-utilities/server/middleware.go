package server

import (
	"context"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

// MiddlewareFunc type
type MiddlewareFunc func(context.Context) context.Context

// MiddlewareGroup type
type MiddlewareGroup map[string][]MiddlewareFunc

// Add register full method to middleware
func (mw MiddlewareGroup) Add(fullMethod string, middlewareFunc ...MiddlewareFunc) {
	mw[fullMethod] = middlewareFunc
}

// AddProto register proto for grpc middleware
func (mw MiddlewareGroup) AddProto(protoDesc protoreflect.FileDescriptor, handler interface{}, middlewareFunc ...MiddlewareFunc) {
	serviceName := fmt.Sprintf("/%s/", protoDesc.Services().Get(0).FullName())
	var fullMethodName string
	switch h := handler.(type) {
	case string:
		fullMethodName = serviceName + h
	default:
		fullMethodName = serviceName + getFuncName(handler)
	}
	mw[fullMethodName] = middlewareFunc
}

func getFuncName(fn interface{}) string {
	defer func() { recover() }()
	handlerName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	return strings.TrimSuffix(strings.TrimPrefix(filepath.Ext(handlerName), "."), "-fm") // if `fn` is method, trim `-fm`
}
